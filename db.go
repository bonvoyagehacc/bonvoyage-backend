package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"archive/zip"
	"crypto/md5"
	"database/sql"
	"encoding/hex"

	_ "github.com/lib/pq"
)

type PostgresConnection struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func dbConnection() *sql.DB {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", ConnectionSecrets.Host, ConnectionSecrets.Port, ConnectionSecrets.User, ConnectionSecrets.Password, ConnectionSecrets.DBName)

	/* validate connection */
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

/* =-=-=-=-=-=-= USERS =-=-=-=-=-=-=-=*/
// CREATE TABLE users(
//     id serial primary key,
//     username varchar(64) not null unique,
//     password varchar(128) not null,
//     hash     varchar(32) not null
// );

func RegisterUser(username string, password string) (int, error) {
	db := dbConnection()
	defer db.Close()

	query := `
    INSERT INTO users (username, password, hash)
    VALUES ($1, $2, $3)
    RETURNING id
    `
	id := 0
	err := db.QueryRow(query, username, password, GenerateMD5(username)).Scan(&id)
	return id, err
}

func AuthenticateUser(username string, password string) (int, error) {
	db := dbConnection()
	defer db.Close()

	query := `
    SELECT id FROM users WHERE username = $1 AND password = $2
    `
	id := 0
	err := db.QueryRow(query, username, password).Scan(&id)
	return id, err
}

func GetUserHash(id int) (string, error) {
	db := dbConnection()
	defer db.Close()

	query := `
    SELECT hash FROM users WHERE id = $1
    `
	hash := ""
	err := db.QueryRow(query, id).Scan(&hash)
	return hash, err
}

/* =-=-=-=-=-=-= PHOTOS =-=-=-=-=-=-=-=*/
// CREATE TABLE photos (
//     id serial primary key,
//     userid int not null,
//     filename varchar(64) not null,
//     CONSTRAINT fk_user FOREIGN KEY(userid) REFERENCES users(id)
// );

func GenerateMD5(raw string) string {
	hasher := md5.New()
	io.WriteString(hasher, raw)
	return hex.EncodeToString(hasher.Sum(nil)[:])
}

func WriteImageFile(file *zip.File, userhash string, filename string) error {

	/* create parent dir if it doesnt already exist */
	if err := os.MkdirAll(filepath.Join(ResourceDir, userhash), os.ModePerm); err != nil {
		return err
	}

	/* copy unziped file to disk */
	newFile, err := os.OpenFile(filepath.Join(ResourceDir, userhash, filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer newFile.Close()

	handle, err := file.Open()
	if err != nil {
		return err
	}
	defer handle.Close()

	_, err = io.Copy(newFile, handle)

	return nil

}

func NewPhoto(userid int, filename string, ahash string) error {
	db := dbConnection()
	defer db.Close()

	/* insert into db */
	query := `
    INSERT INTO photos (userid, filename, ahash)
    VALUES ($1, $2, $3)
    RETURNING id
    `
	id := 0
	err := db.QueryRow(query, userid, filename, ahash).Scan(&id)
	return err
}

func GetUserPhotos(userid int) ([]string, error) {
	db := dbConnection()
	defer db.Close()

	query := `
    SELECT filename FROM photos WHERE userid = $1
    `
    rows, err := db.Query(query, userid)
    if err != nil {
        return []string{}, err
    }
    defer rows.Close()

    filenames := []string{}
    for rows.Next() {
        filename := ""
        rows.Scan(&filename)
        filenames = append(filenames, filename)
    }

    return filenames, nil
}

