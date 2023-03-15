package routes

import (
	"github.com/gorilla/mux"
	"github.com/kendoow/gomastering/internal/controllers"
)

var EntryRoutes = func(router *mux.Router) {
	router.HandleFunc("/entry", controllers.CreateEntry).Methods("POST")
	router.HandleFunc("/entry", controllers.GetAllEntries).Methods("GET")
	router.HandleFunc("/entry/{entryId}", controllers.GetEntryById).Methods("GET")
	router.HandleFunc("/entry/{entryId}", controllers.DeleteEntry).Methods("DELETE")
}
