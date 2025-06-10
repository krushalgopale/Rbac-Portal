package models

import (
	"testing"
)

func TestPatientCreation(t *testing.T) {
	p := Patient{
		ID:      1,
		Name:    "Saitama",
		Email:   "saitama@gmail.com",
		Phone:   9876543210,
		Age:     28,
		Gender:  "Male",
		Disease: "Diabetes",
	}

	if p.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", p.ID)
	}

	if p.Name != "Saitama" {
		t.Errorf("Expected Name to be 'Saitama', got %s", p.Name)
	}

	if p.Email != "saitama@gmail.com" {
		t.Errorf("Expected Email to be 'aryan.blaze@example.com', got %s", p.Email)
	}

	if p.Phone != 9876543210 {
		t.Errorf("Expected Phone to be 9876543210, got %d", p.Phone)
	}

	if p.Age != 28 {
		t.Errorf("Expected Age to be 28, got %d", p.Age)
	}

	if p.Gender != "Male" {
		t.Errorf("Expected Gender to be 'Male', got %s", p.Gender)
	}

	if p.Disease != "Diabetes" {
		t.Errorf("Expected Disease to be 'Diabetes', got %s", p.Disease)
	}
}
