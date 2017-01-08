package modelo

import (
	"time"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Cliente struct {
	Id int
	Nombre string
	Apellido string
	Categoria string
	FechaNacimiento time.Time
}

var OpenDBvar = OpenDB()

func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/AlarmaSys?parseTime=True")

	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	return db
}

//funciones de BD y Open()
func (cliente Cliente) GetAll() []Cliente {
	// Execute the query
	rows, err := OpenDBvar.Query("SELECT id, nombre, apellido, categoria, fechaNacimiento FROM Cliente")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var clientes []Cliente
	var c Cliente
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(&c.Id, &c.Nombre, &c.Apellido, &c.Categoria, &c.FechaNacimiento)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		clientes = append(clientes, c)

	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println(clientes)
	return clientes
}


func (cliente Cliente) GetOne(id string) Cliente {
	// Execute the query
	rows := OpenDBvar.QueryRow("SELECT id, nombre, apellido, categoria, fechaNacimiento FROM Cliente WHERE id = ?", id)
	var c Cliente

	err := rows.Scan(&c.Id, &c.Nombre, &c.Apellido, &c.Categoria, &c.FechaNacimiento)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println(c)
	return c
}