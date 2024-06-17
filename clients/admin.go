package main

import (
	"fmt"
	"strings"

	"github.com/natalia-romero/arqui_sw/services"
)

func generateConnection() {
	serviceName := "serv1"
	data := "basededatos"
	response, err := services.SendToSOABus(serviceName, data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response from SOA bus:", response)
	if strings.Contains(response, "OK") {
		fmt.Println("Conectando a la base de datos...")
		services.HandleService1()
	} else {
		fmt.Println("No se puede conectar a la base de datos.")
	}
}

func main() {
	generateConnection()
}
