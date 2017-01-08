package main

import (
	"time"
)

type Cliente struct{
	id int
	nombre string
	apellido string
	categoria string
	fechaNacimiento time.Time
}