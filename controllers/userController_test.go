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

func TestUserController_GetUsers(t *testing.T) {
	controller := &UserController{}

	r := gin.Default()
	r.GET("/users", controller.GetUsers)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	r.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	var users []models.User
	err := json.Unmarshal(w.Body.Bytes(), &users)
	// assert.NoError(t, err)

	// assert.True(t, len(users) >= 0)
	if err != nil {
		t.Error("Error unmarchalling JSON:", err)
	}
}

func TestUserController_CreateUser(t *testing.T) {
	controller := &UserController{}

	r := gin.Default()
	r.POST("/users", controller.CreateUser)

	w := httptest.NewRecorder()
	userData := `{"Name": "John Doe", "Mobile":"1234567890", "Latitude": 40.7128, "Longitude": -74.0060}`
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(userData))
	r.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	var createdUser models.User
	err := json.Unmarshal(w.Body.Bytes(), &createdUser)
	// assert.NoError(t, err)

	// assert.Equal(t, "John Doe", createdUser.Name)

	if err != nil {
		t.Error("Error unmarshalling JSON:", err)
	}
}
