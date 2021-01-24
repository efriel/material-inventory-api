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
	//var MasterPart MasterPart
	//var AgregateMasterPart AgregateMasterPart
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_mst_part")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//cursor, err := collection.Find(context.TODO(), bson.M{})
	//var results []bson.M
	//var NumberInt = require('mongoose-int32');
	o1 := bson.M{
		"$project": bson.M{"_id": 0, "part": "$$ROOT"},
	}

	o2 := bson.M{
		"$lookup": bson.M{"localField": "part.mg_cat_id",
			"from":         "t_cat_mg",
			"foreignField": "mg_cat_id",
			"as":           "category",
		},
	}

	o3 := bson.M{
		"$unwind": bson.M{
			"path":                       "$category",
			"preserveNullAndEmptyArrays": false,
		},
	}

	o4 := bson.M{
		"$lookup": bson.M{
			"localField":   "part.supplier_id",
			"from":         "t_mst_supplier",
			"foreignField": "supplier_id",
			"as":           "supplier",
		},
	}

	o5 := bson.M{
		"$unwind": bson.M{
			"path":                       "$supplier",
			"preserveNullAndEmptyArrays": false,
		},
	}

	o6 := bson.M{
		"$lookup": bson.M{
			"localField":   "part.site_id",
			"from":         "t_mst_site",
			"foreignField": "site_id",
			"as":           "site",
		},
	}

	o7 := bson.M{
		"$unwind": bson.M{
			"path":                       "$site",
			"preserveNullAndEmptyArrays": false,
		},
	}

	o8 := bson.M{
		"$lookup": bson.M{
			"localField":   "part.user_id",
			"from":         "users",
			"foreignField": "user_id",
			"as":           "user",
		},
	}

	o9 := bson.M{
		"$unwind": bson.M{
			"path":                       "$user",
			"preserveNullAndEmptyArrays": false,
		},
	}

	o10 := bson.M{
		"$lookup": bson.M{
			"localField":   "part.part_id",
			"from":         "t_stock_view",
			"foreignField": "mg_id",
			"as":           "stock",
		},
	}

	o11 := bson.M{
		"$unwind": bson.M{
			"path":                       "$stock",
			"preserveNullAndEmptyArrays": true,
		},
	}

	o12 := bson.M{
		"$project": bson.M{
			"part.part_id":           "$part.part_id",
			"part.mg_cat_id":         "$part.mg_cat_id",
			"category.mg_cat_name":   "$category.mg_cat_name",
			"part.part_code":         "$part.part_code",
			"part.part_name":         "$part.part_name",
			"part.part_unit":         "$part.part_unit",
			"part.supplier_id":       "$part.supplier_id",
			"supplier.supplier_name": "$supplier.supplier_name",
			"part.min_stock":         "$part.min_stock",
			"part.cost_price":        "$part.cost_price",
			"part.expired_date":      "$part.expired_date",
			"part.site_id":           "$part.site_id",
			"site.site_name":         "$site.site_name",
			"part.part_notes":        "$part.part_notes",
			"part.user_id":           "$part.user_id",
			"user.name":              "$user.name",
			"stock.quantity":         "$stock.quantity",
			"part.insert_date":       "$part.insert_date",
			"part.update_date":       "$part.update_date",
		},
	}

	pipeline := bson.A{o1, o2, o3, o4, o5, o6, o7, o8, o9, o10, o11, o12}

	//pipe := mongo.Pipeline([]bson.A{o1, o2, o3, o4, o5, o6, o7, o8, o9, o10})
	showInfoCursor, err := collection.Aggregate(ctx, pipeline)
	results := []bson.M{}
	//results := AgregateMasterPart
	//err := pipe.One(&results)
	if err = showInfoCursor.All(ctx, &results); err != nil {
		panic(err)
	}

	//err = cursor.All(ctx, &results)

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
	fmt.Println(fmt.Sprintf("%#v", NewPart))
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
