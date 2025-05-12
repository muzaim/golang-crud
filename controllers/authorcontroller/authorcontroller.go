package authorcontroller

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
	var authors []models.Author

	if err := config.DB.Find(&authors).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var authorResponses []models.AuthorResponse
	for _, a := range authors {
		authorResponses = append(authorResponses, models.AuthorResponse{
			ID:     a.ID,
			Name:   a.Name,
			Gender: a.Gender,
			Email:  a.Email,
			Age:    a.Age,
		})
	}

	helper.Response(w, 200, "GET ALL AUTHOR", authorResponses)
}

func Index2(w http.ResponseWriter, r *http.Request) {
	var author *models.Author

	if err := config.DB.First(&author).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	authorResponse := models.AuthorResponse{
		ID:     author.ID,
		Name:   author.Name,
		Gender: author.Gender,
		Email:  author.Email,
		Age:    author.Age,
	}

	helper.Response(w, 200, "GET AUTHOR RESPONSE SUCCESS", authorResponse)
}

func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)
	var author *models.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "GET AUTHOR SUCCESS", author)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Create(&author).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success", nil)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)
	var author *models.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Where("id=?", id).Updates(&author).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success", author)

}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var author models.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Delete(&author, id).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Delete berhasil", nil)

}
