package doctor

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/krushalgopale/internal/database"
	"github.com/krushalgopale/internal/models"
)

func TestDoctorController(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	DoctorController(c)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", w.Code)
	}
	if w.Body.String() == "" || w.Body.String() == "{}" {
		t.Error("Expected dashboard message, got empty response")
	}
}

func TestGetPatientRecords(t *testing.T) {
	p := models.Patient{
		Name:    "Test",
		Email:   "test@gmail.com",
		Age:     40,
		Phone:   9876543210,
		Gender:  "male",
		Disease: "Cough",
	}
	database.DB.Create(&p)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	GetPatientRecords(c)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", w.Code)
	}
	if !contains(w.Body.String(), p.Email) {
		t.Errorf(
			"Expected patient record with email %s in response, got: %s",
			p.Email,
			w.Body.String(),
		)
	}

	database.DB.Delete(&p)
}

func TestUpdatePatientRecord(t *testing.T) {
	original := models.Patient{
		Name:    "Test2",
		Email:   "test2@gmail.com",
		Age:     30,
		Phone:   1122334455,
		Gender:  "female",
		Disease: "Fever",
	}
	database.DB.Create(&original)

	updated := models.Patient{
		Name:    "Test3",
		Email:   "test3@gmail.com",
		Age:     31,
		Phone:   5566778899,
		Gender:  "female",
		Disease: "diabetes",
	}
	body, _ := json.Marshal(updated)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: strconv.Itoa(int(original.ID))}}
	c.Request = httptest.NewRequest(
		"PUT",
		"/patients/"+strconv.Itoa(int(original.ID)),
		bytes.NewBuffer(body),
	)
	c.Request.Header.Set("Content-Type", "application/json")

	UpdatePatientRecord(c)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", w.Code)
	}
	if !contains(w.Body.String(), "updated successfully") {
		t.Errorf("Expected update confirmation message, got: %s", w.Body.String())
	}

	database.DB.Delete(&original)
}

func contains(haystack, needle string) bool {
	return len(haystack) >= len(needle) && string(haystack[:len(needle)]) == needle
}
