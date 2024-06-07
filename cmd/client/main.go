package main

import (
	"fmt"
	"net"
	"time"

	"github.com/natalia-romero/arqui_sw/services"
)

const busAddress = "localhost:5001"

func sendTransaction(service, data string) (string, error) {
	service = fmt.Sprintf("%-5s", service)
	dataLength := len(service) + len(data)
	transaction := fmt.Sprintf("%05d%s%s", dataLength, service, data)
	fmt.Printf("Enviando transacción: %s\n", transaction)

	conn, err := net.Dial("tcp", busAddress)
	if err != nil {
		return "", fmt.Errorf("error al conectar al bus: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(transaction))
	if err != nil {
		return "", fmt.Errorf("error al escribir en la conexión: %v", err)
	}

	conn.SetReadDeadline(time.Now().Add(60 * time.Second)) // Aumenta el tiempo de espera

	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		return "", fmt.Errorf("error al leer la respuesta: %v", err)
	}

	return string(response[:n]), nil //xd
}

func main() {
	// Envía una transacción al servicio de administración de usuarios
	response, err := sendTransaction("serv3", "administracionusuarios")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Respuesta: %s\n", response)
		services.HandleService3()
	}
}
