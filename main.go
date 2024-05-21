package main

import (
    "go-api/controllers"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/books", controllers.GetBooks)
    r.GET("/books/:id", controllers.GetBook)
    r.POST("/books", controllers.CreateBook)
    r.PUT("/books/:id", controllers.UpdateBook)
    r.DELETE("/books/:id", controllers.DeleteBook)

    r.Run() // Inicia o servidor na porta 8080
}
