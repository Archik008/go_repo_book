package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) != 6 {
		fmt.Println("Please provide: hostname port username password db")
		return
	}
	host := arguments[1]
	p := arguments[2]
	user := arguments[3]
	pass := arguments[4]
	database := arguments[5]
	// номер порта ДОЛЖЕН БЫТЬ целым числом
	// Port number SHOULD BE an integer
	port, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Not a valid port number:", err)
		return
	}
	// строка подключения
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, database)

	// открыть базу данных PostgreSQL
	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("Open():", err)
		return
	}
	defer db.Close()

	// получить все базы данных
	rows, err := db.Query(`SELECT "datname" FROM "pg_database" WHERE datistemplate = false`)
	if err != nil {
		fmt.Println("Query", err)
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan", err)
			return
		}
		fmt.Println("*", name)
	}
	defer rows.Close()

	// получить все таблицы из текущей базы данных
	query := `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' ORDER BY table_name`
	rows, err = db.Query(query)
	if err != nil {
		fmt.Println("Query", err)
		return
	}

	// вот как вы обрабатываете строки, возвращаемые из SELECT
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan", err)
			return
		}
		fmt.Println("+T", name)
	}
	defer rows.Close()
}
