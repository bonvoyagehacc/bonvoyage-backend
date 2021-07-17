package main

import (
    "net/http"
    "encoding/json"
    "gopkg.in/validator.v2"
)

type uploadRequest struct {
    Zipped string `json:"zipped"`
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

    var pics uploadRequest
    if err := json.NewDecoder(r.Body).Decode(&pics); err != nil {
        http.Error(w, "malformed request", http.StatusBadRequest); return
    }
    if err := validator.Validate(pics); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest); return
    }

    /* unzip received data */

}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
}

func PhotoRoutes(mux *http.ServeMux) {
    mux.Handle("/photo/upload", RestrictMethod("POST", RestrictAuth(http.HandlerFunc(uploadHandler))));
    mux.Handle("/photo/gallery", RestrictMethod("GET", RestrictAuth(http.HandlerFunc(galleryHandler))));
}
