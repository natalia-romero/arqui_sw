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
		fmt.Println("4. Service 4 (Table CRUD)")
		fmt.Println("5. Service 5 (Meals CRUD)")
		fmt.Println("6. Service 6 (Receipt)")
		fmt.Println("7. Service 7 (Ranking)")
		fmt.Println("8. Service 8 (Sales Export)")
		fmt.Println("9. Service 9 (Orders)")
		fmt.Println("10. Exit")
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
			fmt.Println("Exiting.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
