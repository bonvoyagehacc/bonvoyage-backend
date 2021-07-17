package main

import (
    "net/http"
    "encoding/json"
    "fmt"
)

type credRequest struct {
    Username string
    Password string
}

type tokenResponse struct {
    AccessToken string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    var creds credRequest
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Println(creds)

    /* return access token */
    w.Header().Add("content-type", "application/json")
    token := tokenResponse{AccessToken: "peepeepoopoo"}
    json.NewEncoder(w).Encode(token)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }
}

func AuthRoutes() {
    // http.HandleFunc("/photo", photoHandler);
    http.HandleFunc("/auth/login", loginHandler);
    http.HandleFunc("/auth/register", registerHandler);
}

