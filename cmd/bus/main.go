package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const busAddress = "localhost:5001"
const dockerAddress = "localhost:5000"

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("Error al leer la conexión: %v\n", err)
		return
	}

	transaction := string(buf[:n])
	fmt.Printf("Recibido: %s\n", transaction)

	length, _ := strconv.Atoi(transaction[:5])
	service := strings.TrimSpace(transaction[5:10])
	data := transaction[10:length]

	var response string

	switch service {
	case "serv1", "auth", "serv3":
		response, err = forwardToDocker(transaction)
		if err != nil {
			response = fmt.Sprintf("%05d%s%sNK Error al redirigir al Docker", len(transaction), service, data)
		}
	default:
		response = fmt.Sprintf("%05d%s%sNK Servicio no encontrado", len(transaction), service, data)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Printf("Error al escribir la respuesta: %v\n", err)
	}
}

func forwardToDocker(transaction string) (string, error) {
	conn, err := net.Dial("tcp", dockerAddress)
	if err != nil {
		return "", fmt.Errorf("error al conectar al Docker: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(transaction))
	if err != nil {
		return "", fmt.Errorf("error al escribir en la conexión: %v", err)
	}

	conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		return "", fmt.Errorf("error al leer la respuesta: %v", err)
	}

	return string(response[:n]), nil
}

func main() {
	listener, err := net.Listen("tcp", busAddress)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Bus de Datos escuchando en %s\n", busAddress)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error al aceptar la conexión: %v\n", err)
			continue
		}

		go handleConnection(conn)
	}
}
