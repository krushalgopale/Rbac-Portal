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
)

func TestRegisterSimple(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")

	userData := map[string]string{
		"name":     "Test",
		"email":    "test@gmail.com",
		"password": "pass123",
		"role":     "doctor",
	}

	database.DB.Where("email = ?", userData["email"]).Delete(&models.User{})

	body, _ := json.Marshal(userData)
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	Register(c)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d", w.Code)
	}

	if !bytes.Contains(w.Body.Bytes(), []byte("Registered successfully")) {
		t.Errorf("Expected success message, got: %s", w.Body.String())
	}

	database.DB.Where("email = ?", userData["email"]).Delete(&models.User{})
}
