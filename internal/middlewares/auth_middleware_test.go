package middlewares

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func generateToken(secret string, claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}

func TestAuthMiddleware_MissingToken(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest("GET", "/", nil)
	c.Request = req

	AuthMiddleware(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Token missing")
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: "jwt", Value: "invalidtoken"})
	c.Request = req

	os.Setenv("JWT_SECRET", "secret")
	AuthMiddleware(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid token")
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
	// Setup valid token
	claims := jwt.MapClaims{
		"user_id": "1",
		"role":    "receptionist",
		"exp":     time.Now().Add(time.Hour).Unix(),
	}
	secret := "testsecret"
	tokenString := generateToken(secret, claims)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: "jwt", Value: tokenString})
	c.Request = req

	os.Setenv("JWT_SECRET", secret)
	AuthMiddleware(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "123", c.GetString("userID"))
	assert.Equal(t, "admin", c.GetString("role"))
}
