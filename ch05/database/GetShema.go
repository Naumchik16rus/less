package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	argumens := os.Args
	if len(argumens) != 6 {
		fmt.Println("Please provide: hostname port username password db")
		return
	}
	host := argumens[1]
	port := argumens[2]
	user := argumens[3]
	pass := argumens[4]
	database := argumens[5]

	//номер порта должен быть целым числом!!!!
	/*if err != nil {
		fmt.Println("Not a valid port number:", err)
		return
	}*/
	//строка подключения
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, database)
	//откыть базу PostgreSQL
	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("Open():", err)
		return
	}
	defer db.Close()
	//получить все базы данных
	rows, err := db.Query(`SELECT "datname" FROM "pg_database" WHERE datistemplate = false`)
	if err != nil {
		fmt.Println("Query", err)
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
	//получить все таблицы из текущей базы данных
	query := `SELECT table_name FROM information_schema.tables WHERE
	table_schema = 'public' ORDER BY table_name`
	rows, err = db.Query(query)
	if err != nil {
		fmt.Println("Query", err)
		return
	}
	//воткак вы обрабатываете строки.возвращаемые из Select
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
