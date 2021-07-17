package main

import (
    "time"
    "github.com/golang-jwt/jwt"
)

type Claims struct {
    Username string
    jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {

    expire := time.Now().Add(time.Duration(JWTLifetime) * time.Minute)
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expire.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(JWTSecret)

    return tokenString, err
}
