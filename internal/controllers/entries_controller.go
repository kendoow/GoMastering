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


var NewEntry models.Entry

func GetAllEntries(w http.ResponseWriter, r *http.Request) {
	NewEntry := models.GetAllEntries()
	res, _ := json.Marshal(NewEntry)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetEntryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entryID := vars["entryId"]
	ID, err := strconv.ParseInt(entryID, 0, 0)
	if err != nil {
		fmt.Println("ошибка при парсинге id книги в GetEntryById")
	}
	entryDetails, _ := models.GetEntryById(ID)
	res, _ := json.Marshal(entryDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateEntry(w http.ResponseWriter, r *http.Request) {
	CreateEntry := &models.Entry{}
	utils.ParseBody(r, CreateEntry)
	b := CreateEntry.CreateEntry() // возвращает тему с базы данных
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entryID := vars["entryId"]
	ID, err := strconv.ParseInt(entryID, 0, 0)
	if err != nil {
		fmt.Println("ошибка при парсинге id аккаунта в DeleteEntry")
	}
	deletedEntry := models.DeleteEntry(ID)
	res, _ := json.Marshal(deletedEntry)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}