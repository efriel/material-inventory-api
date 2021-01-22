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
	route.HandleFunc("/userDetails", GetUserDetails).Methods("GET")
	route.HandleFunc("/api/signin", SignInUser).Methods("POST")
	route.HandleFunc("/api/signup", SignUpUser).Methods("POST")
	route.HandleFunc("/api/userDetails", GetUserDetails).Methods("GET")
	route.HandleFunc("/api/master/part", GetMasterPart).Methods("GET")
	route.HandleFunc("/api/master/part/{part_id}", GetMasterPartDetail).Methods("GET")
	route.HandleFunc("/api/master/part", CreateMasterPart).Methods("POST")
	route.HandleFunc("/api/master/part/{part_id}", RemoveMasterPart).Methods("DELETE")
	route.HandleFunc("/api/master/part/{part_id}", UpdateMasterPart).Methods("PUT")
	route.HandleFunc("/api/combo/supplier", GetMasterSupplier).Methods("GET")
	route.HandleFunc("/api/combo/site", GetMasterSite).Methods("GET")
	route.HandleFunc("/api/combo/warehouse", GetMasterWarehouse).Methods("GET")
	route.HandleFunc("/api/combo/category", GetCategory).Methods("GET")
	route.HandleFunc("/api/combo/docs", GetDocs).Methods("GET")
	log.Println("Routes are Loaded.")
}
