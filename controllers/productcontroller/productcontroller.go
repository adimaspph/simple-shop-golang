package productcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeypc/go-restapi-mux/helper"
	"github.com/jeypc/go-restapi-mux/models"
	"gorm.io/gorm"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	// if error exist
	if err := models.DB.Find(&products).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
	}

	ResponseJson(w, http.StatusOK, products)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Get ID from URL
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	// Error handling
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
	}

	// Get Product by ID
	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Product not found")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ResponseJson(w, http.StatusOK, product)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Get ID from URL
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	// Error handling
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
	}

	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	// update product
	if models.DB.Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Failed to update product")
		return
	}

	product.Id = id

	ResponseJson(w, http.StatusOK, product)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	inputId := map[string]int{"id" : 0}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&inputId); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	var product models.Product

	if models.DB.Delete(&product, inputId["id"]).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Failed to delete product")
		return
	}

	message := map[string]string{"message" : "Product has been deleted successfully"}
	ResponseJson(w, http.StatusOK, message)
}
