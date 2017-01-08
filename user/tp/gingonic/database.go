package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Cliente struct{
	id int
	nombre string
	apellido string
	categoria string
	fechaNacimiento time.Time
}

func main() {
	// Open database connection
	db, err := sql.Open("mysql", "root:root@/AlarmaSys?parseTime=True")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT id, nombre, apellido FROM cliente")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var c Cliente
	// Fetch rows
	for rows.Next() {
		err = rows.Scan(&c.id,&c.nombre,&c.apellido)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// Now do something with the data.
		// Here we just print each column as a string.
		fmt.Println(c)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
