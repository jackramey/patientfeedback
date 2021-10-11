package service

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
	patientsGroup := e.Group("/patients") // TODO add middleware here to require authN+authZ
	patientsGroup.GET("/", patientsHandler.GetAllPatients)
	patientsGroup.GET("/:patientId", patientsHandler.GetPatient)
	patientsGroup.GET("/:patientId/appointments", patientsHandler.GetAppointmentsForPatient)

	appointmentsHandler := appointmentsHandler{store: store}
	appointmentsGroup := e.Group("/appointments") // TODO add middleware here to require authN+authZ
	appointmentsGroup.GET("/:appointmentId", appointmentsHandler.GetAppointment)
	appointmentsGroup.GET("/:appointmentId/feedback", appointmentsHandler.GetFeedbackForAppointment)
	appointmentsGroup.POST("/:appointmentId/feedback", appointmentsHandler.PostFeedbackForAppointment)

	adminHandler := adminHandler{store: store}
	adminGroup := e.Group("/admin") // TODO add middleware here to require authN+authZ
	adminGroup.GET("/dumpdb", adminHandler.DumpDB)

	return e
}
