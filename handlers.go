package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	u "./utils"

	"github.com/go-chi/chi"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcmoe!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{}
	db.Find(&todos)
	u.Respond(w, http.StatusOK, todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todo := Todo{}
	if err := db.Where("id = ?", todoID).First(&todo).Error; err != nil {
		msg := fmt.Sprintf("Couldn't find a todo with ID %s", todoID)
		u.Respond(w, http.StatusNotFound, u.Message(false, msg))
		return
	}
	u.Respond(w, http.StatusOK, todo)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // 1MiB
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// Unmarshal make todo with decoded json
	if err := json.Unmarshal(body, &todo); err != nil {
		msg := fmt.Sprintf("Initernal Server Error")
		u.Respond(w, http.StatusInternalServerError, u.Message(false, msg))
		return
	}
	db.Create(&todo)
	location := fmt.Sprintf("http://%s/%d", r.Host, todo.ID)
	w.Header().Set("Location", location)
	u.Respond(w, http.StatusCreated, todo)
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todo := Todo{}
	if err := db.Where("id = ?", todoID).First(&todo).Error; err != nil {
		msg := fmt.Sprintf("Couldn't find a todo with ID %s", todoID)
		u.Respond(w, http.StatusNotFound, u.Message(false, msg))
		return
	}
	db.Delete(&todo)
	u.Respond(w, http.StatusOK, u.Message(true, "Success to delete"))
	return
}
