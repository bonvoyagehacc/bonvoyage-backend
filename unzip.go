package main

import (
    "fmt"
    "bytes"
    "io"
    "archive/zip"
    "io/ioutil"
)

func Unzip(zipped io.ReadCloser) error {

    raw, err := ioutil.ReadAll(zipped)
    if err != nil { return err }

    reader, err := zip.NewReader(bytes.NewReader(raw), int64(len(raw)))
    if err != nil { return err }

    /* get list of files */
    for _, file := range reader.File {
        fmt.Println("Reading", file.Name)
    }

    return nil

}
