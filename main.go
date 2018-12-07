package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", Index)
	r.Get("/todos", TodoIndex)
	r.Post("/todos", TodoCreate)
	r.Get("/todos/{todoID}", TodoShow)
	r.Delete("/todos/{todoID}", TodoDelete)

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
	log.Fatal(http.ListenAndServe(":8080", r))
}
