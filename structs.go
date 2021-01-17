package main

import jwt "github.com/dgrijalva/jwt-go"

type ErrorResponse struct {
	Code    int
	Message string
}

type SuccessResponse struct {
	Code     int
	Message  string
	Response interface{}
}

type Claims struct {
	User_id string
	Name    string
	Email   string
	jwt.StandardClaims
}

type RegistrationParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SuccessfulLoginResponse struct {
	Email     string
	AuthToken string
}

type UserDetails struct {
	User_id  string
	Name     string
	Email    string
	Password string
}
