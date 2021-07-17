package main

import (
    "net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" { w.WriteHeader(http.StatusMethodNotAllowed); return }


}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" { w.WriteHeader(http.StatusMethodNotAllowed); return }

}

func PhotoRoutes(mux *http.ServeMux) {
    http.HandleFunc("/photo/upload", uploadHandler);
    http.HandleFunc("/photo/gallery", galleryHandler);
}
