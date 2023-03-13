package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kendoow/gomastering/models"
	"github.com/kendoow/gomastering/internal/utils"
)

var NewAccount models.Account

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	NewAccount := models.GetAllAccounts()
	res, _ := json.Marshal(NewAccount)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	CreateAccount := &models.Account{}
	utils.ParseBody(r, CreateAccount)
	b := CreateAccount.CreateAccount() // возвращает тему с базы данных
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["accountId"]
	ID, err := strconv.ParseInt(accountID, 0, 0)
	if err != nil {
		fmt.Println("ошибка при парсинге id аккаунта в DeleteAccount")
	}
	deletedAccount := models.DeleteAccount(ID)
	res, _ := json.Marshal(deletedAccount)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var updateAccount = &models.Account{}
	utils.ParseBody(r, updateAccount)
	vars := mux.Vars(r)
	accountID := vars["accountId"]
	ID, err := strconv.ParseInt(accountID, 0, 0)
	if err != nil {
		fmt.Println("ошибка при парсинге id аккаунта в UpdateAccount")
	}
	accountDetails, _ := models.GetAccountById(ID)
	
	
	res, _ := json.Marshal(accountDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}