package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Logging(Index, "index"))
	router.GET("/todos", Logging(TodoIndex, "todo-index"))
	router.POST("/todos", Logging(TodoCreate, "todo-create"))
	router.GET("/todos/:todoId", Logging(TodoShow, "todo-show"))
	router.DELETE("/todos/:todoId", Logging(TodoDelete, "todo-delete"))

	log.Fatal(http.ListenAndServe(":8080", router))
}
