package main

import (
    "net/http"
)

func main() {
    port := ":8080"
    AuthRoutes()
    http.ListenAndServe(port, nil)
}

