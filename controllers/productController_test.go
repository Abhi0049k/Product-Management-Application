package controllers

import (
	"Product-Management-Application/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestProductController_GetProducts(t *testing.T) {
	controller := &ProductController{}

	r := gin.Default()
	r.GET("/products", controller.GetProducts)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	var products []models.Product
	err := json.Unmarshal(w.Body.Bytes(), &products)

	if err != nil {
		t.Error("Error unmarshalling JSON:", err)
	}
}

func TestProductController_CreateProduct(t *testing.T) {
	controllers := &ProductController{}

	r := gin.Default()
	r.POST("/products", controllers.CreateProduct)
	w := httptest.NewRecorder()
	productData := `{"user_id": 1, "product_name": "Laptop", "product_description": "Powerful laptop", "product_price": 1000.0, "product_images" : ["https://cdn.pixabay.com/photo/2016/02/17/19/08/lotus-1205631_1280.jpg", "https://cdn.pixabay.com/photo/2016/02/17/19/08/lotus-1205631_1280.jpg"]}`
	req, _ := http.NewRequest("POST", "/products", strings.NewReader(productData))
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	var createdProduct models.Product
	err := json.Unmarshal(w.Body.Bytes(), &createdProduct)

	if err != nil {
		t.Error("Error unmarshalling JSON:", err)
	}
}
