package main
 
import (
    "log" 
    "github.com/gorilla/mux"
)
 
func AddApproutes(route *mux.Router) { 
    log.Println("Loadeding Routes...") 
    route.HandleFunc("/signin", SignInUser).Methods("POST") 
    route.HandleFunc("/signup", SignUpUser).Methods("POST") 
    route.HandleFunc("/userDetails", GetUserDetails).Methods("GET") 
    log.Println("Routes are Loaded.")
}