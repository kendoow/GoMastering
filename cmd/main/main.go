package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kendoow/gomastering/internal/routes"
)


func main () {
	r := mux.NewRouter()

	routes.AccountRoutes(r)
	routes.EntryRoutes(r)
	routes.TransferRoutes(r)
	
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}