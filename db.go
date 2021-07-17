package main

import (
    "fmt"
    "io"
    "encoding/hex"
    "database/sql"
    "crypto/md5"
    _ "github.com/lib/pq"
)

type PostgresConnection struct {
    Host string
    Port int
    User string
    Password string
    DBName string
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
//     password varchar(128) not null
// );

func RegisterUser(username string, password string) (int, error) {
    db := dbConnection()
    defer db.Close()

    query := `
    INSERT INTO users (username, password)
    VALUES ($1, $2)
    RETURNING id
    `
    id := 0
    err := db.QueryRow(query, username, password).Scan(&id)
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

/* =-=-=-=-=-=-= PHOTOS =-=-=-=-=-=-=-=*/
func GenerateMD5(raw string) string {
    hasher := md5.New()
    io.WriteString(hasher, raw)
    return hex.EncodeToString(hasher.Sum(nil)[:])
}

func NewPhoto(userid int, filename string) {
    db := dbConnection()
    defer db.Close()

    /* grab extension */

    /* hash filename */
    filehash := GenerateMD5(filename)
    fmt.Println(filehash)

    /* insert into db */

}
