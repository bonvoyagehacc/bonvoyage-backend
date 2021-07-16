package main

import (
    "net/http"
)

func main() {
    Routes()
    http.ListenAndServe(":8080", nil)
}

