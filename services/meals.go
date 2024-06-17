package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/manifoldco/promptui"
)

type Service5 struct {
	service1 *Service1
}

func (s *Service5) Execute() {
	s.service1.Execute()
	s.mealsCRUD()
}
func (s *Service5) mealsCRUD() {
	for {
		menu := promptui.Select{
			Label: "ADMINISTRACIÓN DE PLATOS, INGRESE OPCIÓN",
			Items: []string{"Agregar", "Ver", "Editar", "Salir"},
		}

		_, choice, err := menu.Run()
		if err != nil {
			log.Fatal(err)
		}

		switch choice {
		case "Agregar":
			fmt.Println("Ha ingresado a la opción de crear plato.")
			var name, description string
			var price int
			fmt.Print("Ingrese nombre: ")
			fmt.Scanln(&name)
			fmt.Print("Ingrese descripción: ")
			fmt.Scanln(&description)
			fmt.Print("Ingrese precio: ")
			fmt.Scanln(&price)

			var existingMeal string
			err := s.service1.db.QueryRow("SELECT name FROM meals WHERE name = ?", name).Scan(&existingMeal)
			if err != nil && err != sql.ErrNoRows {
				log.Fatal(err)
			}

			if existingMeal != "" {
				fmt.Println("El plato ya existe.")
			} else {
				_, err := s.service1.db.Exec("INSERT INTO meals (name, description, price) VALUES (?, ?, ?)", name, description, price)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Plato creado exitosamente.")
			}
		case "Ver":
			fmt.Println("Ha ingresado a la opción de ver platos.")
			rows, err := s.service1.db.Query("SELECT id, name, description, price FROM meals")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			for rows.Next() {
				var id, price int
				var name, description string
				err := rows.Scan(&id, &name, &description, &price)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("ID: %d, Name: %s, Description: %s, Price: %d\n", id, name, description, price)
			}

			if err := rows.Err(); err != nil {
				log.Fatal(err)
			}
		case "Editar":
			fmt.Println("Ha ingresado a la opción de editar plato.")
			var id int
			fmt.Print("Ingrese ID del plato a editar: ")
			fmt.Scanln(&id)
			for {
				editMenu := promptui.Select{
					Label: "Seleccione la opción a editar",
					Items: []string{"Nombre", "Descripción", "Precio", "Salir"},
				}

				_, editChoice, err := editMenu.Run()
				if err != nil {
					log.Fatal(err)
				}

				switch editChoice {
				case "Nombre":
					fmt.Print("Ingrese nuevo nombre: ")
					var name string
					fmt.Scanln(&name)
					_, err := s.service1.db.Exec("UPDATE meals SET name = ? WHERE id = ?", name, id)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Nombre del plato actualizado exitosamente.")
				case "Descripción":
					fmt.Print("Ingrese nueva descripción: ")
					var description string
					fmt.Scanln(&description)
					_, err := s.service1.db.Exec("UPDATE meals SET description = ? WHERE id = ?", description, id)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Descripción actualizada exitosamente.")
				case "Precio":
					fmt.Print("Ingrese nuevo precio: ")
					var price int
					fmt.Scanln(&price)
					_, err := s.service1.db.Exec("UPDATE meals SET price = ? WHERE id = ?", price, id)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Precio actualizado exitosamente.")
				case "Salir":
					fmt.Println("Saliendo de la opción de editar plato.")
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

func HandleService5() {
	service1 := &Service1{}
	service5 := &Service5{service1: service1}
	service5.Execute()
}
