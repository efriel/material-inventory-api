package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server will start at http://localhost:9000/")
	ConnectDatabase()
	route := mux.NewRouter()
	AddApproutes(route)
	log.Fatal(http.ListenAndServe(":9000", route))
}
