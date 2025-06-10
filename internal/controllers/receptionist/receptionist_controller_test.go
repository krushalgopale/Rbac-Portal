package receptionist

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

func contains(haystack, needle string) bool {
	return len(haystack) >= len(needle) && string(haystack[:len(needle)]) == needle
}

func TestReceptionistController(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	ReceptionistController(c)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", w.Code)
	}
	if !contains(w.Body.String(), "Receptionist Dashboard") {
		t.Errorf("Expected dashboard message, got: %s", w.Body.String())
	}
}

func TestCreatePatientRecord(t *testing.T) {
	patient := models.Patient{
		Name:    "Test",
		Email:   "test@gmail.com",
		Age:     28,
		Phone:   1234567890,
		Gender:  "female",
		Disease: "Flu",
	}
	body, _ := json.Marshal(patient)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/patients", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	CreatePatientRecord(c)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", w.Code)
	}
	if !contains(w.Body.String(), "Successfully Created") {
		t.Errorf("Expected creation message, got: %s", w.Body.String())
	}

	database.DB.Where("email = ?", patient.Email).Delete(&models.Patient{})
}

func TestGetPatientRecords(t *testing.T) {
	p := models.Patient{
		Name:    "Test2",
		Email:   "test2@gmail.com",
		Age:     30,
		Phone:   9988776655,
		Gender:  "male",
		Disease: "Cold",
	}
	database.DB.Create(&p)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	GetPatientRecords(c)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", w.Code)
	}
	if !contains(w.Body.String(), p.Email) {
		t.Errorf("Expected to find patient, got: %s", w.Body.String())
	}

	database.DB.Delete(&p)
}

func TestUpdatePatientRecord(t *testing.T) {
	p := models.Patient{
		Name:    "Test3",
		Email:   "test3@gmail.com",
		Age:     45,
		Phone:   0000000000,
		Gender:  "male",
		Disease: "Cough",
	}
	database.DB.Create(&p)

	updated := models.Patient{
		Name:    "Test3",
		Email:   "test3@gmail.com",
		Age:     50,
		Phone:   1111111111,
		Gender:  "male",
		Disease: "Cold",
	}
	body, _ := json.Marshal(updated)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: strconv.Itoa(int(p.ID))}}
	c.Request = httptest.NewRequest(
		"PUT",
		"/patients/"+strconv.Itoa(int(p.ID)),
		bytes.NewBuffer(body),
	)
	c.Request.Header.Set("Content-Type", "application/json")

	UpdatePatientRecord(c)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", w.Code)
	}
	if !contains(w.Body.String(), "updated successfully") {
		t.Errorf("Expected update confirmation, got: %s", w.Body.String())
	}

	database.DB.Delete(&p)
}

func TestDeletePatientRecord(t *testing.T) {
	p := models.Patient{
		Name:    "Test4",
		Email:   "test4@gmail.com",
		Age:     33,
		Phone:   1234509876,
		Gender:  "male",
		Disease: "dengu",
	}
	database.DB.Create(&p)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: strconv.Itoa(int(p.ID))}}

	DeletePatientRecord(c)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", w.Code)
	}
	if !contains(w.Body.String(), "deleted") {
		t.Errorf("Expected deletion message, got: %s", w.Body.String())
	}
}
