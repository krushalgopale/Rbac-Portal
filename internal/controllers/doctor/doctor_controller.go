package doctor

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krushalgopale/internal/database"
	"github.com/krushalgopale/internal/models"
)

func DoctorController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Doctor Dashboard"})
}

// Get patient records
func GetPatientRecords(c *gin.Context) {
	var records []models.Patient

	if err := database.DB.Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get patient record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Records of patients": records})
}

// Get patient record by patient-id
func GetPatientRecordById(c *gin.Context) {
	var record models.Patient

	patientID := c.Param("id")

	result := database.DB.First(&record, patientID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"Record of patient": record})

}

// Update patient records
func UpdatePatientRecord(c *gin.Context) {
	var record models.Patient

	patientID := c.Param("id")

	result := database.DB.First(&record, patientID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var updatedRecord models.Patient
	if err := c.ShouldBindJSON(&updatedRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	record.Name = updatedRecord.Name
	record.Email = updatedRecord.Email
	record.Age = updatedRecord.Age
	record.Phone = updatedRecord.Phone
	record.Gender = updatedRecord.Gender
	record.Disease = updatedRecord.Disease

	if err := database.DB.Save(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient record updated successfully"})
}
