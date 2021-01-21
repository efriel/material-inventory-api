package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

//GetMasterPart to get list of master part data documents
func GetMasterPart(response http.ResponseWriter, request *http.Request) {
	//var result MasterPart
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_mst_part")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(context.TODO(), bson.M{})
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

//GetMasterPartDetail to get the master deatil based on the part_id
func GetMasterPartDetail(response http.ResponseWriter, request *http.Request) {
	var results MasterPart
	vars := mux.Vars(request)
	PartID := vars["part_id"]
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_mst_part")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var err = collection.FindOne(ctx, bson.M{
		"part_id": PartID,
	}).Decode(&results)

	defer cancel()
	fmt.Println(fmt.Sprintf("%#v", results))

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

//CreateMasterPart to create new master part
func CreateMasterPart(response http.ResponseWriter, request *http.Request) {
	var NewPart MasterPart
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	decoder := json.NewDecoder(request.Body)
	decoderErr := decoder.Decode(&NewPart)

	defer request.Body.Close()

	if decoderErr != nil {
		returnErrorResponse(response, request, errorResponse)
	} else {

		if NewPart.Partcode == "" {
			errorResponse.Message = "Partcode can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else if NewPart.Partname == "" {
			errorResponse.Message = "Name can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else {
			collection := Client.Database("msdb").Collection("t_mst_part")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			//count, _ := collection.CountDocuments(ctx, bson.M{})
			//count = count + 1
			//NewPartID := fmt.Sprint("R", count)

			NewPart := MasterPart{
				Partid:      NewPart.Partid,
				Mgcatid:     NewPart.Mgcatid,
				Partcode:    NewPart.Partcode,
				Partname:    NewPart.Partname,
				Partunit:    NewPart.Partunit,
				Supplierid:  NewPart.Supplierid,
				Minstock:    NewPart.Minstock,
				Costprice:   NewPart.Costprice,
				Expireddate: NewPart.Expireddate,
				Siteid:      NewPart.Siteid,
				Partnotes:   NewPart.Partnotes,
				Userid:      NewPart.Userid,
				Insertdate:  time.Now(),
				Updatedate:  time.Now(),
			}

			//inserted, jsonerror := json.Marshal(inserted)

			_, databaseErr := collection.InsertOne(ctx, bson.M{
				"part_id":      NewPart.Partid,
				"mg_cat_id":    NewPart.Mgcatid,
				"part_code":    NewPart.Partcode,
				"part_name":    NewPart.Partname,
				"part_unit":    NewPart.Partunit,
				"supplier_id":  NewPart.Supplierid,
				"min_stock":    NewPart.Minstock,
				"cost_price":   NewPart.Costprice,
				"expired_date": NewPart.Expireddate,
				"site_id":      NewPart.Siteid,
				"part_notes":   NewPart.Partnotes,
				"user_id":      NewPart.Userid,
				"insert_date":  time.Now(),
				"update_date":  time.Now(),
			})
			defer cancel()

			if databaseErr != nil {
				returnErrorResponse(response, request, errorResponse)
			}

			var successResponse = SuccessResponse{
				Code:     http.StatusOK,
				Message:  "Success",
				Response: NewPart,
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

//RemoveMasterPart to remove the document based on the part_id
func RemoveMasterPart(response http.ResponseWriter, request *http.Request) {
	//var results MasterPart
	vars := mux.Vars(request)
	PartID := vars["part_id"]
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_mst_part")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	results, err := collection.DeleteOne(ctx, bson.M{
		"part_id": PartID,
	})

	defer cancel()
	fmt.Println(fmt.Sprintf("Deleted %#v", results.DeletedCount))

	if err != nil {
		errorResponse.Message = "Document not found"
		returnErrorResponse(response, request, errorResponse)
	} else {
		var successResponse = SuccessResponse{
			Code:     http.StatusOK,
			Message:  "Success",
			Response: fmt.Sprint("Deleted Doc Number ", results.DeletedCount),
		}

		successJSONResponse, jsonError := json.Marshal(successResponse)

		if jsonError != nil {
			returnErrorResponse(response, request, errorResponse)
		}
		response.Header().Set("Content-Type", "application/json")
		response.Write(successJSONResponse)
	}
}

//UpdateMasterPart to remove the document based on the part_id
func UpdateMasterPart(response http.ResponseWriter, request *http.Request) {
	var NewPart MasterPart
	vars := mux.Vars(request)
	PartID := vars["part_id"]
	fmt.Printf("%v\n", PartID)
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	decoder := json.NewDecoder(request.Body)
	decoderErr := decoder.Decode(&NewPart)

	defer request.Body.Close()

	if decoderErr != nil {
		returnErrorResponse(response, request, errorResponse)
	} else {

		if NewPart.Partcode == "" {
			errorResponse.Message = "Partcode can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else if NewPart.Partname == "" {
			errorResponse.Message = "Name can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else {

			collection := Client.Database("msdb").Collection("t_mst_part")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			NewPart = MasterPart{
				Partid:      PartID,
				Mgcatid:     NewPart.Mgcatid,
				Partcode:    NewPart.Partcode,
				Partname:    NewPart.Partname,
				Partunit:    NewPart.Partunit,
				Supplierid:  NewPart.Supplierid,
				Minstock:    NewPart.Minstock,
				Costprice:   NewPart.Costprice,
				Expireddate: NewPart.Expireddate,
				Siteid:      NewPart.Siteid,
				Partnotes:   NewPart.Partnotes,
				Userid:      NewPart.Userid,
				Insertdate:  time.Now(),
				Updatedate:  time.Now(),
			}

			_, err := collection.UpdateOne(
				ctx,
				bson.M{"part_id": PartID},
				bson.M{
					"$set": bson.M{
						"mg_cat_id":    NewPart.Mgcatid,
						"part_code":    NewPart.Partcode,
						"part_name":    NewPart.Partname,
						"part_unit":    NewPart.Partunit,
						"supplier_id":  NewPart.Supplierid,
						"min_stock":    NewPart.Minstock,
						"cost_price":   NewPart.Costprice,
						"expired_date": NewPart.Expireddate,
						"site_id":      NewPart.Siteid,
						"part_notes":   NewPart.Partnotes,
						"user_id":      NewPart.Userid,
						"update_date":  time.Now(),
					},
				},
			)

			defer cancel()
			//fmt.Printf("%v\n", err)
			if err != nil {
				errorResponse.Message = "Document not found"
				returnErrorResponse(response, request, errorResponse)
			} else {
				var successResponse = SuccessResponse{
					Code:     http.StatusOK,
					Message:  "Success",
					Response: NewPart,
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
