package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/manifoldco/promptui"
)

type Service4 struct {
	service1 *Service1
}

func (s *Service4) Execute() {
	s.service1.Execute()
	s.tablesCRUD()
}
func (s *Service4) tablesCRUD() {
	for {
		menu := promptui.Select{
			Label: "ADMINISTRACIÓN DE MESAS, INGRESE OPCIÓN",
			Items: []string{"Agregar", "Ver", "Editar", "Salir"},
		}

		_, choice, err := menu.Run()
		if err != nil {
			log.Fatal(err)
		}

		switch choice {
		case "Agregar":
			fmt.Println("Ha ingresado a la opción de crear mesa.")
			var number int
			fmt.Print("Ingrese número de la mesa: ")
			fmt.Scanln(&number)

			var existingTable string
			err := s.service1.db.QueryRow("SELECT number FROM restaurant_table WHERE number = ?", number).Scan(&existingTable)
			if err != nil && err != sql.ErrNoRows {
				log.Fatal(err)
			}
			if existingTable != "" {
				fmt.Println("La mesa ya existe.")
			} else {
				_, err := s.service1.db.Exec("INSERT INTO restaurant_table (number) VALUES (?)", number)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Mesa creada exitosamente.")
			}
		case "Ver":
			fmt.Println("Ha ingresado a la opción de ver mesas.")
			rows, err := s.service1.db.Query("SELECT id, number, user_id FROM restaurant_table")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			for rows.Next() {
				var id, number int
				var user_id sql.NullInt64
				err := rows.Scan(&id, &number, &user_id)
				if err != nil {
					log.Fatal(err)
				}
				if user_id.Valid {
					fmt.Printf("ID: %d, Number: %d, User ID: %d\n", id, number, user_id.Int64)
				} else {
					fmt.Printf("ID: %d, Number: %d, User ID: NULL\n", id, number)
				}
			}

			if err := rows.Err(); err != nil {
				log.Fatal(err)
			}
		case "Editar":
			fmt.Println("Ha ingresado a la opción de editar mesa.")
			var id int
			fmt.Print("Ingrese ID de la mesa a editar: ")
			fmt.Scanln(&id)
			for {
				editMenu := promptui.Select{
					Label: "Seleccione la opción a editar",
					Items: []string{"Número", "Garzón", "Salir"},
				}

				_, editChoice, err := editMenu.Run()
				if err != nil {
					log.Fatal(err)
				}

				switch editChoice {
				case "Número":
					fmt.Print("Ingrese nuevo numero de mesa: ")
					var number int
					fmt.Scanln(&number)
					_, err := s.service1.db.Exec("UPDATE restaurant_table SET number = ? WHERE id = ?", number, id)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Número actualizado exitosamente.")
				case "Garzón":
					fmt.Print("Ingrese ID del garzon: ")
					var id_waiter int
					fmt.Scanln(&id_waiter)
					_, err := s.service1.db.Exec("UPDATE restaurant_table SET user_id = ? WHERE id = ?", id_waiter, id)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Garzón actualizado exitosamente.")
				case "Salir":
					fmt.Println("Saliendo de la opción de editar mesa.")
				}
				if editChoice == "Salir" {
					break
				}
			}
		case "Salir":
			fmt.Println("Ha ingresado a la opción de salir.")
			return
		}
	}
}

func HandleService4() {
	service1 := &Service1{}
	service4 := &Service4{service1: service1}
	service4.Execute()
}
