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

//RenderHome to render profile page if needed
func RenderHome(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "views/profile.html")
}

//RenderLogin to render login page if needed
func RenderLogin(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "views/login.html")
}

//RenderRegister to render register page if needed
func RenderRegister(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "views/register.html")
}

//SignInUser to accept request from user login
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

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
					tokenString, _ := CreateJWT(result.Name, loginRequest.Email, result.Userid)

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
							Userid:    result.Userid,
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

//SignUpUser to accept request from user signup
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
			ntsecint, _ := strconv.Atoi(ntsec)

			tokenString, _ := CreateJWT(registrationRequest.Name, registrationRequest.Email, ntsecint)

			if tokenString == "" {
				returnErrorResponse(response, request, errorResponse)
			}

			var registrationResponse = SuccessfulLoginResponse{
				AuthToken: tokenString,
				Email:     registrationRequest.Email,
				Name:      registrationRequest.Name,
			}

			collection := Client.Database("msdb").Collection("users")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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

//GetUserDetails to accept request for user detail
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
		//result := []bson.M{}
		fmt.Printf("%v\n", result)
		defer cancel()

		if err != nil {
			returnErrorResponse(response, request, errorResponse)
		} else {
			var successResponse = SuccessResponse{
				Code:     http.StatusOK,
				Message:  "Successfully logged in ",
				Response: result,
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

//GetCompleteUserDetails to accept request for user detail
func GetCompleteUserDetails(response http.ResponseWriter, request *http.Request) {
	var results CompleteUserDetails
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
		}).Decode(&results)

		defer cancel()
		fmt.Printf("%v\n", results)

		if err != nil {
			errorResponse.Message = "User not found"
			returnErrorResponse(response, request, errorResponse)
		} else {
			var successResponse = SuccessResponse{
				Code:     http.StatusOK,
				Message:  "Success",
				Response: results,
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

//CheckPasswordHash to compare the password given
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
