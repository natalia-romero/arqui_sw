package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

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

func HandleService1() {
	service1 := &Service1{}
	service1.Execute()

}

func soaConnection1() {
	serviceConnection("serv1")
}
