package main

import "fmt"

func mainMenu() {
	fmt.Println("\n Главное меню")
	fmt.Println("1. Поиск в базе данных")
	fmt.Println("2. Удаление данных")
	fmt.Println("3. Изменение данных")
	fmt.Println("0. Выход")
}

func searchMenu() {
	fmt.Println("\n Меню поиска:")
	fmt.Println("1. Поиск по ID")
	fmt.Println("2. Поиск по имени")
	fmt.Println("3. Поиск по фамилии")
	fmt.Println("4. Поиск по должности")
	fmt.Println("5. Сортировка по зарплате")
	fmt.Println("6. Сортировка по фамилии (алфавит)")
	fmt.Println("7. Сортировка по времени в компании")
	fmt.Println("0. Назад")
}

func deleteMenu() {
	fmt.Println("\n Меню удаления:")
	fmt.Println("1. Удалить по ID")
	fmt.Println("2. Удалить по имени")
	fmt.Println("3. Удалить по фамилии")
	fmt.Println("0. Назад")
}

func updateMenu() {
	fmt.Println("\n Меню изменения:")
	fmt.Println("1. Изменить зарплату")
	fmt.Println("2. Изменить имя")
	fmt.Println("3. Изменить фамилию")
	fmt.Println("4. Изменить должность")
	fmt.Println("0. Назад")
}

func backSearchSelect(choice int, value string) string {
	switch choice {
	case 1:
		return fmt.Sprintf("SELECT * FROM employees WHERE id = %s", value)
	case 2:
		return fmt.Sprintf("SELECT * FROM employees WHERE first_name = '%s'", value)
	case 3:
		return fmt.Sprintf("SELECT * FROM employees WHERE last_name = '%s'", value)
	case 4:
		return fmt.Sprintf("SELECT * FROM employees WHERE position = '%s'", value)
	case 5:
		return "SELECT * FROM employees ORDER BY salary DESC"
	case 6:
		return "SELECT * FROM employees ORDER BY last_name ASC"
	case 7:
		return "SELECT * FROM employees ORDER BY employment_months DESC"
	default:
		return ""
	}
}

func backDelete(choice int, value string) string {
	switch choice {
	case 1:
		return fmt.Sprintf("DELETE FROM employees WHERE id = %s", value)
	case 2:
		return fmt.Sprintf("DELETE FROM employees WHERE first_name = '%s'", value)
	case 3:
		return fmt.Sprintf("DELETE FROM employees WHERE last_name = '%s'", value)
	default:
		return ""
	}
}

func backUpdate(choice int, id, value string) string {
	switch choice {
	case 1:
		return fmt.Sprintf("UPDATE employees SET salary = %s WHERE id = %s", value, id)
	case 2:
		return fmt.Sprintf("UPDATE employees SET first_name = '%s' WHERE id = %s", value, id)
	case 3:
		return fmt.Sprintf("UPDATE employees SET last_name = '%s' WHERE id = %s", value, id)
	case 4:
		return fmt.Sprintf("UPDATE employees SET position = '%s' WHERE id = %s", value, id)
	default:
		return ""
	}
}
