package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type employee struct {
	id                  int
	first_name          string
	last_name           string
	gender              string
	birth_year          int
	position            string
	employment_months   int
	salary              float64
	employment_duration string
}

func main() {
	fmt.Println("ds")
	connStr := "user=postgres dbname=employees sslmode=disable password=postgre host=localhost port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Введите SQL запрос: ")
	//zapr, _ := reader.ReadString('\n')
	//zapr = strings.TrimSpace(zapr)
	//zapr = strings.TrimSuffix(zapr, ";")
	rows, err := db.Query(//zapr)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	employees := []employee{}

	for rows.Next() {
		e := employee{}
		err := rows.Scan(&e.id, &e.first_name, &e.last_name, &e.gender, &e.birth_year, &e.position, &e.employment_months, &e.salary, &e.employment_duration)
		if err != nil {
			fmt.Println(err)
			continue
		}
		employees = append(employees, e)
	}
	for _, e := range employees {
		fmt.Println(e.id, e.first_name, e.last_name, e.salary)
	}
}
