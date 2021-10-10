package http

import (
	"patientfeedback/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEchoServer(store storage.ResourceStore) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	patientsHandler := patientsHandler{store: store}
	patientsGroup := e.Group("/patients")
	patientsGroup.GET("/", patientsHandler.GetAllPatients)
	patientsGroup.GET("/:patientId", patientsHandler.GetPatient)
	patientsGroup.GET("/:patientId/appointments", patientsHandler.GetAppointmentsForPatient)

	appointmentsHandler := appointmentsHandler{store: store}
	appointmentsGroup := e.Group("/appointments")
	appointmentsGroup.GET("/:appointmentId", appointmentsHandler.GetAppointment)
	appointmentsGroup.POST("/:appointmentId/feedback", appointmentsHandler.PostFeedbackForAppointment)

	return e
}
