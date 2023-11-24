package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func BenchmarkUserController_GetUsers(b *testing.B) {
	controller := &UserController{}

	r := gin.Default()
	r.GET("/users", controller.GetUsers)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/users", nil)
		r.ServeHTTP(w, req)

		assert.Equal(b, http.StatusOK, w.Code)
	}
}

func BenchmarkUserController_CreateUser(b *testing.B) {
	controller := &UserController{}

	r := gin.Default()
	r.POST("/users", controller.CreateUser)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		userData := `{"name": "John Doe", "mobile": "1234567890", "latitude": 54.6589, "longitude": -24.6532}`
		req, _ := http.NewRequest("POST", "/users", strings.NewReader(userData))
		r.ServeHTTP(w, req)

		assert.Equal(b, http.StatusOK, w.Code)
	}
}
