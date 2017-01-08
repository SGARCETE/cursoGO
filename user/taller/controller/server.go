package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
	"github.com/user/taller/modelo"
)

func getAll(c *gin.Context) {
	var cliente modelo.Cliente
	var clientes []modelo.Cliente = cliente.GetAll()
	c.JSON(200, gin.H{
		"clientes": clientes,
	})
}

func getOne(c *gin.Context) {
	id := c.Query("id")
	var cliente modelo.Cliente
	var clienteConsulta modelo.Cliente = cliente.GetOne(id)
	c.JSON(200, gin.H{
		"clientes": clienteConsulta,
	})
}

func main() {
	r := gin.Default()
	r.GET("/clientes", getAll )
	r.GET("/cliente", getOne )
	r.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	r.Run() // listen and server on 0.0.0.0:8080
	defer modelo.OpenDBvar.Close()
}

