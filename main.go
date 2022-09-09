package main

// tutorial: https://www.youtube.com/watch?v=d_L64KT3SFM
// rest api using gin framework

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "First Item", Completed: false},
	{ID: "2", Item: "Second Item", Completed: false},
	{ID: "3", Item: "Third Item", Completed: false},
}

func addTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(id string) (*todo, error) {
	for _, t := range todos {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}
func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)
	router.Run("localhost:8080")
}
