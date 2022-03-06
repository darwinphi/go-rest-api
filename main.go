package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID string `json:"id"`
	Item string	`json:"item"`
	Completed bool `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Workout", Completed: false},
	{ID: "2", Item: "Clean room", Completed: false},
	{ID: "3", Item: "Wash dishes", Completed: false},
}

func getTodos(context	*gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
 var newTodo todo

 if err := context.BindJSON(&newTodo); err != nil {
	 return
 }

 todos = append(todos, newTodo)
 context.IndentedJSON(http.StatusOK, newTodo)
}

func getTodoById(id string) (*todo, error){
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9000")
}


// curl  --location --request POST 'http://localhost:9000/todos' \
//       --header 'Content-Type: application/json' \
//       --data '{"id": "4", "item": "Make bed", "completed": false}'

// curl  --location --request PATCH 'http://localhost:9000/todos/1' \
//       --header 'Content-Type: application/json'