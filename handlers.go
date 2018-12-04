package main

import (
  "fmt"
  "net/http"
  "time"
  "encoding/json"

  "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Fprintf(w, "Welcmoe!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  todos := Todos{
    Todo{Name: "Write presentation", Due: time.Now()},
    Todo{Name: "Host meetup"},
  }

  if err := json.NewEncoder(w).Encode(todos); err != nil {
    panic(err)
  }
}

func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "Todo show: %s", ps.ByName("todoId"))
}
