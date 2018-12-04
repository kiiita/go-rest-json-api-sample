package main

import (
  "fmt"
  "net/http"
  "encoding/json"

  "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Fprintf(w, "Welcmoe!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(w).Encode(todos); err != nil {
    panic(err)
  }
}

func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "Todo show: %s", ps.ByName("todoId"))
}
