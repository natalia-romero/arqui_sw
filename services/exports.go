package services

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/manifoldco/promptui"
)

type Service8 struct {
	service1 *Service1
}

func (s *Service8) Execute() {
	s.service1.Execute()
	s.exportOrders()
}

func (s *Service8) exportOrders() {
	fmt.Println("Exportar ventas")
	menu := promptui.Select{
		Label: "Seleccione opción",
		Items: []string{"Exportar a PDF", "Exportar a CSV"},
	}
	_, choice, err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := s.service1.db.Query("SELECT r.id, r.number, r.date, r.amount, rt.number AS table_number, u.id AS waiter_id, u.username AS waiter_name  FROM receipt r INNER JOIN restaurant_table rt ON r.table_id = rt.id INNER JOIN users u ON rt.user_id = u.id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	switch choice {
	case "Exportar a PDF":
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 12)

		pdf.Cell(20, 10, "ID Boleta")
		pdf.Cell(20, 10, "N Boleta")
		pdf.Cell(30, 10, "Fecha")
		pdf.Cell(30, 10, "Monto")
		pdf.Cell(20, 10, "N Mesa")
		pdf.Cell(30, 10, "ID Garzon")
		pdf.Cell(40, 10, "Nombre Garzon")
		pdf.Ln(10)

		pdf.SetFont("Arial", "", 12)

		for rows.Next() {
			var (
				id          int
				number      int
				date        time.Time
				amount      sql.NullInt64
				tableNumber int
				waiterID    int
				waiterName  string
			)
			err := rows.Scan(&id, &number, &date, &amount, &tableNumber, &waiterID, &waiterName)
			if err != nil {
				log.Fatal(err)
			}
			amountStr := "NULL"
			if amount.Valid {
				amountStr = fmt.Sprintf("%d", amount.Int64)
			}
			pdf.Cell(20, 10, fmt.Sprintf("%d", id))
			pdf.Cell(20, 10, fmt.Sprintf("%d", number))
			pdf.Cell(40, 10, date.Format("2006-01-02"))
			pdf.Cell(30, 10, amountStr)
			pdf.Cell(20, 10, fmt.Sprintf("%d", tableNumber))
			pdf.Cell(20, 10, fmt.Sprintf("%d", waiterID))
			pdf.Cell(40, 10, waiterName)
			pdf.Ln(10)
		}

		err = pdf.OutputFileAndClose("receipts.pdf")
		if err != nil {
			log.Fatal("Error al crear PDF", err)
		}

		fmt.Println("Se ha exportado la boleta al archivo receipts.pdf")
	case "Exportar a CSV":
		file, err := os.Create("receipts.csv")
		if err != nil {
			log.Fatal("Error al crear PDF", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		//header
		writer.Write([]string{"ID", "N Boleta", "Fecha", "Monto", "N Mesa", "ID Garzon", "Nombre Garzon"})

		//rows
		for rows.Next() {
			var (
				id          int
				number      int
				date        time.Time
				amount      sql.NullInt64
				tableNumber int
				waiterID    int
				waiterName  string
			)
			err := rows.Scan(&id, &number, &date, &amount, &tableNumber, &waiterID, &waiterName)
			if err != nil {
				log.Fatal(err)
			}
			amountStr := "NULL"
			if amount.Valid {
				amountStr = fmt.Sprintf("%d", amount.Int64)
			}
			writer.Write([]string{
				fmt.Sprintf("%d", id),
				fmt.Sprintf("%d", number),
				date.Format("2006-01-02 15:04:05"),
				amountStr,
				fmt.Sprintf("%d", tableNumber),
				fmt.Sprintf("%d", waiterID),
				waiterName,
			})
		}
		fmt.Println("Se ha exportado la boleta al archivo receipts.csv")
	}

}

func HandleService8() {
	service1 := &Service1{}
	service8 := &Service8{service1: service1}
	service8.Execute()
}
