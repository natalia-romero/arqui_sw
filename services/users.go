package services

import (
    "fmt"
    "log"
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
                fmt.Print("Ingrese rol (1 admin, 2
