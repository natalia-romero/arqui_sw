package main

import (
	"fmt"
	"os"

	"github.com/natalia-romero/arqui_sw/services"
)

func main() {
	for {
		fmt.Println("Select a service to connect:")
		fmt.Println("1. Service 1 (Database)")
		fmt.Println("2. Service 2 (Authentication)")
		fmt.Println("3. Service 3 (User CRUD)")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
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
			fmt.Println("Exiting.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
