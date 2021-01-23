package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = []byte("jwt_secret_key_rumeh")

//CreateJWT to generate Token with payload name and email
func CreateJWT(name string, email string, userid int) (response string, err error) {
	expirationTime := time.Now().Add(100 * time.Minute)
	claims := &Claims{
		Name:   name,
		Email:  email,
		Userid: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err == nil {
		return tokenString, nil
	}
	return "", err
}

// VerifyToken to Verify the accepted token from http request
func VerifyToken(tokenString string) (email string, err error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if token != nil {
		return claims.Email, nil
	}
	return "", err
}
