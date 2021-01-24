package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//GetMasterSupplier to get list of master supplier
func GetMasterSupplier(response http.ResponseWriter, request *http.Request) {
	//var results MstSupplier
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_mst_supplier")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	var results []bson.M
	err = cursor.All(ctx, &results)

	defer cancel()

	if err != nil {
		errorResponse.Message = "Document not found"
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

//GetMasterSite to get list of master site
func GetMasterSite(response http.ResponseWriter, request *http.Request) {
	//var results MstSite
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_mst_site")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	var results []bson.M
	err = cursor.All(ctx, &results)

	defer cancel()

	if err != nil {
		errorResponse.Message = "Document not found"
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

//GetMasterWarehouse to get list of master warehouse
func GetMasterWarehouse(response http.ResponseWriter, request *http.Request) {
	//var results MstWarehouse
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_mst_warehouse")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	var results []bson.M
	err = cursor.All(ctx, &results)

	defer cancel()

	if err != nil {
		errorResponse.Message = "Document not found"
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

//GetCategory to get list of category
func GetCategory(response http.ResponseWriter, request *http.Request) {
	//var results TCategory
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_cat_mg")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	var results []bson.M
	err = cursor.All(ctx, &results)

	defer cancel()

	if err != nil {
		errorResponse.Message = "Document not found"
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

//GetDocs to get list of documents
func GetDocs(response http.ResponseWriter, request *http.Request) {
	//var results TDoc
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_doc")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	var results []bson.M
	err = cursor.All(ctx, &results)

	defer cancel()

	if err != nil {
		errorResponse.Message = "Document not found"
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

//GetParts to get list of part
func GetParts(response http.ResponseWriter, request *http.Request) {
	//var results TDoc
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_mst_part")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	var results []bson.M
	err = cursor.All(ctx, &results)

	defer cancel()

	if err != nil {
		errorResponse.Message = "Document not found"
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

//GetUsers to get list of User
func GetUsers(response http.ResponseWriter, request *http.Request) {
	//var results TDoc
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	var results []bson.M
	err = cursor.All(ctx, &results)

	defer cancel()

	if err != nil {
		errorResponse.Message = "Document not found"
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
