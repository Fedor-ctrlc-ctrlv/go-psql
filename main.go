package main

import (
	"database/sql"
	"fmt"
	"log"

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
	for {
		mainMenu()
		choice := getInt("Выберите пункт меню: ")

		switch choice {
		case 0:
			fmt.Println("Выход из программы")
			return
		case 1:
			SearchMenu(db)
		case 2:
			DeleteMenu(db)
		case 3:
			UpdateMenu(db)
		default:
			fmt.Println("Неверный выбор, попробуйте снова")
		}
	}
}

func SearchMenu(db *sql.DB) {
	for {
		frontsearchMenu()
		choice := getInt("Выберите тип поиска: ")

		if choice == 0 {
			return
		}

		var query string
		if choice >= 1 && choice <= 4 {
			value := getInput("Введите значение для поиска: ")
			query = backSearchSelect(choice, value)
		} else if choice >= 5 && choice <= 7 {
			query = backSearchSelect(choice, "")
		} else {
			fmt.Println("Неверный выбор")
			continue
		}

		Printbd(db, query)
	}
}

func DeleteMenu(db *sql.DB) {
	for {
		frontdeleteMenu()
		choice := getInt("Выберите тип удаления: ")

		if choice == 0 {
			return
		}

		if choice < 1 || choice > 3 {
			fmt.Println("Неверный выбор")
			continue
		}

		value := getInput("Введите значение для удаления: ")
		query := backDelete(choice, value)

		result, err := db.Exec(query)
		if err != nil {
			fmt.Printf("Ошибка при удалении: %v\n", err)
			continue
		}

		rowsAffected, _ := result.RowsAffected()
		fmt.Printf("Удалено %d строк\n", rowsAffected)
	}
}

func UpdateMenu(db *sql.DB) {
	for {
		frontupdateMenu()
		choice := getInt("Выберите что изменить: ")

		if choice == 0 {
			return
		}

		if choice < 1 || choice > 4 {
			fmt.Println("Неверный выбор")
			continue
		}

		id := getInput("Введите ID сотрудника: ")
		value := getInput("Введите новое значение: ")
		query := backUpdate(choice, id, value)

		_, err := db.Exec(query)
		if err != nil {
			fmt.Printf("Ошибка при обновлении: %v\n", err)
			continue
		}

		fmt.Println("Данные успешно обновлены")
	}
}

func Printbd(db *sql.DB, query string) {
	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("Ошибка выполнения запроса: %v\n", err)
		return
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

	if len(employees) == 0 {
		fmt.Println("Данные не найдены")
	} else {
		fmt.Println("\n Результаты:")
		for _, e := range employees {
			fmt.Printf("ID: %d, Имя: %s, Фамилия: %s, Должность: %s, Зарплата: %.2f, В компании: %d мес.\n",
				e.id, e.first_name, e.last_name, e.position, e.salary, e.employment_months)
		}
	}
}
