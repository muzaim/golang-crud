package bookcontroller

import (
	"encoding/json"
	"golang-crud/config"
	"golang-crud/helper"
	"golang-crud/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var book []models.Book

	if err := config.DB.Find(&book).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var bookResponse []models.BookResponse
	for _, a := range book {
		bookResponse = append(bookResponse, models.BookResponse{
			ID:     a.ID,
			Title:  a.Title,
			Price:  a.Price,
			Author: a.Author,
		})
	}

	helper.Response(w, 200, "Success", bookResponse)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Create(&book).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Book created", nil)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)
	var book *models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	bookResponse := models.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Price:       book.Price,
		Author:      book.Author,
	}

	helper.Response(w, 200, "GET BOOK BY ID", bookResponse)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book *models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Where("id=?", id).Updates(&book).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 201, "Berhasil", book)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Delete(&book, id).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Deleted", nil)
}
