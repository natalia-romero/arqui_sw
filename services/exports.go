package services

import (
	"database/sql"
	"fmt"
	"log"
)

type Service8 struct {
	service1 *Service1
}

func (s *Service8) Execute() bool {
	s.service1.Execute()
	return s.authenticateUser()
}

func (s *Service8) authenticateUser() bool {
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

func HandleService8() bool {
	service1 := &Service1{}
	service8 := &Service8{service1: service1}
	return service8.Execute()
}
