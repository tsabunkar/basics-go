package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	// HOST     = "172.17.0.2"
	HOST 	 = "127.0.0.2"
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
	// fmt.Println("db", db)
	if err != nil {
		fmt.Println("failed while connecting to the db")
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	// Perform a sample query
	rows, err := db.Query("SELECT 1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	

	 // Iterate through the results
	 for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
}
