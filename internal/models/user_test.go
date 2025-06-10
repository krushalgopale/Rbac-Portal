package models

import (
	"testing"
)

func TestUserInitialization(t *testing.T) {
	user := User{
		ID:       1,
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: "test123",
		Role:     "doctor",
	}

	if user.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", user.ID)
	}

	if user.Name != "Test" {
		t.Errorf("Expected Name to be 'Test User', got %s", user.Name)
	}

	if user.Email != "test@gmail.com" {
		t.Errorf("Expected Email to be 'user@example.com', got %s", user.Email)
	}

	if user.Password != "test123" {
		t.Errorf("Expected Password to be 'secret123', got %s", user.Password)
	}

	if user.Role != "doctor" {
		t.Errorf("Expected Role to be 'doctor', got %s", user.Role)
	}
}
