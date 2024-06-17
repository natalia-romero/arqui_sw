package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/manifoldco/promptui"
)

type Service3 struct {
	service2 *Service2
}

func (s *Service3) Execute() {
	if s.service2.Execute() {
		for {
			menu := promptui.Select{
				Label: "ADMINISTRACIÓN DE USUARIOS, INGRESE OPCIÓN",
				Items: []string{"Agregar", "Ver", "Editar", "Desactivar", "Salir"},
			}

			_, choice, err := menu.Run()
			if err != nil {
				log.Fatal(err)
			}

			switch choice {
			case "Agregar":
				fmt.Println("Ha ingresado a la opción de crear usuario.")
				var username, password string
				var role int
				fmt.Print("Ingrese nombre de usuario: ")
				fmt.Scanln(&username)
				fmt.Print("Ingrese contraseña: ")
				fmt.Scanln(&password)
				fmt.Print("Ingrese rol (1 admin, 2 garzon): ")
				fmt.Scanln(&role)

				var existingUser string
				err := s.service2.service1.db.QueryRow("SELECT username FROM users WHERE username = ?", username).Scan(&existingUser)
				if err != nil && err != sql.ErrNoRows {
					log.Fatal(err)
				}

				if existingUser != "" {
					fmt.Println("El usuario ya existe.")
				} else {
					_, err := s.service2.service1.db.Exec("INSERT INTO users (username, password, type_user_id, status_id) VALUES (?, ?, ?, 1)", username, password, role)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Usuario creado exitosamente.")
				}
			case "Ver":
				fmt.Println("Ha ingresado a la opción de ver usuarios.")
				rows, err := s.service2.service1.db.Query("SELECT id, username, password, type_user_id, status_id FROM users")
				if err != nil {
					log.Fatal(err)
				}
				defer rows.Close()

				for rows.Next() {
					var id, typeUserId, statusId int
					var username, password string
					err := rows.Scan(&id, &username, &password, &typeUserId, &statusId)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("ID: %d, Username: %s, Password: %s, Role: %d, Status: %d\n", id, username, password, typeUserId, statusId)
				}

				if err := rows.Err(); err != nil {
					log.Fatal(err)
				}
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
						_, err := s.service2.service1.db.Exec("UPDATE users SET username = ? WHERE id = ?", username, id)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println("Nombre de usuario actualizado exitosamente.")
					case "Contraseña":
						fmt.Print("Ingrese nueva contraseña: ")
						var password string
						fmt.Scanln(&password)
						_, err := s.service2.service1.db.Exec("UPDATE users SET password = ? WHERE id = ?", password, id)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println("Contraseña actualizada exitosamente.")
					case "Tipo de usuario":
						fmt.Print("Ingrese nuevo tipo de usuario (1 admin, 2 garzon): ")
						var typeUserId int
						fmt.Scanln(&typeUserId)
						_, err := s.service2.service1.db.Exec("UPDATE users SET type_user_id = ? WHERE id = ?", typeUserId, id)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println("Tipo de usuario actualizado exitosamente.")
					case "Estado":
						fmt.Print("Ingrese nuevo estado (1 activo, 2 inactivo): ")
						var statusId int
						fmt.Scanln(&statusId)
						_, err := s.service2.service1.db.Exec("UPDATE users SET status_id = ? WHERE id = ?", statusId, id)
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
				_, err := s.service2.service1.db.Exec("UPDATE users SET status_id = 2 WHERE id = ?", id)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Usuario desactivado exitosamente.")
			case "Salir":
				fmt.Println("Ha ingresado a la opción de salir.")
				return
			}
		}
	} else {
		fmt.Println("Autenticación fallida. No se puede continuar.")
	}
}

func HandleService3() {
	service1 := &Service1{}
	service2 := &Service2{service1: service1}
	service3 := &Service3{service2: service2}
	service3.Execute()
}

func soaConnection3() {
	serviceConnection("serv3")
}
