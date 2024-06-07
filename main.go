package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nexidian/gocliselect"
)

var db *sql.DB

// Service 1 DATABASE
// Service1 DATABASE
type Service1 struct {
	User string
	Pass string
	db   *sql.DB
}

func (s *Service1) Execute() {
	s.User = os.Getenv("DBUSER")
	s.Pass = os.Getenv("DBPASS")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/arqui_sw?charset=utf8&parseTime=True&loc=Local", s.User, s.Pass)
	var err error
	s.db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	pingErr := s.db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Conectado a la base de datos")
	res, err := s.db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	var table string
	fmt.Println("Listado de tablas:")
	for res.Next() {
		res.Scan(&table)
		fmt.Println(" - ", table)
	}
}

// Service2 AUTHENTICATION
type Service2 struct {
	service1 *Service1
}

func (s *Service2) Execute() bool {
	s.service1.Execute()
	return s.authenticateUser()
}

func (s *Service2) authenticateUser() bool {
	var username, password string

	fmt.Print("Ingrese nombre de usuario: ")
	fmt.Scanln(&username)
	fmt.Print("Ingrese contraseña: ")
	fmt.Scanln(&password)

	var storedPassword string
	err := s.service1.db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Usuario no encontrado.")
		} else {
			log.Fatal(err)
		}
		return false
	}

	if password == storedPassword {
		var status int
		err2 := s.service1.db.QueryRow("SELECT status_id FROM users WHERE username = ? AND password = ?", username, password).Scan(&status)
		if err2 != nil {
			log.Fatal(err2)
		}

		if status == 1 {
			fmt.Println("Inicio de sesión exitoso.")
			return true
		} else {
			fmt.Println("Usuario inactivo.")
			return false
		}
	} else {
		fmt.Println("Contraseña incorrecta.")
		return false
	}
}

// Service 3 USERS ADMINISTRATION
type Service3 struct {
	service2 *Service2
}

func (s *Service3) Execute() {
	if s.service2.Execute() {
		for {
			menu := gocliselect.NewMenu("ADMINISTRACIÓN DE USUARIOS, INGRESE OPCIÓN")
			menu.AddItem("Agregar", "create")
			menu.AddItem("Ver", "read")
			menu.AddItem("Editar", "update")
			menu.AddItem("Desactivar", "deactivate")
			menu.AddItem("Salir", "exit")

			choice := menu.Display()
			//fmt.Printf("Opción escogida: %s\n", choice)
			if choice == "create" {
				fmt.Printf("Ha ingresado a la opción de crear usuario.\n")
				var username, password string
				var role int
				fmt.Print("Ingrese nombre de usuario: ")
				fmt.Scanln(&username)
				fmt.Print("Ingrese contraseña: ")
				fmt.Scanln(&password)
				fmt.Print("Ingrese rol (1 admin, 2 garzon): ")
				fmt.Scanln(&role)
				err := s.service2.service1.db.QueryRow("SELECT username FROM users WHERE username = ?", username).Scan(&username)
				if err != nil {
					if err == sql.ErrNoRows {
						s.service2.service1.db.QueryRow("INSERT INTO users (id, username, password, type_user_id, status_id) VALUES (NULL, ?, ?, ?, 1);", username, password, role)
					} else {
						log.Fatal(err)
					}
				}
			} else if choice == "read" {
				fmt.Printf("Ha ingresado a la opción de ver usuario.\n")
			} else if choice == "update" {
				fmt.Printf("Ha ingresado a la opción de editar usuario.\n")
			} else if choice == "deactivate" {
				fmt.Printf("Ha ingresado a la opción de desactivar usuario.\n")
			} else if choice == "exit" {
				fmt.Printf("Ha ingresado a la opción de salir.\n")
				break
			}
		}
	} else {
		fmt.Println("Autenticación fallida. No se puede continuar.")
	}
}

// Main
func main() {
	// create instances of Service1 and Service2
	service1 := &Service1{}
	service2 := &Service2{service1: service1}
	service3 := &Service3{service2: service2}
	service3.Execute()
}
