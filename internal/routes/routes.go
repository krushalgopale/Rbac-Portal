package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/krushalgopale/internal/controllers/auth"
	"github.com/krushalgopale/internal/controllers/doctor"
	"github.com/krushalgopale/internal/controllers/receptionist"
	"github.com/krushalgopale/internal/middlewares"
)

func Routes(r *gin.Engine) {

	// Auth routes
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/register", auth.Register)
	}

	// Protected routes requiring authentication
	protectedGroup := r.Group("/")
	protectedGroup.Use(middlewares.AuthMiddleware)
	{

		// Receptionist routes
		receptionistGroup := protectedGroup.Group("/receptionist")
		receptionistGroup.Use(middlewares.ReceptionistMiddleware)
		{
			receptionistGroup.GET("/dashboard", receptionist.ReceptionistController)
			receptionistGroup.POST("/createpatient", receptionist.CreatePatientRecord)
			receptionistGroup.GET("/patients", receptionist.GetPatientRecords)
			receptionistGroup.PUT("/patient/:id", receptionist.UpdatePatientRecord)
			receptionistGroup.DELETE("/patient/:id", receptionist.DeletePatientRecord)
		}

		// Doctor routes
		doctorGroup := protectedGroup.Group("/doctor")
		doctorGroup.Use(middlewares.DoctorMiddleware)
		{
			doctorGroup.GET("/dashboard", doctor.DoctorController)
			doctorGroup.GET("/patients", doctor.GetPatientRecords)
			doctorGroup.PUT("/patient/:id", doctor.UpdatePatientRecord)
		}
	}
}
