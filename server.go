package main

import (
    "net/http"
    "encoding/json"
)

type Photo struct {
    Name string
    Description string
}

func photoHandler(w http.ResponseWriter, r *http.Request) {
    photoEntry := Photo{Name: "girl's last tour", Description: "potato"}
    jsonBytes, err := json.Marshal(photoEntry)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
    }

    w.Header().Add("content-type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonBytes)
}

func routes() {
    http.HandleFunc("/photo", photoHandler);
}

func main() {
    routes()
    http.ListenAndServe(":8080", nil)
}
