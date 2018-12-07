package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcmoe!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	todos := Todos{}
	db.Find(&todos)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "todoID")
	id, err := strconv.Atoi(idParam)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	todo := Todo{}
	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
	return
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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	db.Create(&todo)
	location := fmt.Sprintf("http://%s/%d", r.Host, todo.ID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
	return
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "todoID")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	todo := Todo{}
	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	db.Delete(&todo)

	w.WriteHeader(204) // 204 No Content
	return
}
