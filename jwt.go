package main

import (
    "errors"
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

func ValidateToken(tokenString string) (string, error) {

    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return JWTSecret, nil
    })
    if err != nil {
        return "", err
    }
    if !token.Valid {
        return "", errors.New("invalid access token")
    }
    return claims.Username, nil
}

