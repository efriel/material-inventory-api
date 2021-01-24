package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

//GetPurchaseView to get list of purchase data documents
func GetPurchaseView(response http.ResponseWriter, request *http.Request) {
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_purchase_view")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	showInfoCursor, err := collection.Find(context.TODO(), bson.M{})
	results := []bson.M{}
	if err = showInfoCursor.All(ctx, &results); err != nil {
		panic(err)
	}

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

//GetPurchaseViewFlag to get list of purchase data documents by status flag
func GetPurchaseViewFlag(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	StatusFlag := vars["status_flag"]
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	collection := Client.Database("msdb").Collection("t_purchase_view_" + StatusFlag)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	showInfoCursor, err := collection.Find(context.TODO(), bson.M{})
	results := []bson.M{}
	if err = showInfoCursor.All(ctx, &results); err != nil {
		panic(err)
	}

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

//CreatePurchase to create new master fg
func CreatePurchase(response http.ResponseWriter, request *http.Request) {
	var NewPurchase TPurchase
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	decoder := json.NewDecoder(request.Body)
	decoderErr := decoder.Decode(&NewPurchase)

	defer request.Body.Close()

	if decoderErr != nil {
		returnErrorResponse(response, request, errorResponse)
	} else {

		if NewPurchase.Partid == "" {
			errorResponse.Message = "PartID can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else if NewPurchase.Statusflag == "" {
			errorResponse.Message = "Status can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else {
			collection := Client.Database("msdb").Collection("t_purchase")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			//count, _ := collection.CountDocuments(ctx, bson.M{})
			//count = count + 1
			//NewFgID := fmt.Sprint("R", count)
			tnow := time.Now()
			tbid := tnow.AddDate(0, 0, 7*1)
			tclose := tnow.AddDate(0, 0, 7*2)
			tsec := tnow.Unix()
			ntsec := strconv.FormatInt(tsec, 10)
			NPurchaseid := ntsec
			NewPurchase := TPurchase{
				Purchaseid:    NPurchaseid,
				Supplierid:    NewPurchase.Supplierid,
				Whid:          NewPurchase.Whid,
				Partid:        NewPurchase.Partid,
				Qty:           NewPurchase.Qty,
				Purchasedate:  time.Now(),
				Estimatedcost: NewPurchase.Estimatedcost,
				Invoice:       NewPurchase.Invoice,
				Receipt:       NewPurchase.Receipt,
				Buyerid:       NewPurchase.Buyerid,
				Originatorid:  NewPurchase.Originatorid,
				Userid:        NewPurchase.Userid,
				Notes:         NewPurchase.Notes,
				Statusflag:    NewPurchase.Statusflag,
				Bidoutdate:    tbid,
				Closeddate:    tclose,
				Insertdate:    time.Now(),
				Updatedate:    time.Now(),
			}
			fmt.Println(fmt.Sprintf("%#v", NewPurchase))
			//inserted, jsonerror := json.Marshal(inserted)

			_, databaseErr := collection.InsertOne(ctx, bson.M{
				"purchase_id":    NewPurchase.Purchaseid,
				"supplier_id":    NewPurchase.Supplierid,
				"wh_id":          NewPurchase.Whid,
				"part_id":        NewPurchase.Partid,
				"qty":            NewPurchase.Qty,
				"purchase_date":  NewPurchase.Purchasedate,
				"estimated_cost": NewPurchase.Estimatedcost,
				"invoice":        NewPurchase.Invoice,
				"receipt":        NewPurchase.Receipt,
				"buyer_id":       NewPurchase.Buyerid,
				"originator_id":  NewPurchase.Originatorid,
				"user_id":        NewPurchase.Userid,
				"notes":          NewPurchase.Notes,
				"status_flag":    NewPurchase.Statusflag,
				"bidout_date":    NewPurchase.Bidoutdate,
				"closed_date":    NewPurchase.Closeddate,
				"insert_date":    NewPurchase.Insertdate,
				"update_date":    NewPurchase.Updatedate,
			})
			defer cancel()

			if databaseErr != nil {
				returnErrorResponse(response, request, errorResponse)
			}

			var successResponse = SuccessResponse{
				Code:     http.StatusOK,
				Message:  "Success",
				Response: NewPurchase,
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

//UpdatePurchase to remove the document based on the part_id
func UpdatePurchase(response http.ResponseWriter, request *http.Request) {
	var NewPurchase TPurchase
	vars := mux.Vars(request)
	PurchaseID := vars["purchase_id"]
	fmt.Printf("%v\n", PurchaseID)
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal Server Error.",
	}

	decoder := json.NewDecoder(request.Body)
	decoderErr := decoder.Decode(&NewPurchase)

	defer request.Body.Close()

	if decoderErr != nil {
		returnErrorResponse(response, request, errorResponse)
	} else {

		if NewPurchase.Statusflag == "" {
			errorResponse.Message = "Status can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else {

			collection := Client.Database("msdb").Collection("t_purchase")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			NewPurchase = TPurchase{
				Purchaseid:    PurchaseID,
				Supplierid:    NewPurchase.Supplierid,
				Whid:          NewPurchase.Whid,
				Partid:        NewPurchase.Partid,
				Qty:           NewPurchase.Qty,
				Purchasedate:  NewPurchase.Purchasedate,
				Estimatedcost: NewPurchase.Estimatedcost,
				Invoice:       NewPurchase.Invoice,
				Receipt:       NewPurchase.Receipt,
				Buyerid:       NewPurchase.Buyerid,
				Originatorid:  NewPurchase.Originatorid,
				Userid:        NewPurchase.Userid,
				Notes:         NewPurchase.Notes,
				Statusflag:    NewPurchase.Statusflag,
				Bidoutdate:    NewPurchase.Bidoutdate,
				Closeddate:    NewPurchase.Closeddate,
				Insertdate:    NewPurchase.Insertdate,
				Updatedate:    time.Now(),
			}

			_, err := collection.UpdateOne(
				ctx,
				bson.M{"purchase_id": PurchaseID},
				bson.M{
					"$set": bson.M{
						"supplier_id":    NewPurchase.Supplierid,
						"wh_id":          NewPurchase.Whid,
						"part_id":        NewPurchase.Partid,
						"qty":            NewPurchase.Qty,
						"purchase_date":  NewPurchase.Purchasedate,
						"estimated_cost": NewPurchase.Estimatedcost,
						"invoice":        NewPurchase.Invoice,
						"receipt":        NewPurchase.Receipt,
						"buyer_id":       NewPurchase.Buyerid,
						"originator_id":  NewPurchase.Originatorid,
						"user_id":        NewPurchase.Userid,
						"notes":          NewPurchase.Notes,
						"status_flag":    NewPurchase.Statusflag,
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
					Response: NewPurchase,
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
