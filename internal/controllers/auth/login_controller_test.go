package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/krushalgopale/internal/database"
	"github.com/krushalgopale/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginSimple(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")

	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := models.User{
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: string(password),
		Role:     "doctor",
	}
	database.DB.Where("email = ?", user.Email).Delete(&models.User{}) // clean previous
	database.DB.Create(&user)

	body := map[string]string{
		"email":    "test@gmail.com",
		"password": "pass123",
	}
	jsonBody, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	Login(c)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}
	if !bytes.Contains(w.Body.Bytes(), []byte("Login successful")) {
		t.Fatalf("Expected login success message, got: %s", w.Body.String())
	}
}
