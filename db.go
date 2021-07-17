package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

type PostgresConnection struct {
    Host string
    Port int
    User string
    Password string
    DBName string
}

// CREATE TABLE users(
//     id serial primary key,
//     username varchar(64) not null unique,
//     password varchar(128) not null
// );

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

func RegisterUser(username string, password string) error {
    db := dbConnection()
    defer db.Close()

    query := `
    INSERT INTO users (username, password)
    VALUES ($1, $2)
    `
    _, err := db.Exec(query, username, password)
    return err
}

func AuthenticateUser(username string, password string) error {
    db := dbConnection()
    defer db.Close()

    query := `
    SELECT id FROM users WHERE username = $1 AND password = $2
    `
    id := 0
    err := db.QueryRow(query, username, password).Scan(&id)
    return err
}

