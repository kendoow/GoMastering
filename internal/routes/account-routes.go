package routes

import (
	"github.com/gorilla/mux"
	"github.com/kendoow/gomastering/internal/controllers"
)

var AccountRoutes = func (router *mux.Router)  {
	router.HandleFunc("/account", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/account", controllers.GetAccounts).Methods("GET")
	router.HandleFunc("/account/{accountId}",controllers.UpdateAccount).Methods("PUT")
	router.HandleFunc("/account/{accountId}",controllers.DeleteAccount).Methods("DELETE")
}	