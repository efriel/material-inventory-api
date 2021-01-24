package main

import (
	"log"

	"github.com/gorilla/mux"
)

// AddApproutes is to handle all routes coming from the http request
func AddApproutes(route *mux.Router) {
	log.Println("Loadeding Routes...")
	route.HandleFunc("/signin", SignInUser).Methods("POST")
	route.HandleFunc("/signup", SignUpUser).Methods("POST")
	route.HandleFunc("/userdetails", GetUserDetails).Methods("GET")
	route.HandleFunc("/api/signin", SignInUser).Methods("POST")
	route.HandleFunc("/api/signup", SignUpUser).Methods("POST")
	route.HandleFunc("/api/userdetails", GetUserDetails).Methods("GET")
	route.HandleFunc("/api/completeuserdetails", GetCompleteUserDetails).Methods("GET")
	route.HandleFunc("/api/master/part", GetMasterPart).Methods("GET")
	route.HandleFunc("/api/master/part/{part_id}", GetMasterPartDetail).Methods("GET")
	route.HandleFunc("/api/master/part", CreateMasterPart).Methods("POST")
	route.HandleFunc("/api/master/part/{part_id}", RemoveMasterPart).Methods("DELETE")
	route.HandleFunc("/api/master/part/{part_id}", UpdateMasterPart).Methods("PUT")
	route.HandleFunc("/api/master/goods", GetMasterGoodsView).Methods("GET")
	route.HandleFunc("/api/master/goods/{fg_id}", GetMasterGoodsDetail).Methods("GET")
	route.HandleFunc("/api/master/goods", CreateMasterGoods).Methods("POST")
	route.HandleFunc("/api/master/goods/{fg_id}", RemoveMasterGoods).Methods("DELETE")
	route.HandleFunc("/api/master/goods/{fg_id}", UpdateMasterGoods).Methods("PUT")
	route.HandleFunc("/api/purchase/status", GetPurchaseView).Methods("GET")
	route.HandleFunc("/api/purchase/status/{status_flag}", GetPurchaseViewFlag).Methods("GET")
	route.HandleFunc("/api/purchase/status", CreatePurchase).Methods("POST")
	route.HandleFunc("/api/purchase/status/{purchase_id}", UpdatePurchase).Methods("PUT")
	route.HandleFunc("/api/combo/supplier", GetMasterSupplier).Methods("GET")
	route.HandleFunc("/api/combo/site", GetMasterSite).Methods("GET")
	route.HandleFunc("/api/combo/warehouse", GetMasterWarehouse).Methods("GET")
	route.HandleFunc("/api/combo/category", GetCategory).Methods("GET")
	route.HandleFunc("/api/combo/docs", GetDocs).Methods("GET")
	route.HandleFunc("/api/combo/parts", GetParts).Methods("GET")
	route.HandleFunc("/api/combo/users", GetUsers).Methods("GET")
	log.Println("Routes are Loaded.")
}
