package services

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// SendToSOABus sends a message to the SOA bus and returns the response
func SendToSOABus(serviceName, data string) (string, error) {
	// Create a TCP connection to the SOA bus
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		return "", fmt.Errorf("could not connect to SOA bus: %v", err)
	}
	defer conn.Close()

	// Prepare the message according to the format NNNNNSSSSSDATA
	dataLen := len(data)
	msg := fmt.Sprintf("%05d%-5s%s", dataLen+len(serviceName), serviceName, data)
	fmt.Println(msg)
	// Send the message
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return "", fmt.Errorf("could not send data to SOA bus: %v", err)
	}

	// Read the length of the response
	reader := bufio.NewReader(conn)
	lenBuf := make([]byte, 5)
	_, err = reader.Read(lenBuf)
	if err != nil {
		return "", fmt.Errorf("could not read response length: %v", err)
	}
	responseLen, err := strconv.Atoi(strings.TrimSpace(string(lenBuf)))
	if err != nil {
		return "", fmt.Errorf("invalid response length: %v", err)
	}

	// Read the actual response
	responseBuf := make([]byte, responseLen)
	_, err = reader.Read(responseBuf)
	if err != nil {
		return "", fmt.Errorf("could not read response: %v", err)
	}

	return string(responseBuf), nil
}
