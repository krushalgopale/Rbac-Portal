package database

import (
	"os"
	"testing"

	"github.com/krushalgopale/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func resetDB() {
	DB = nil
}

func TestConnDB_Success(t *testing.T) {
	resetDB()

	testDBURL := os.Getenv("DB_URL")
	if testDBURL == "" {
		t.Skip("DB_URL not set, skipping DB integration test")
	}

	os.Setenv("DB_URL", testDBURL)

	ConnDB()

	assert.NotNil(t, DB, "Expected DB to be initialized")
	assert.IsType(t, &gorm.DB{}, DB, "Expected DB to be of type *gorm.DB")

	ok := DB.Migrator().HasTable(&models.User{})
	assert.True(t, ok, "User table should be created")

	ok = DB.Migrator().HasTable(&models.Patient{})
	assert.True(t, ok, "Patient table should be created")
}
