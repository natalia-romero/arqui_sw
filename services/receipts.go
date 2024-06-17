package services

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/manifoldco/promptui"
)

type Service6 struct {
	service1 *Service1
}

func (s *Service6) Execute() {
	s.service1.Execute()
	s.receipts()
}
func (s *Service6) receipts() {
	for {
		menu := promptui.Select{
			Label: "ADMINISTRACIÓN DE BOLETAS, INGRESE OPCIÓN",
			Items: []string{"Cerrar mesa", "Ver boleta"},
		}

		_, choice, err := menu.Run()
		if err != nil {
			log.Fatal(err)
		}

		switch choice {
		case "Cerrar mesa":
			fmt.Println("Ha ingresado a la opción de cerrar pedido de mesa.")
			var id int
			fmt.Print("Ingrese ID del pedido: ")
			fmt.Scanln(&id)

			var amount int
			err := s.service1.db.QueryRow("SELECT SUM(m.price * rd.quantity) AS total_amount FROM receipt r INNER JOIN receipt_detail rd ON r.id = rd.receipt_id INNER JOIN meals m ON rd.meal_id = m.id  WHERE r.id = ?", id).Scan(&amount)

			if err != nil {
				log.Fatal(err)
			}

			_, err2 := s.service1.db.Exec("UPDATE receipt SET amount = ? WHERE id = ?", amount, id)
			if err2 != nil {
				log.Fatal(err)
			}
			fmt.Println("Mesa cerrada exitosamente.")

		case "Ver boleta":
			fmt.Println("Ha ingresado a la opción de ver boletas.")
			var id int
			fmt.Print("Ingrese ID del pedido: ")
			fmt.Scanln(&id)
			var amount sql.NullFloat64
			err := s.service1.db.QueryRow("SELECT amount FROM receipt WHERE id = ?", id).Scan(&amount)
			if err != nil {
				log.Fatal(err)
			}
			if amount.Valid {
				fmt.Println("****************** BOLETA GENERADA ******************")
				var (
					receipt_number int
					receipt_date   time.Time
					receipt_amount sql.NullFloat64
					table_id       int
					table_number   int
				)

				err := s.service1.db.QueryRow(`
					SELECT r.number, r.date, r.amount, rt.id, rt.number
					FROM receipt r
					INNER JOIN restaurant_table rt ON r.table_id = rt.id
					WHERE r.id = ?`, id).Scan(&receipt_number, &receipt_date, &receipt_amount, &table_id, &table_number)
				if err != nil {
					if err == sql.ErrNoRows {
						fmt.Println("No se encontró una boleta con ese ID.")
					} else {
						log.Fatal(err)
					}
					return
				}
				fmt.Printf("---------------- INFORMACIÓN ----------------\n")
				fmt.Printf("- ID boleta: %d\n", id)
				fmt.Printf("- Número boleta: %d\n", receipt_number)
				fmt.Printf("- Mesa: ID %d - N° %d\n", table_id, table_number)
				fmt.Printf("- Fecha: %s\n", receipt_date.Format("2006-01-02 15:04:05"))
				fmt.Printf("---------------- PRODUCTOS ----------------\n")
				fmt.Printf("|producto|\t|cantidad|\t|precio|\n\n")

				rows, err := s.service1.db.Query(`
					SELECT m.name, rd.quantity, m.price
					FROM receipt_detail rd
					INNER JOIN meals m ON rd.meal_id = m.id
					WHERE rd.receipt_id = ?`, id)
				if err != nil {
					log.Fatal(err)
				}
				defer rows.Close()

				var total_amount float64
				for rows.Next() {
					var name string
					var quantity int
					var price float64

					err := rows.Scan(&name, &quantity, &price)
					if err != nil {
						log.Fatal(err)
					}

					item_total := price * float64(quantity)
					total_amount += item_total
					fmt.Printf("|%s| \t |%d| \t |%.2f (%.2f)|\n", name, quantity, price, item_total)
				}

				if err := rows.Err(); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("\n---------------- TOTAL ----------------\n")
				if receipt_amount.Valid {
					fmt.Printf("- Total boleta: %.2f\n", receipt_amount.Float64)
				} else {
					fmt.Printf("- Total boleta: %.2f\n", total_amount)
				}
			} else {
				fmt.Printf("- La mesa aún no está cerrada.")
			}
			fmt.Println("")
		case "Editar":
			fmt.Println("Ha ingresado a la opción de editar usuario.")
			var id int
			fmt.Print("Ingrese ID del usuario a editar: ")
			fmt.Scanln(&id)
			for {
				editMenu := promptui.Select{
					Label: "Seleccione la opción a editar",
					Items: []string{"Nombre", "Contraseña", "Tipo de usuario", "Estado", "Salir"},
				}

				_, editChoice, err := editMenu.Run()
				if err != nil {
					log.Fatal(err)
				}

				switch editChoice {
				case "Nombre":
					fmt.Print("Ingrese nuevo nombre de usuario: ")
					var username string
					fmt.Scanln(&username)
					_, err := s.service1.db.Exec("UPDATE users SET username = ? WHERE id = ?", username, id)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Nombre de usuario actualizado exitosamente.")
				case "Contraseña":
					fmt.Print("Ingrese nueva contraseña: ")
					var password string
					fmt.Scanln(&password)
					_, err := s.service1.db.Exec("UPDATE users SET password = ? WHERE id = ?", password, id)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Contraseña actualizada exitosamente.")
				case "Tipo de usuario":
					fmt.Print("Ingrese nuevo tipo de usuario (1 admin, 2 garzon): ")
					var typeUserId int
					fmt.Scanln(&typeUserId)
					_, err := s.service1.db.Exec("UPDATE users SET type_user_id = ? WHERE id = ?", typeUserId, id)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Tipo de usuario actualizado exitosamente.")
				case "Estado":
					fmt.Print("Ingrese nuevo estado (1 activo, 2 inactivo): ")
					var statusId int
					fmt.Scanln(&statusId)
					_, err := s.service1.db.Exec("UPDATE users SET status_id = ? WHERE id = ?", statusId, id)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Estado del usuario actualizado exitosamente.")
				case "Salir":
					fmt.Println("Saliendo de la opción de editar usuario.")
				}
				if editChoice == "Salir" {
					break
				}
			}
		case "Desactivar":
			fmt.Println("Ha ingresado a la opción de desactivar usuario.")
			var id int
			fmt.Print("Ingrese ID del usuario a desactivar: ")
			fmt.Scanln(&id)
			_, err := s.service1.db.Exec("UPDATE users SET status_id = 2 WHERE id = ?", id)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Usuario desactivado exitosamente.")
		case "Salir":
			fmt.Println("Ha ingresado a la opción de salir.")
			return
		}
	}
}

func HandleService6() {
	service1 := &Service1{}
	service6 := &Service6{service1: service1}
	service6.Execute()
}
