package main

import (
    "net/http"
    "database/sql"
    "encoding/json"
    "gopkg.in/validator.v2"
)

type credRequest struct {
    Username string `json:"username" validate:"min=3,max=64,nonzero"`
    Password string `json:"password" validate:"min=8,max=128,nonzero"`
}

type tokenResponse struct {
    AccessToken string `json:"accessToken"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

    var creds credRequest
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, "malformed request", http.StatusBadRequest); return
    }
    if err := validator.Validate(creds); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest); return
    }

    /* query db */
    id, err := RegisterUser(creds.Username, creds.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest); return
    }

    /* return access token */
    tokenString, err := GenerateToken(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError); return
    }
    token := tokenResponse{AccessToken: tokenString}

    w.Header().Add("content-type", "application/json")
    json.NewEncoder(w).Encode(token)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

    var creds credRequest
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, "malformed request", http.StatusBadRequest); return
    }
    if err := validator.Validate(creds); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest); return
    }

    /* query db */
    id, err := AuthenticateUser(creds.Username, creds.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            w.WriteHeader(http.StatusUnauthorized)
        } else {
            w.WriteHeader(http.StatusInternalServerError)
        }
        return
    }

    /* return token */
    tokenString, err := GenerateToken(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError); return
    }
    token := tokenResponse{AccessToken: tokenString}

    w.Header().Add("content-type", "application/json")
    json.NewEncoder(w).Encode(token)
}

func AuthRoutes(mux *http.ServeMux) {
    mux.Handle("/auth/register", RestrictMethod("POST", http.HandlerFunc(registerHandler)))
    mux.Handle("/auth/login", RestrictMethod("POST", http.HandlerFunc(loginHandler)))
}

