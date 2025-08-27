package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JamestarKurbah/go-auth-app/src/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	router := gin.Default()
	router.POST("/login", handlers.LoginHandler)

	// Create a test request
	req, _ := http.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response code
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTokenHandler(t *testing.T) {
	router := gin.Default()
	router.GET("/token", handlers.TokenHandler)

	// Create a test request
	req, _ := http.NewRequest("GET", "/token", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response code
	assert.Equal(t, http.StatusOK, w.Code)
}
