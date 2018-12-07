package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

var db *gorm.DB

func main() {
	router := httprouter.New()
	router.GET("/", Logging(Index, "index"))
	router.GET("/todos", Logging(TodoIndex, "todo-index"))
	router.POST("/todos", Logging(TodoCreate, "todo-create"))
	router.GET("/todos/:todoId", Logging(TodoShow, "todo-show"))
	router.DELETE("/todos/:todoId", Logging(TodoDelete, "todo-delete"))

	// Setup DB
	var err error
	db, err = setupDB()
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	migrateDB()
	defer db.Close()

	// logs
	db.LogMode(true)
	log.Fatal(http.ListenAndServe(":8080", router))
}
