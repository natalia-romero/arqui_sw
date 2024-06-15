package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func createDB(name string) *sql.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name)
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println("BDD creada")

	return db
}

func createTablesSQL(db *sql.DB, name string) {
	query := fmt.Sprintf("USE %s", name)
	_, err := db.Exec(query)

	if err != nil {
		panic(err)
	}

	createTablesSQL := []string{
		`CREATE TABLE IF NOT EXISTS user_status (id BIGINT AUTO_INCREMENT NOT NULL, name VARCHAR(255), PRIMARY KEY(id) );`,
		`CREATE TABLE IF NOT EXISTS type_user (id BIGINT AUTO_INCREMENT NOT NULL, name VARCHAR(255), description varchar(255), PRIMARY KEY(id) );`,
		`CREATE TABLE IF NOT EXISTS users (id BIGINT AUTO_INCREMENT NOT NULL, username VARCHAR(255), password VARCHAR(255), type_user_id BIGINT, status_id BIGINT, PRIMARY KEY(id), FOREIGN KEY(type_user_id) REFERENCES type_user(id), FOREIGN KEY(status_id) REFERENCES user_status(id) );`,
		`CREATE TABLE IF NOT EXISTS restaurant_table (id BIGINT AUTO_INCREMENT NOT NULL , number INT NOT NULL, user_id BIGINT, PRIMARY KEY(id), FOREIGN KEY(user_id) REFERENCES users(id) );`,
		`CREATE TABLE IF NOT EXISTS receipt (id BIGINT AUTO_INCREMENT NOT NULL, number INT NOT NULL, date DATE, amount INT, table_id BIGINT, PRIMARY KEY(id), FOREIGN KEY(table_id) REFERENCES restaurant_table(id) );`,
		`CREATE TABLE IF NOT EXISTS meals (id BIGINT AUTO_INCREMENT NOT NULL, name VARCHAR(255), description TEXT, price INT, PRIMARY KEY(id) );`,
		`CREATE TABLE IF NOT EXISTS receipt_detail (id BIGINT AUTO_INCREMENT NOT NULL, receipt_id BIGINT, meal_id BIGINT, quantity INT, PRIMARY KEY(id), FOREIGN KEY(receipt_id) REFERENCES receipt(id), FOREIGN KEY(meal_id) REFERENCES meals(id) );`}

	for _, query := range createTablesSQL {
		_, err = db.Exec(query)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Tablas creadas")
}

func addData(db *sql.DB, name string) {
	query := fmt.Sprintf("USE %s", name)
	_, err := db.Exec(query)

	if err != nil {
		panic(err)
	}

	data := make(map[string][]string)

	data["user_status"] = []string{
		`INSERT INTO user_status (id, name) VALUES (NULL, 'Activo')`,
		`INSERT INTO user_status (id, name) VALUES (NULL, 'Inactivo');`,
	}

	data["type_user"] = []string{
		`INSERT INTO type_user (id, name, description) VALUES (NULL, 'Garzon', 'Garzon o camarero del restaurante');`,
		`INSERT INTO type_user (id, name, description) VALUES (NULL, 'Administrador', 'Administrador del restaurante');`,
	}

	data["users"] = []string{
		`INSERT INTO users (id, username, password, type_user_id, status_id) VALUES (NULL, 'admin', 'admin', 2, 1);`,
		`INSERT INTO users (id, username, password, type_user_id, status_id) VALUES (NULL, 'felipe', 'felipe', 1, 1);`,
		`INSERT INTO users (id, username, password, type_user_id, status_id) VALUES (NULL, 'maxi', 'maxi', 1, 1);`,
		`INSERT INTO users (id, username, password, type_user_id, status_id) VALUES (NULL, 'natalia', 'natalia', 1, 1);`,
		`INSERT INTO users (id, username, password, type_user_id, status_id) VALUES (NULL, 'benjamin', 'benjamin', 1, 1);`,
		`INSERT INTO users (id, username, password, type_user_id, status_id) VALUES (NULL, 'nicolas', 'nicolas', 1, 1);`,
	}

	data["restaurant_table"] = []string{
		`INSERT INTO restaurant_table (id, number, user_id) VALUES (NULL, 1, 2);`,
		`INSERT INTO restaurant_table (id, number, user_id) VALUES (NULL, 1, 3);`,
		`INSERT INTO restaurant_table (id, number, user_id) VALUES (NULL, 2, 4);`,
		`INSERT INTO restaurant_table (id, number, user_id) VALUES (NULL, 3, 5);`,
		`INSERT INTO restaurant_table (id, number, user_id) VALUES (NULL, 4, 6);`,
	}

	data["meals"] = []string{
		`INSERT INTO meals (id, name, description, price) VALUES (NULL, 'Milanesa', 'Milanesa con papas fritas', 5000);`,
		`INSERT INTO meals (id, name, description, price) VALUES (NULL, 'Pizza', 'Pizza todas las carnes', 7000);`,
		`INSERT INTO meals (id, name, description, price) VALUES (NULL, 'Hamburguesa', 'Hamburguesa con papas fritas', 5500);`,
		`INSERT INTO meals (id, name, description, price) VALUES (NULL, 'Ensalada', 'Ensalada cesar', 3500);`,
		`INSERT INTO meals (id, name, description, price) VALUES (NULL, 'Pasta', 'Pasta con salsa bolognesa', 6000);`,
	}

	for key, values := range data {
		for _, value := range values {
			_, err := db.Exec(value)
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("Datos añadidos a:", key)
	}

}

func main() {
	dBName := "bdd_golang"
	db := createDB(dBName)

	createTablesSQL(db, dBName)
	addData(db, dBName)
}
