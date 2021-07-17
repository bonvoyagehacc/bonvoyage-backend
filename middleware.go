package main

import (
    "fmt"
    "net/http"
)

/* allows only one type of method to be used on endpoint */
func RestrictMethod(method string, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != method {
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func RestrictAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        accessToken := r.Header.Get("x-access-token")

        username, err := ValidateToken(accessToken)
        if err != nil {
            fmt.Println(err)
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        fmt.Println(username)

        next.ServeHTTP(w, r)
    })
}
