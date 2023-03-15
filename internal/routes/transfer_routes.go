package routes

import (
	"github.com/gorilla/mux"
	"github.com/kendoow/gomastering/internal/controllers"
)

var TransferRoutes = func(router *mux.Router) {
	router.HandleFunc("/transfer", controllers.CreateTransfer).Methods("POST")
	router.HandleFunc("/transfer", controllers.GetAllTransfers).Methods("GET")
	router.HandleFunc("/transfer/{transferId}", controllers.GetTransferById).Methods("GET")
	router.HandleFunc("/transfer/{transferId}", controllers.DeleteTransfer).Methods("DELETE")
}
