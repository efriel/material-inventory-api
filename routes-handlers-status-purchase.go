package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
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
				if NewPurchase.Statusflag == "D" {
					newQty, _ := strconv.Atoi(NewPurchase.Qty)
					partid := NewPurchase.Partid
					qty := NewPurchase.Qty
					InsertStock(partid, newQty)
					getUserAndEmail(NewPurchase.Buyerid, partid, qty)
					getUserAndEmail(NewPurchase.Originatorid, partid, qty)
					getUserAndEmail(NewPurchase.Userid, partid, qty)
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
				response.Write(successJSONResponse)
			}
		}
	}
}

//InsertStock function to update new stock value
func InsertStock(Mgid string, Qty int) {
	var NewStock TStock

	tnow := time.Now()
	tsec := tnow.Unix()
	ntsec := strconv.FormatInt(tsec, 10)
	Stockid := ntsec

	collection := Client.Database("msdb").Collection("t_stock")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	NewStock = TStock{
		Stockid:    Stockid,
		Mgid:       Mgid,
		Quantity:   Qty,
		Insertdate: time.Now(),
		Updatedate: time.Now(),
	}

	_, databaseErr := collection.InsertOne(ctx, bson.M{
		"stock_id":    NewStock.Stockid,
		"mg_id":       NewStock.Mgid,
		"quantity":    NewStock.Quantity,
		"insert_date": NewStock.Insertdate,
		"update_date": NewStock.Updatedate,
	})
	defer cancel()

	if databaseErr != nil {

	}
}

func getUserAndEmail(userid int, partid string, qty string) error {
	var resUserDetails UserDetails
	collection := Client.Database("msdb").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var err = collection.FindOne(ctx, bson.M{
		"user_id": userid,
	}).Decode(&resUserDetails)
	defer cancel()
	Email := resUserDetails.Email
	Part := "Part ID: " + partid
	Qty := qty
	to := []string{Email}
	cc := []string{"efriel.apps@gmail.com"}
	subject := "Penambahan Stock"
	message := "Dengan Hormat, Ini adalah penambahan stock " + Part + " Sebanyak " + Qty + ", Terimakasih"
	err2 := sendMail(to, cc, subject, message)
	if err2 != nil {
		log.Fatal(err.Error())
	}
	log.Println("Mail sent!")
	return nil
}

func sendMail(to []string, cc []string, subject, message string) error {
	CfgSMTPHost := "smtp.gmail.com"
	CfgSMTPPort := 587
	CfgSenderName := "MIS System"
	CfgAuthEmail := "efriel.apps@gmail.com"
	CfgAuthPwd := "CapitalGainApps!"

	body := "From: " + CfgSenderName + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CfgAuthEmail, CfgAuthPwd, CfgSMTPHost)
	smtpAddr := fmt.Sprintf("%s:%d", CfgSMTPHost, CfgSMTPPort)

	err := smtp.SendMail(smtpAddr, auth, CfgAuthEmail, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
