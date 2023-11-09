package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)

}
func AddTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

var todos = []todo{
	{ID: "1", Item: "item1", Completed: false},
	{ID: "2", Item: "item2", Completed: false},
	{ID: "3", Item: "item3", Completed: false},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", AddTodo)
	router.Run("localhost:9090")

}
