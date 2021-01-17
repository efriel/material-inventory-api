package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func RenderHome(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "views/profile.html")
}

func RenderLogin(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "views/login.html")
}

func RenderRegister(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "views/register.html")
}

func SignInUser(response http.ResponseWriter, request *http.Request) {
	var loginRequest LoginParams
	var result UserDetails
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	decoder := json.NewDecoder(request.Body)
	decoderErr := decoder.Decode(&loginRequest)

	defer request.Body.Close()

	if decoderErr != nil {
		returnErrorResponse(response, request, errorResponse)
	} else {
		errorResponse.Code = http.StatusBadRequest

		if loginRequest.Email == "" {
			errorResponse.Message = "Email can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else if loginRequest.Password == "" {
			errorResponse.Message = "Password can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else {

			collection := Client.Database("msdb").Collection("users")

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var err = collection.FindOne(ctx, bson.M{
				"email": loginRequest.Email,
			}).Decode(&result)
			if err != nil {
				errorResponse.Message = "User not found"
				returnErrorResponse(response, request, errorResponse)
			}

			defer cancel()

			if err != nil {
				returnErrorResponse(response, request, errorResponse)
			} else {
				pwdmatch := CheckPasswordHash(loginRequest.Password, result.Password)
				if pwdmatch != true {
					errorResponse.Message = "Password not match"
					returnErrorResponse(response, request, errorResponse)
				} else {
					tokenString, _ := CreateJWT(result.Name, loginRequest.Email)

					if tokenString == "" {
						returnErrorResponse(response, request, errorResponse)
					}

					var successResponse = SuccessResponse{
						Code:    http.StatusOK,
						Message: "You are registered, login again",
						Response: SuccessfulLoginResponse{
							AuthToken: tokenString,
							Email:     loginRequest.Email,
							Name:      result.Name,
						},
					}

					successJSONResponse, jsonError := json.Marshal(successResponse)

					if jsonError != nil {
						returnErrorResponse(response, request, errorResponse)
					}
					response.Header().Set("Content-Type", "application/json")
					response.Write(successJSONResponse)
				}
			}
		}
	}
}

func SignUpUser(response http.ResponseWriter, request *http.Request) {
	var registrationRequest RegistrationParams
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	decoder := json.NewDecoder(request.Body)
	decoderErr := decoder.Decode(&registrationRequest)
	defer request.Body.Close()

	if decoderErr != nil {
		returnErrorResponse(response, request, errorResponse)
	} else {
		errorResponse.Code = http.StatusBadRequest
		fmt.Printf(registrationRequest.Email)
		if registrationRequest.Name == "" {
			errorResponse.Message = "Name can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else if registrationRequest.Email == "" {
			errorResponse.Message = "Email can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else if registrationRequest.Password == "" {
			errorResponse.Message = "Password can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else {
			tnow := time.Now()
			tsec := tnow.Unix()
			ntsec := strconv.FormatInt(tsec, 10)

			tokenString, _ := CreateJWT(registrationRequest.Name, registrationRequest.Email)

			if tokenString == "" {
				returnErrorResponse(response, request, errorResponse)
			}

			var registrationResponse = SuccessfulLoginResponse{
				AuthToken: tokenString,
				Email:     registrationRequest.Email,
				Name:      registrationRequest.Name,
			}

			collection := Client.Database("msdb").Collection("users")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, databaseErr := collection.InsertOne(ctx, bson.M{
				"user_id":  ntsec,
				"email":    registrationRequest.Email,
				"password": getHash([]byte(registrationRequest.Password)),
				"name":     registrationRequest.Name,
			})
			defer cancel()

			if databaseErr != nil {
				returnErrorResponse(response, request, errorResponse)
			}

			var successResponse = SuccessResponse{
				Code:     http.StatusOK,
				Message:  "Successfully registered, login again",
				Response: registrationResponse,
			}

			successJSONResponse, jsonError := json.Marshal(successResponse)

			if jsonError != nil {
				returnErrorResponse(response, request, errorResponse)
			}
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(successResponse.Code)
			response.Write(successJSONResponse)
		}
	}
}

func GetUserDetails(response http.ResponseWriter, request *http.Request) {
	var result UserDetails
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}
	bearerToken := request.Header.Get("Authorization")
	var authorizationToken = strings.Split(bearerToken, " ")[1]

	email, _ := VerifyToken(authorizationToken)
	if email == "" {
		returnErrorResponse(response, request, errorResponse)
	} else {
		collection := Client.Database("msdb").Collection("users")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var err = collection.FindOne(ctx, bson.M{
			"email": email,
		}).Decode(&result)

		defer cancel()

		if err != nil {
			returnErrorResponse(response, request, errorResponse)
		} else {
			var successResponse = SuccessResponse{
				Code:     http.StatusOK,
				Message:  "Successfully logged in ",
				Response: result.Name,
			}

			successJSONResponse, jsonError := json.Marshal(successResponse)

			if jsonError != nil {
				returnErrorResponse(response, request, errorResponse)
			}
			response.Header().Set("Content-Type", "application/json")
			response.Write(successJSONResponse)
		}
	}
}

func returnErrorResponse(response http.ResponseWriter, request *http.Request, errorMesage ErrorResponse) {
	httpResponse := &ErrorResponse{Code: errorMesage.Code, Message: errorMesage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMesage.Code)
	response.Write(jsonResponse)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
