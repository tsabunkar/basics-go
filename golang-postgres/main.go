package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST     = "172.17.0.2"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "root"
	DBNAME   = "postgres"
)

func main() {
	fmt.Println("Starting the db connection")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME)
	fmt.Println("psqlInfo", psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("failed while connecting to the db")
		log.Fatal(err)
	}
	defer db.Close()


	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
}
