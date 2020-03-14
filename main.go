package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohamedabdelghani/myfirstgoapp/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterPersonRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
