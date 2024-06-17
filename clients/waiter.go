package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/natalia-romero/arqui_sw/services"
)

func GenerateConnectionWaiter() {
	serviceName := "serv2"
	data := "auth"
	response, err := services.SendToSOABus(serviceName, data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response from SOA bus:", response)
	if strings.Contains(response, "OK") {
		fmt.Println("Conexión correcta con el bus.")
		services.HandleService2()
		for {
			menu := promptui.Select{
				Label: "CLIENTE GARZÓN",
				Items: []string{"Gestión de pedidos", "Salir"},
			}

			_, choice, err := menu.Run()
			if err != nil {
				log.Fatal(err)
			}

			switch choice {
			case "Gestión de pedidos": //listo
				if services.ServiceExec("serv9", "orders") {
					fmt.Println("Servicio ejecutado.")
				} else {
					fmt.Println("No se puede conectar al servicio.")
				}
			case "Salir":
				fmt.Println("Ha ingresado a la opción de salir.")
				return
			}
		}
	} else {
		fmt.Println("No se puede conectar a los servicios.")
	}

}
func main() {
	GenerateConnectionWaiter()
}
