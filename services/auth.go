package services

import (
	"database/sql"
	"fmt"
	"log"
)

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

func HandleService2() bool {
	service1 := &Service1{}
	service2 := &Service2{service1: service1}
	return service2.Execute()
}
