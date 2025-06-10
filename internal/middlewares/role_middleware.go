package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Receptionist middleware
func ReceptionistMiddleware(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
		c.Abort()
		return
	}

	if role != "receptionist" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. Receptionist role required"})
		c.Abort()
		return
	}
	c.Next()
}

// Doctor middleware
func DoctorMiddleware(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
		c.Abort()
		return
	}

	if role != "doctor" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. Doctor role required"})
		c.Abort()
		return
	}
	c.Next()
}
