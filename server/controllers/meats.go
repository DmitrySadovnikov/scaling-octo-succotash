package controllers

import (
    "net/http"
    "../services"
    "github.com/gorilla/mux"
    "encoding/json"
)

type streamType string

func Meat(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    service := services.MeatService{}
    result := service.Call(id)
    js, err := json.Marshal(result)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write([]byte(js))
}

func Meats(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello Meats"))
}
