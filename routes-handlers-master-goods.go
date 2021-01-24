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

//GetMasterGoods to get list of master fg data documents
func GetMasterGoods(response http.ResponseWriter, request *http.Request) {
	//var MasterGoods MasterGoods
	//var AgregateMasterGoods AgregateMasterGoods
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_mst_finish_goods")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//cursor, err := collection.Find(context.TODO(), bson.M{})
	//var results []bson.M
	//var NumberInt = require('mongoose-int32');
	o1 := bson.M{
		"$project": bson.M{"_id": 0, "goods": "$$ROOT"},
	}

	o2 := bson.M{
		"$lookup": bson.M{"localField": "goods.mg_cat_id",
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
			"localField":   "goods.site_id",
			"from":         "t_mst_site",
			"foreignField": "site_id",
			"as":           "site",
		},
	}

	o5 := bson.M{
		"$unwind": bson.M{
			"path":                       "$site",
			"preserveNullAndEmptyArrays": false,
		},
	}

	o6 := bson.M{
		"$lookup": bson.M{
			"localField":   "goods.user_id",
			"from":         "users",
			"foreignField": "user_id",
			"as":           "user",
		},
	}

	o7 := bson.M{
		"$unwind": bson.M{
			"path":                       "$user",
			"preserveNullAndEmptyArrays": false,
		},
	}

	o8 := bson.M{
		"$lookup": bson.M{
			"localField":   "goods.fg_id",
			"from":         "t_stock_view",
			"foreignField": "mg_id",
			"as":           "stock",
		},
	}

	o9 := bson.M{
		"$unwind": bson.M{
			"path":                       "$stock",
			"preserveNullAndEmptyArrays": true,
		},
	}

	o10 := bson.M{
		"$project": bson.M{
			"goods.fg_id":            "$goods.fg_id",
			"goods.mg_cat_id":        "$goods.mg_cat_id",
			"category.mg_cat_name":   "$category.mg_cat_name",
			"goods.fg_code":          "$goods.fg_code",
			"goods.fg_name":          "$goods.fg_name",
			"goods.fg_unit":          "$goods.fg_unit",
			"goods.min_stock":        "$goods.min_stock",
			"goods.production_cost":  "$goods.production_cost",
			"goods.percent_markup":   "$goods.percent_markup",
			"goods.percent_discount": "$goods.percent_discount",
			"goods.expired_date":     "$goods.expired_date",
			"goods.site_id":          "$goods.site_id",
			"site.site_name":         "$site.site_name",
			"goods.fg_notes":         "$goods.fg_notes",
			"goods.user_id":          "$goods.user_id",
			"user.name":              "$user.name",
			"stock.quantity":         "$stock.quantity",
			"goods.insert_date":      "$goods.insert_date",
			"goods.update_date":      "$goods.update_date",
		},
	}

	pipeline := bson.A{o1, o2, o3, o4, o5, o6, o7, o8, o9, o10}

	//pipe := mongo.Pipeline([]bson.A{o1, o2, o3, o4, o5, o6, o7, o8, o9, o10})
	showInfoCursor, err := collection.Aggregate(ctx, pipeline)
	results := []bson.M{}
	//results := AgregateMasterGoods
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

//GetMasterGoodsDetail to get the master deatil based on the fg_id
func GetMasterGoodsDetail(response http.ResponseWriter, request *http.Request) {
	var results MasterGoods
	vars := mux.Vars(request)
	FgID := vars["fg_id"]
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_mst_finish_goods")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var err = collection.FindOne(ctx, bson.M{
		"fg_id": FgID,
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

//GetMasterGoodsView to get list of master fg data documents
func GetMasterGoodsView(response http.ResponseWriter, request *http.Request) {
	//var MasterGoods MasterGoods
	//var AgregateMasterGoods AgregateMasterGoods
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_mst_finish_goods_view")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	showInfoCursor, err := collection.Find(context.TODO(), bson.M{})
	//var results []bson.M
	//var NumberInt = require('mongoose-int32');
	//showInfoCursor, err := collection.Aggregate(ctx, pipeline)
	results := []bson.M{}
	//results := AgregateMasterGoods
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

//CreateMasterGoods to create new master fg
func CreateMasterGoods(response http.ResponseWriter, request *http.Request) {
	var NewGoods MasterGoods
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	decoder := json.NewDecoder(request.Body)
	decoderErr := decoder.Decode(&NewGoods)
	fmt.Println(fmt.Sprintf("%#v", NewGoods))
	defer request.Body.Close()

	if decoderErr != nil {
		returnErrorResponse(response, request, errorResponse)
	} else {

		if NewGoods.Fgcode == "" {
			errorResponse.Message = "Fgcode can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else if NewGoods.Fgname == "" {
			errorResponse.Message = "Name can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else {
			collection := Client.Database("msdb").Collection("t_mst_finish_goods")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			//count, _ := collection.CountDocuments(ctx, bson.M{})
			//count = count + 1
			//NewFgID := fmt.Sprint("R", count)

			NewGoods := MasterGoods{
				Fgid:            NewGoods.Fgid,
				Mgcatid:         NewGoods.Mgcatid,
				Fgcode:          NewGoods.Fgcode,
				Fgname:          NewGoods.Fgname,
				Fgunit:          NewGoods.Fgunit,
				Minstock:        NewGoods.Minstock,
				Costprice:       NewGoods.Costprice,
				Percentmarkup:   NewGoods.Percentmarkup,
				Percentdiscount: NewGoods.Percentdiscount,
				Netprice:        NewGoods.Netprice,
				Expireddate:     NewGoods.Expireddate,
				Siteid:          NewGoods.Siteid,
				Fgnotes:         NewGoods.Fgnotes,
				Userid:          NewGoods.Userid,
				Insertdate:      time.Now(),
				Updatedate:      time.Now(),
			}

			//inserted, jsonerror := json.Marshal(inserted)

			_, databaseErr := collection.InsertOne(ctx, bson.M{
				"fg_id":            NewGoods.Fgid,
				"mg_cat_id":        NewGoods.Mgcatid,
				"fg_code":          NewGoods.Fgcode,
				"fg_name":          NewGoods.Fgname,
				"fg_unit":          NewGoods.Fgunit,
				"min_stock":        NewGoods.Minstock,
				"production_cost":  NewGoods.Costprice,
				"percent_markup":   NewGoods.Percentmarkup,
				"percent_discount": NewGoods.Percentdiscount,
				"net_price":        NewGoods.Netprice,
				"expired_date":     NewGoods.Expireddate,
				"site_id":          NewGoods.Siteid,
				"fg_notes":         NewGoods.Fgnotes,
				"user_id":          NewGoods.Userid,
				"insert_date":      time.Now(),
				"update_date":      time.Now(),
			})
			defer cancel()

			if databaseErr != nil {
				returnErrorResponse(response, request, errorResponse)
			}

			var successResponse = SuccessResponse{
				Code:     http.StatusOK,
				Message:  "Success",
				Response: NewGoods,
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

//RemoveMasterGoods to remove the document based on the part_id
func RemoveMasterGoods(response http.ResponseWriter, request *http.Request) {
	//var results MasterGoods
	vars := mux.Vars(request)
	FgID := vars["fg_id"]
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_mst_finish_goods")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	results, err := collection.DeleteOne(ctx, bson.M{
		"fg_id": FgID,
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

//UpdateMasterGoods to remove the document based on the part_id
func UpdateMasterGoods(response http.ResponseWriter, request *http.Request) {
	var NewGoods MasterGoods
	vars := mux.Vars(request)
	FgID := vars["fg_id"]
	fmt.Printf("%v\n", FgID)
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	decoder := json.NewDecoder(request.Body)
	decoderErr := decoder.Decode(&NewGoods)

	defer request.Body.Close()

	if decoderErr != nil {
		returnErrorResponse(response, request, errorResponse)
	} else {

		if NewGoods.Fgcode == "" {
			errorResponse.Message = "Fgcode can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else if NewGoods.Fgname == "" {
			errorResponse.Message = "Name can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else {

			collection := Client.Database("msdb").Collection("t_mst_finish_goods")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			NewGoods = MasterGoods{
				Fgid:            FgID,
				Mgcatid:         NewGoods.Mgcatid,
				Fgcode:          NewGoods.Fgcode,
				Fgname:          NewGoods.Fgname,
				Fgunit:          NewGoods.Fgunit,
				Minstock:        NewGoods.Minstock,
				Costprice:       NewGoods.Costprice,
				Percentmarkup:   NewGoods.Percentmarkup,
				Percentdiscount: NewGoods.Percentdiscount,
				Netprice:        NewGoods.Netprice,
				Expireddate:     NewGoods.Expireddate,
				Siteid:          NewGoods.Siteid,
				Fgnotes:         NewGoods.Fgnotes,
				Userid:          NewGoods.Userid,
				Insertdate:      time.Now(),
				Updatedate:      time.Now(),
			}

			_, err := collection.UpdateOne(
				ctx,
				bson.M{"fg_id": FgID},
				bson.M{
					"$set": bson.M{
						"mg_cat_id":        NewGoods.Mgcatid,
						"fg_code":          NewGoods.Fgcode,
						"fg_name":          NewGoods.Fgname,
						"fg_unit":          NewGoods.Fgunit,
						"min_stock":        NewGoods.Minstock,
						"production_cost":  NewGoods.Costprice,
						"percent_markup":   NewGoods.Percentmarkup,
						"percent_discount": NewGoods.Percentdiscount,
						"net_price":        NewGoods.Netprice,
						"expired_date":     NewGoods.Expireddate,
						"site_id":          NewGoods.Siteid,
						"fg_notes":         NewGoods.Fgnotes,
						"user_id":          NewGoods.Userid,
						"update_date":      time.Now(),
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
					Response: NewGoods,
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
