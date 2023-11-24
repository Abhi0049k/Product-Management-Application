package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func BenchmarkProductController_GetProducts(b *testing.B) {
	controller := &ProductController{}

	r := gin.Default()
	r.GET("/products", controller.GetProducts)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/products", nil)
		r.ServeHTTP(w, req)

		assert.Equal(b, http.StatusOK, w.Code)
	}
}

func BenchmarkProductController_CreateProduct(b *testing.B) {
	controller := &ProductController{}

	r := gin.Default()
	r.POST("/products", controller.CreateProduct)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		productData := `{"user_id": 1, "product_name": "Laptop", "product_description": "Powerful laptop", "product_price": 1000.0, "product_images" : ["https://cdn.pixabay.com/photo/2016/02/17/19/08/lotus-1205631_1280.jpg", "https://cdn.pixabay.com/photo/2016/02/17/19/08/lotus-1205631_1280.jpg"]}`
		req, _ := http.NewRequest("POST", "/products", strings.NewReader(productData))

		r.ServeHTTP(w, req)

		assert.Equal(b, http.StatusOK, w.Code)
	}
}
