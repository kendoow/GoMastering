package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kendoow/gomastering/internal/utils"
	"github.com/kendoow/gomastering/models"
)


var NewTransfer models.Transfer

func GetAllTransfers(w http.ResponseWriter, r *http.Request) {
	NewTransfer := models.GetAllTransfers()
	res, _ := json.Marshal(NewTransfer)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetTransferById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)	
	transferID := vars["transferId"]
	ID, err := strconv.ParseInt(transferID, 0, 0)
	if err != nil {
		fmt.Println("ошибка при парсинге id книги в GetTransferById")
	}
	transferDetails, _ := models.GetTransferById(ID)
	res, _ := json.Marshal(transferDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTransfer(w http.ResponseWriter, r *http.Request) {
	CreateTransfer := &models.Transfer{}
	utils.ParseBody(r, CreateTransfer)
	b := CreateTransfer.CreateTransfer() // возвращает тему с базы данных
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTransfer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transferID := vars["transferId"]
	ID, err := strconv.ParseInt(transferID, 0, 0)
	if err != nil {
		fmt.Println("ошибка при парсинге id аккаунта в DeleteTransfer")
	}
	deletedTransfer := models.DeleteTransfer(ID)
	res, _ := json.Marshal(deletedTransfer)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
