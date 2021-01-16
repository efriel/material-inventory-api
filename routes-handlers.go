package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func SignInUser(response http.ResponseWriter, request *http.Request) {
	var loginRequest LoginParams
	var result UserDetails
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal server error.",
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
				"email":    loginRequest.Email,
				"password": loginRequest.Password,
			}).Decode(&result)

			defer cancel()

			if err != nil {
				returnErrorResponse(response, request, errorResponse)
			} else {
				tokenString, _ := CreateJWT(loginRequest.Email)

				if tokenString == "" {
					returnErrorResponse(response, request, errorResponse)
				}

				var successResponse = SuccessResponse{
					Code:    http.StatusOK,
					Message: "Successfully registered, please login again",
					Response: SuccessfulLoginResponse{
						AuthToken: tokenString,
						Email:     loginRequest.Email,
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
			tokenString, _ := CreateJWT(registrationRequest.Email)

			if tokenString == "" {
				returnErrorResponse(response, request, errorResponse)
			}

			var registrationResponse = SuccessfulLoginResponse{
				AuthToken: tokenString,
				Email:     registrationRequest.Email,
			}

			tnow := time.Now()
			tsec := tnow.Unix()

			collection := Client.Database("msdb").Collection("users")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, databaseErr := collection.InsertOne(ctx, bson.M{
				"user_id":  tsec,
				"email":    registrationRequest.Email,
				"password": registrationRequest.Password,
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
