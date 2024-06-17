package services

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/manifoldco/promptui"
)

type Service7 struct {
	service1 *Service1
}

func (s *Service7) Execute() {
	s.service1.Execute()
	s.viewRanking()
}
func (s *Service7) viewRanking() {
	for {
		menu := promptui.Select{
			Label: "VISTA DE RANKING, INGRESE OPCIÓN",
			Items: []string{"Ranking garzones", "Ranking platos", "Salir"},
		}

		_, choice, err := menu.Run()
		if err != nil {
			log.Fatal(err)
		}

		switch choice {
		case "Ranking garzones":
			fmt.Println("Ha ingresado a la opción de ver ranking de garzones.")
			rows, err := s.service1.db.Query("SELECT u.username, COUNT(rt.id) AS num_tables FROM users u INNER JOIN restaurant_table rt ON u.id = rt.user_id WHERE u.type_user_id = 2 AND u.status_id = 1 GROUP BY u.username ORDER BY num_tables DESC")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			fmt.Println("- - - - - Ranking de Garzones - - - - -")
			fmt.Println("Usuario garzon - Cantidad de mesas atendidas")
			for rows.Next() {
				var username string
				var numTables int
				err := rows.Scan(&username, &numTables)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s \t\t-\t\t%d\n", username, numTables)
			}
			if err := rows.Err(); err != nil {
				log.Fatal(err)
			}

		case "Ranking platos":
			fmt.Println("- - - - - Ranking de Platos - - - - -")
			fmt.Println("Ha ingresado a la opción de ver los 5 platos más vendidos.")
			rows, err := s.service1.db.Query("SELECT m.name, SUM(rd.quantity) AS total_quantity FROM meals m INNER JOIN receipt_detail rd ON m.id = rd.meal_id GROUP BY m.name ORDER BY total_quantity DESC LIMIT 5")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			fmt.Println("Ranking de los 5 platos más vendidos:")
			fmt.Println("Plato \t-\t Cantidad Vendida")
			for rows.Next() {
				var name string
				var totalQuantity int
				err := rows.Scan(&name, &totalQuantity)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s \t-\t %d\n", name, totalQuantity)
			}
			if err := rows.Err(); err != nil {
				log.Fatal(err)
			}
		case "Salir":
			fmt.Println("Ha ingresado a la opción de salir.")
			return
		}
	}
}

func HandleService7() {
	service1 := &Service1{}
	service7 := &Service7{service1: service1}
	service7.Execute()
}
