package main

import (
	"log"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main() {
	fmt.Println("Hello World")

	r := gin.Default()

	todos := []Todo{}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "hello world"})
	})

	r.POST("/api/todos", func (c *gin.Context) {
		todo := &Todo{}

		if err := c.ShouldBindJSON(todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return 
		}
		
		if todo.Body == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Todo body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		c.JSON(http.StatusCreated, todo)
	})

	log.Fatal(r.Run(":4000"))
}