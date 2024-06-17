package main

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/natalia-romero/arqui_sw/services"
)

func GenerateConnectionAdmin() {
	if services.ServiceExec("serv2", "auth") {
		fmt.Println("Servicio ejecutado.")
		for {
			menu := promptui.Select{
				Label: "CLIENTE ADMINISTRADOR",
				Items: []string{"Base de datos", "Gestión de usuarios", "Gestión de mesas", "Gestión de platos", "Boletas", "Ranking", "Exportar ventas", "Salir"},
			}

			_, choice, err := menu.Run()
			if err != nil {
				log.Fatal(err)
			}

			switch choice {
			case "Base de datos":
				if services.ServiceExec("serv1", "databases") { //listo
					fmt.Println("Servicio ejecutado.")
				} else {
					fmt.Println("No se puede conectar al servicio.") //listo
				}
			case "Gestión de usuarios":
				if services.ServiceExec("serv3", "users") {
					fmt.Println("Servicio ejecutado.")
				} else {
					fmt.Println("No se puede conectar al servicio.") //listo
				}
			case "Gestión de mesas":
				if services.ServiceExec("serv4", "tables") {
					fmt.Println("Servicio ejecutado.")
				} else {
					fmt.Println("No se puede conectar al servicio.") //listo
				}
			case "Gestión de platos":
				if services.ServiceExec("serv5", "meals") { //listo
					fmt.Println("Servicio ejecutado.")
				} else {
					fmt.Println("No se puede conectar al servicio.")
				}
			case "Boletas":
				if services.ServiceExec("serv6", "receipt") { //listo
					fmt.Println("Servicio ejecutado.")
				} else {
					fmt.Println("No se puede conectar al servicio.")
				}
			case "Ranking":
				if services.ServiceExec("serv7", "ranking") {
					fmt.Println("Servicio ejecutado.")
				} else {
					fmt.Println("No se puede conectar al servicio.")
				}
			case "Exportar ventas":
				if services.ServiceExec("serv8", "export") {
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
	GenerateConnectionAdmin()
}
