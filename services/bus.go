package services

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func ServiceConnection(serviceName string) {
	// Create a TCP connection to the SOA bus
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Printf("Could not connect to SOA bus: %v\n", err)
		return
	}
	defer conn.Close()

	// Send initialization message
	initMsg := "00010sinit" + serviceName
	fmt.Printf("Sending init message: %s\n", initMsg)
	_, err = conn.Write([]byte(initMsg))
	if err != nil {
		fmt.Printf("Could not send init message: %v\n", err)
		return
	}

	sinit := true

	for {
		fmt.Println("Waiting for transaction")

		// Read the length of the response
		lenBuf := make([]byte, 5)
		_, err := conn.Read(lenBuf)
		if err != nil {
			fmt.Printf("Could not read response length: %v\n", err)
			return
		}

		responseLen, err := strconv.Atoi(strings.TrimSpace(string(lenBuf)))
		if err != nil {
			fmt.Printf("Invalid response length: %v\n", err)
			return
		}
		fmt.Printf("Expected response length: %d\n", responseLen)

		// Read the actual response
		responseBuf := make([]byte, responseLen)
		_, err = conn.Read(responseBuf)
		if err != nil {
			fmt.Printf("Could not read response: %v\n", err)
			return
		}

		fmt.Printf("Processing response: %s\n", string(responseBuf))

		if sinit {
			sinit = false
			fmt.Println("Received sinit answer")
		} else {
			fmt.Println("Send answer")
			answerMsg := "00013" + serviceName + "Received"
			fmt.Printf("Sending answer message: %s\n", answerMsg)
			_, err = conn.Write([]byte(answerMsg))
			if err != nil {
				fmt.Printf("Could not send answer message: %v\n", err)
				return
			}
		}
	}
}
