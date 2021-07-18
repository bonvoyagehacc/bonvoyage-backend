package main

import (
    "fmt"
	"bytes"
	"archive/zip"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"

	// "github.com/pixolous/pixolousAnalyze"
)

type galleryResponse struct {
	Images []string `json:"image"`
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	/* unzip received data */
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "malformed zip file", http.StatusBadRequest)
		return
	}
	reader, err := zip.NewReader(bytes.NewReader(raw), int64(len(raw)))
	if err != nil {
		http.Error(w, "malformed zip file", http.StatusBadRequest)
		return
	}

	/* get userid from context */
	userid := r.Context().Value("userid")
	if userid == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	/* get user's hash */
	userhash, err := GetUserHash(userid.(int))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	/* get list of files */
	for _, file := range reader.File {
		ext := filepath.Ext(file.Name)
		filehash := GenerateMD5(file.Name)

		/* write photo to drive */
		WriteImageFile(file, userhash, filehash+ext)

		/* compute ahash */
		// ahash := pixolousAnalyze.AHash(filepath.Join(ResourceDir, userhash, filehash+ext))
        ahash := "temp"

		/* write photo to db */
		NewPhoto(userid.(int), filehash+ext, ahash)
	}
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {

	/* get userid from context */
	userid := r.Context().Value("userid")
	if userid == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	/* get user's hash */
	userhash, err := GetUserHash(userid.(int))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

    filenames, err := GetUserPhotos(userid.(int))
    if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
    }

    imageURLs := []string{}
    for _, f := range filenames {
        imageURLs = append(imageURLs, SERVERURL+"/"+userhash+"/"+f)
    }
    fmt.Println(imageURLs)

	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(imageURLs)
}

func PhotoRoutes(mux *http.ServeMux) {
	mux.Handle("/photo/upload", RestrictMethod("POST", RestrictAuth(http.HandlerFunc(uploadHandler))))
	mux.Handle("/photo/gallery", RestrictMethod("GET", RestrictAuth(http.HandlerFunc(galleryHandler))))
}

