package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type employee struct {
	id                int
	first_name        string
	last_name         string
	gender            string
	birth_year        int
	position          string
	employment_months int
	salary            float64
}

func main() {
	fmt.Println("ds")
	connStr := "user=postgres dbname=employees sslmode=disable password=postgre host=localhost port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
