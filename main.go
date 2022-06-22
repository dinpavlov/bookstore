package main

import (
	"github.com/gin-gonic/gin"

	"bookstore/controllers" // new
	"bookstore/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)         // new
	r.POST("/books", controllers.CreateBook)       // new
	r.GET("/books/:id", controllers.FindBook)      // new
	r.PATCH("/books/:id", controllers.UpdateBook)  // new
	r.DELETE("/books/:id", controllers.DeleteBook) // new

	r.Run()
}
