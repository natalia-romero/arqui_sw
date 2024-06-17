package services

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/manifoldco/promptui"
)

type Service9 struct {
	service1 *Service1
}

func (s *Service9) Execute() {
	s.service1.Execute()
	s.orders()
}
func (s *Service9) orders() {
	for {
		fmt.Println("Ha ingresado a la opción de gestionar pedidos.")
		for {
			editMenu := promptui.Select{
				Label: "Seleccione opción",
				Items: []string{"Crear nuevo pedido", "Editar pedido existente", "Salir"},
			}
			_, editChoice, err := editMenu.Run()
			if err != nil {
				log.Fatal(err)
			}

			switch editChoice {
			case "Crear nuevo pedido":
				fmt.Print("Ingrese ID de la mesa: ")
				var table_id string
				fmt.Scanln(&table_id)
				date := time.Now()
				number := rand.Intn(9000) + 1000
				_, err := s.service1.db.Exec("INSERT INTO receipt (table_id, date, number) VALUES (?, ?, ?)", table_id, date, number)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Pedido creado exitosamente.")
			case "Editar pedido existente":
				fmt.Print("Ingrese ID del pedido: ")
				var order_id int
				fmt.Scanln(&order_id)
				fmt.Print("Ingrese ID del plato: ")
				var meal_id int
				fmt.Scanln(&meal_id)
				fmt.Print("Ingrese cantidad: ")
				var quantity int
				fmt.Scanln(&quantity)
				_, err := s.service1.db.Exec("INSERT INTO receipt_detail (receipt_id, meal_id, quantity) VALUES (?, ?, ?)", order_id, meal_id, quantity)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Plato añadido exitosamente.")
			case "Salir":
				fmt.Println("Saliendo de la opción de agregar plato.")
				return
			}
		}
	}
}

func HandleService9() {
	service1 := &Service1{}
	service9 := &Service9{service1: service1}
	service9.Execute()
}
