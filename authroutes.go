package main

import (
    "net/http"
    "encoding/json"
    "gopkg.in/validator.v2"
)

type credRequest struct {
    Username string `validate:"min=3,max=64,nonzero"`
    Password string `validate:"min=8,max=128,nonzero"`
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
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    /* validate query */
    if err := validator.Validate(creds); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    /* query db */
    if err := RegisterUser(creds.Username, creds.Password); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

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
    http.HandleFunc("/auth/login", loginHandler);
    http.HandleFunc("/auth/register", registerHandler);
}

