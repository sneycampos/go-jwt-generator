package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

var (
	secret = "my-super-secret-jwt-key"
)

func main() {
	jwtToken := generateJWT()
	fmt.Println("Generated JWT:", jwtToken)

	fmt.Println("Validating JWT...")
	validateJWT(jwtToken)
}

func generateJWT() string {
	claims := &jwt.RegisteredClaims{
		Issuer:    "my-app",
		Subject:   "user-id",
		Audience:  []string{"audience"},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		NotBefore: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)),
		ID:        "unique-jwt-id",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return tokenString
}

func validateJWT(tokenString string) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if token.Valid {
		fmt.Println("Token is valid")
	} else {
		fmt.Println("Token is invalid")
	}

	fmt.Println("Claims:")
	fmt.Printf("Issuer: %s\n", claims.Issuer)
	fmt.Printf("Subject: %s\n", claims.Subject)
	fmt.Printf("Audience: %v\n", claims.Audience)
	fmt.Printf("Issued At: %s\n", claims.IssuedAt.Time)
	fmt.Printf("Expires At: %s\n", claims.ExpiresAt.Time)
	fmt.Printf("Not Before: %s\n", claims.NotBefore.Time)
	fmt.Printf("ID: %s\n", claims.ID)
}
