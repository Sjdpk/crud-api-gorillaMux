package controllers

import (
	"encoding/json"
	"golang-crud/database"
	"golang-crud/models"
	"net/http"

	"github.com/gorilla/mux"
)

// @desc -> Create Books
// @access -> Public [POST]
// @route -> /api/v1/book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book models.BooksModel
	json.NewDecoder(r.Body).Decode(&book)
	database.Database.Create(&book)
	json.NewEncoder(w).Encode(book)

}

// @desc -> List all Books
// @access -> Public [GET]
// @route -> /api/v1/books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	// var book []models.BooksModel
	// database.Database.Find(&book)

	var book []models.BookSubModel
	database.Database.Model(&models.BooksModel{}).Limit(10).Find(&book)
	json.NewEncoder(w).Encode(book)
}

// @desc -> get single book by id
// @access -> Public [GET]
// @route -> /api/v1/book/{id}
func GetSingleBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	bookId := mux.Vars(r)["bid"]
	if checkIfProductExists(bookId) == false {
		json.NewEncoder(w).Encode("Book Not Found!")
		return
	}
	var book models.BooksModel
	database.Database.First(&book, bookId)
	json.NewEncoder(w).Encode(book)

}

// @desc -> Update Books by id
// @access -> Public [PUT]
// @route -> /api/v1/book/{id}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	bookId := mux.Vars(r)["bid"]
	if checkIfProductExists(bookId) == false {
		json.NewEncoder(w).Encode("Book Not Found!")
		return
	}

	var book models.BooksModel
	database.Database.First(&book, bookId)
	json.NewDecoder(r.Body).Decode(&book)
	database.Database.Save(&book)
	json.NewEncoder(w).Encode(book)

}

// @desc -> Delete Books by id
// @access -> Public [DELETE]
// @route -> /api/v1/book/{id}
func Deletebook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	bookId := mux.Vars(r)["bid"]
	if checkIfProductExists(bookId) == false {
		json.NewEncoder(w).Encode("Book Not Found!")
		return
	}
	var book models.BooksModel
	database.Database.Delete(&book, bookId)
	json.NewEncoder(w).Encode("Delete Sucess!!!")

}

//check if book is present or not
func checkIfProductExists(bookId string) bool {
	var book models.BooksModel
	database.Database.First(&book, bookId)
	if book.ID == 0 {
		return false
	}
	return true
}
