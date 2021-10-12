package service

import (
	"fmt"

	"patientfeedback/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const defaultPort = 1323

type Server struct {
	*echo.Echo

	config Config
}

type Config struct {
	Store storage.ResourceStore
	Port int
}

func NewServer(config Config) Server {
	server := Server{
		config: config,
		Echo: echo.New(),
	}

	// Middleware
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	patientsHandler := patientsHandler{store: config.Store}
	patientsGroup := server.Group("/patients") // TODO add middleware here to require authN+authZ
	patientsGroup.GET("/", patientsHandler.GetAllPatients)
	patientsGroup.GET("/:patientId", patientsHandler.GetPatient)
	patientsGroup.GET("/:patientId/appointments", patientsHandler.GetAppointmentsForPatient)

	appointmentsHandler := appointmentsHandler{store: config.Store}
	appointmentsGroup := server.Group("/appointments") // TODO add middleware here to require authN+authZ
	appointmentsGroup.GET("/:appointmentId", appointmentsHandler.GetAppointment)
	appointmentsGroup.GET("/:appointmentId/feedback", appointmentsHandler.GetFeedbackForAppointment)
	appointmentsGroup.POST("/:appointmentId/feedback", appointmentsHandler.PostFeedbackForAppointment)

	securedHandler := securedHandler{store: config.Store}
	securedGroup := server.Group("/secured") // TODO add middleware here to require authN+authZ
	securedGroup.GET("/dumpdb", securedHandler.DumpDB)
	securedGroup.POST("/bundle", securedHandler.ProcessBundle)

	return server
}

func (s *Server) Run() error {
	var port = s.config.Port
	// Well known ports for Unix systems are ports 0-1023
	// Ports over 65535 are invalid
	if s.config.Port < 1023 || s.config.Port > 65535 {
		fmt.Printf("invalid port in configuration: %d\n", s.config.Port)
		port = defaultPort
	}

	return s.Start(fmt.Sprintf(":%d", port))
}
