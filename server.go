package main

import (
    "net/http"
    "github.com/rs/cors"
)

func main() {
    port := ":8080"

    var mux *http.ServeMux = http.NewServeMux()
    AuthRoutes(mux)
    PhotoRoutes(mux)

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"https://pixolo.us", "http://localhost:3000"},
        AllowCredentials: true,
    })
    handler := c.Handler(mux)
    http.ListenAndServe(port, handler)
}

