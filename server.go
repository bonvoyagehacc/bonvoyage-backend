package main

import (
    "fmt"
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
        AllowedHeaders: []string{"x-access-token", "Content-Type"},
        AllowCredentials: true,
        // Debug: true,
    })
    handler := c.Handler(mux)

    fmt.Printf("Listening on port %s\n", port)
    http.ListenAndServe(port, handler)
}

