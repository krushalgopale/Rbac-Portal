package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func makeTestRouter(middleware gin.HandlerFunc, role string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		if role != "" {
			c.Set("role", role)
		}
		middleware(c)
		if !c.IsAborted() {
			c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
		}
	})

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestReceptionistMiddleware_NoRole(t *testing.T) {
	w := makeTestRouter(ReceptionistMiddleware, "")
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Role not found")
}

func TestReceptionistMiddleware_WrongRole(t *testing.T) {
	w := makeTestRouter(ReceptionistMiddleware, "doctor")
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Contains(t, w.Body.String(), "Access denied")
}

func TestReceptionistMiddleware_CorrectRole(t *testing.T) {
	w := makeTestRouter(ReceptionistMiddleware, "receptionist")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Access granted")
}

func TestDoctorMiddleware_NoRole(t *testing.T) {
	w := makeTestRouter(DoctorMiddleware, "")
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Role not found")
}

func TestDoctorMiddleware_WrongRole(t *testing.T) {
	w := makeTestRouter(DoctorMiddleware, "receptionist")
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Contains(t, w.Body.String(), "Access denied")
}

func TestDoctorMiddleware_CorrectRole(t *testing.T) {
	w := makeTestRouter(DoctorMiddleware, "doctor")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Access granted")
}
