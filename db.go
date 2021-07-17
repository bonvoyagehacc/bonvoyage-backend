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

func main() {

    connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", ConnectionSecrets.Host, ConnectionSecrets.Port, ConnectionSecrets.User, ConnectionSecrets.Password, ConnectionSecrets.DBName)

    /* validat connection */
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Connected to database")
}

