package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
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
	// Suponiendo que los datos vienen en formato "login=admin&pass=123"
	fmt.Sscanf(data, "login=%s&pass=%s", &username, &password)

	var storedPassword string
	err := s.service1.db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Usuario no encontrado.")
			return false
		} else {
			log.Fatal(err)
		}
	}

	if password == storedPassword {
		fmt.Println("Inicio de sesión exitoso.")
		return true
	} else {
		fmt.Println("Contraseña incorrecta.")
		return false
	}
}

func HandleService2(data string) string {
	service1 := &Service1{}
	service2 := &Service2{service1: service1}
	if service2.Execute() {
		return fmt.Sprintf("%05d%s%sOK Autenticación exitosa", len(data)+18, "auth", data)
	}
	return fmt.Sprintf("%05d%s%sNK Autenticación fallida", len(data)+18, "auth", data)
}
