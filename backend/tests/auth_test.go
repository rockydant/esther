package main

import (
	controller "esther/controllers"
	"esther/routes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Error("Error loading .env file")
	}
	mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`
	validate := validator.New()
	r := routes.SetupRouter(validate)
	r.POST("/login", controller.Login)
	req, _ := http.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
