// bus.go
package main

import (
	"fmt"
	"net"
)

// sendTransaction sends a transaction to the SOA bus and returns the response
func sendTransaction(serviceName, data string) (string, error) {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		return "", fmt.Errorf("could not connect to SOA bus: %w", err)
	}
	defer conn.Close()

	serviceName = fmt.Sprintf("%-5s", serviceName) // Ensure serviceName is exactly 5 characters
	transaction := fmt.Sprintf("%05d%s%s", len(serviceName)+len(data), serviceName, data)
	_, err = conn.Write([]byte(transaction))
	if err != nil {
		return "", fmt.Errorf("could not send transaction: %w", err)
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("could not read response: %w", err)
	}

	response := string(buffer[:n])
	return response, nil
}
