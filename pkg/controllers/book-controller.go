package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/muhammadfassihhaider/go-sql-book-management-system/pkg/models"
	"github.com/muhammadfassihhaider/go-sql-book-management-system/pkg/utils"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {

}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {

}
