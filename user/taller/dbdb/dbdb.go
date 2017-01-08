package main

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/user/taller/modelo"
)

func main() {
	db, err := sql.Open("mysql", "root:jsf732@/AlarmaSys")

	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT id, nombre, apellido FROM Cliente")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var clientes []modelo.Cliente
	var c modelo.Cliente
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(&c.Id, &c.Nombre, &c.Apellido)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		clientes = append(clientes, c)
		fmt.Println(c)

	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}


}