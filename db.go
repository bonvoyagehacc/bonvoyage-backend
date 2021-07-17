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
//     username varchar(64) not null,
//     password varchar(128) not null
// );

type PostgresDB struct {
    db *sql.DB
}

func (p *PostgresDB) RegisterUser(username string, password string) {
    db := p.db

    query := `
    INSERT INTO users (username, password)
    VALUES ($1, $2)
    RETURNING id
    `
    _, err := db.Exec(query, username, password)
    if err != nil {
        panic(err)
    }
    return
}

func (p *PostgresDB) AuthenticateUser(username string, password string) {
    db := p.db

    query := `
    SELECT id FROM users WHERE username = $1 AND password = $2
    `
    _, err := db.Exec(query, username, password)
    if err != nil {
        panic(err)
    }
    return
}

func main() {

    connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", ConnectionSecrets.Host, ConnectionSecrets.Port, ConnectionSecrets.User, ConnectionSecrets.Password, ConnectionSecrets.DBName)

    /* validate connection */
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    psdb := &PostgresDB{db: db}
    psdb.RegisterUser("pinosaur", "peepee")

    fmt.Println("Connected to database")
}

