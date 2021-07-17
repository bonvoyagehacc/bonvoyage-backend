package main

import (
    "bytes"
    "path/filepath"
    "archive/zip"
    "io/ioutil"
    "net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {

    /* unzip received data */
    raw, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "malformed zip file", http.StatusBadRequest); return
    }
    reader, err := zip.NewReader(bytes.NewReader(raw), int64(len(raw)))
    if err != nil {
        http.Error(w, "malformed zip file", http.StatusBadRequest); return
    }

    /* get userid from context */
    userid := r.Context().Value("userid")
    if userid == nil {
        w.WriteHeader(http.StatusInternalServerError); return
    }
    /* get user's hash */
    userhash, err := GetUserHash(userid.(int))
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError); return
    }

    /* get list of files */
    for _, file := range reader.File {
        ext := filepath.Ext(file.Name)
        filehash := GenerateMD5(file.Name)

        /* write photo to db - add a random string at end to take care of duplicate filenames */
        NewPhoto(userid.(int), filehash+ext)
        /* write photo to drive */
        WriteImageFile(file, userhash, filehash+ext)
    }
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
}

func PhotoRoutes(mux *http.ServeMux) {
    mux.Handle("/photo/upload", RestrictMethod("POST", RestrictAuth(http.HandlerFunc(uploadHandler))));
    mux.Handle("/photo/gallery", RestrictMethod("GET", RestrictAuth(http.HandlerFunc(galleryHandler))));
}

