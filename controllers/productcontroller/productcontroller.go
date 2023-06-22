package productcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeypc/go-restapi-mux/models"
)


func Index(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	// if error exist
	if err := models.DB.Find(&products).Error; err != nil {
		fmt.Print(err)
	}

	response, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Show(w http.ResponseWriter, r *http.Request) {
	
}

func Create(w http.ResponseWriter, r *http.Request) {
	
}

func Update(w http.ResponseWriter, r *http.Request) {
	
}

func Delete(w http.ResponseWriter, r *http.Request) {
	
}