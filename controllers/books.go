package main

import (
	"github.com/gin-gonic/gin"

	"bookstore/controllers" // new
	"bookstore/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks) // new

	r.Run()
}
