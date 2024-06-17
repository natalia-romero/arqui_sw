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
func ResponseSoa(serviceName, data string) bool {
	response, err := SendToSOABus(serviceName, data)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	fmt.Println("Response from SOA bus:", response)
	return strings.Contains(response, "OK")
}
func ServiceExec(serviceName, data string) bool {
	if ResponseSoa(serviceName, data) {
		servNumber := serviceName[len(serviceName)-1:]
		if servNumber == "1" {
			HandleService1()
			return true
		} else if servNumber == "2" {
			return HandleService2()
		} else if servNumber == "3" {
			HandleService3()
			return true
		} else if servNumber == "4" {
			HandleService4()
			return true
		} else if servNumber == "5" {
			HandleService5()
			return true
		} else if servNumber == "6" {
			HandleService6()
			return true
		} else if servNumber == "7" {
			HandleService7()
			return true
		} else if servNumber == "8" {
			HandleService9()
			return true
		} else if servNumber == "9" {
			HandleService9()
			return true
		}
	}
	return false
}
