package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server will start at port 9000")
	ConnectDatabase()
	route := mux.NewRouter()
	AddApproutes(route)
	//log.Fatal(http.ListenAndServe(":9000", route))
	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		)(route)))
}
