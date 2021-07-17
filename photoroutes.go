package main

import (
    "fmt"
    "net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {

    /* unzip received data */
    err := Unzip(r.Body)
    if err != nil {
        http.Error(w, "malformed zip file", http.StatusBadRequest); return
    }

}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
}

func PhotoRoutes(mux *http.ServeMux) {
    mux.Handle("/photo/upload", RestrictMethod("POST", RestrictAuth(http.HandlerFunc(uploadHandler))));
    mux.Handle("/photo/gallery", RestrictMethod("GET", RestrictAuth(http.HandlerFunc(galleryHandler))));
}
