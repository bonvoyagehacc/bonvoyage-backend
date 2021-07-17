package main

import (
    "net/http"
    "database/sql"
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

func registerHandler(w http.ResponseWriter, r *http.Request) {
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
    tokenString, err := GenerateToken(creds.Username)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    token := tokenResponse{AccessToken: tokenString}

    w.Header().Add("content-type", "application/json")
    json.NewEncoder(w).Encode(token)
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
    if err := AuthenticateUser(creds.Username, creds.Password); err != nil {
        if err == sql.ErrNoRows {
            w.WriteHeader(http.StatusUnauthorized)
        } else {
            w.WriteHeader(http.StatusInternalServerError)
        }
        return
    }

    /* return token */
    tokenString, err := GenerateToken(creds.Username)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    token := tokenResponse{AccessToken: tokenString}

    w.Header().Add("content-type", "application/json")
    json.NewEncoder(w).Encode(token)
}

func AuthRoutes() {
    http.HandleFunc("/auth/register", registerHandler);
    http.HandleFunc("/auth/login", loginHandler);
}

