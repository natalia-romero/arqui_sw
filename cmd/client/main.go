package main

import (
	"fmt"
	"net"
	"time"
)

const busAddress = "localhost:5000"

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

	conn.SetReadDeadline(time.Now().Add(20 * time.Second))

	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		return "", fmt.Errorf("error al leer la respuesta: %v", err)
	}

	return string(response[:n]), nil
}

func main() {
	response, err := sendTransaction("auth", "login=admin&pass=123")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Respuesta: %s\n", response)
	}
}
