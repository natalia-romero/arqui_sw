package main

import (
	"fmt"
	"os"

	"github.com/natalia-romero/arqui_sw/services"
)

func main() {
	for {
		fmt.Println("Seleccione un servicio a conectar:")
		fmt.Println("1. Servicio 1 (Base de datos)")
		fmt.Println("2. Servicio 2 (Autenticación)")
		fmt.Println("3. Servicio 3 (Usuarios CRUD)")
		fmt.Println("4. Servicio 4 (Mesas CRUD)")
		fmt.Println("5. Servicio 5 (Platos CRUD)")
		fmt.Println("6. Servicio 6 (Boletas)")
		fmt.Println("7. Servicio 7 (Ranking)")
		fmt.Println("8. Servicio 8 (Exportar ventas)")
		fmt.Println("9. Servicio 9 (Pedidos)")
		fmt.Println("10. Salir")
		fmt.Print("Ingrese opción: ")

		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Opción incorrecta.")
			continue
		}

		switch choice {
		case 1:
			services.ServiceConnection("serv1")
		case 2:
			services.ServiceConnection("serv2")
		case 3:
			services.ServiceConnection("serv3")
		case 4:
			services.ServiceConnection("serv4")
		case 5:
			services.ServiceConnection("serv5")
		case 6:
			services.ServiceConnection("serv6")
		case 7:
			services.ServiceConnection("serv7")
		case 8:
			services.ServiceConnection("serv8")
		case 9:
			services.ServiceConnection("serv9")
		case 10:
			fmt.Println("Ha salido del programa.")
			os.Exit(0)
		default:
			fmt.Println("Opción incorrecta.")
		}
	}
}
