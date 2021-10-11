package service

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"patientfeedback/api"
	. "patientfeedback/internal/domain"
	"patientfeedback/storage"
)

type patientsHandler struct {
	store storage.ResourceStore
}

func (handler patientsHandler) GetPatient(c echo.Context) error {
	patientId := c.Param("patientId")
	patient, err := handler.store.GetPatient(patientId)
	if err != nil {
		return err
	}

	fullName := PatientPreferredFullName(*patient)
	return c.JSON(http.StatusOK, api.Patient{
		ID:   patient.ID,
		Name: PatientPreferredFirstName(*patient),
		FullName: fullName,
	})
}

func (handler patientsHandler) GetAllPatients(c echo.Context) error {
	patients, err := handler.store.GetAllPatients()
	if err != nil {
		return err
	}

	var apiPatients []api.Patient
	for _, patient := range patients {
		apiPatients = append(apiPatients, api.Patient{
			ID:   patient.ID,
			Name: PatientPreferredFirstName(patient),
			FullName: PatientPreferredFullName(patient),
		})
	}

	return c.JSON(http.StatusOK, api.GetPatientsResponse{
		Patients: apiPatients,
	})
}

func (handler patientsHandler) GetAppointmentsForPatient(c echo.Context) error {
	patientId := c.Param("patientId")
	patient, err := handler.store.GetPatient(patientId)
	if err != nil {
		return err
	}

	if patient == nil {
		return echo.ErrNotFound
	}

	appointments, err := handler.store.GetAppointmentsForPatient(patientId)
	if err != nil {
		return err
	}

	simpleAppointments, err := handler.simplifyAppointments(appointments)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, api.GetAppointmentsResponse{
		Appointments: simpleAppointments,
	})
}

func (handler patientsHandler) makeAPIAppointment(appointment Appointment) (api.Appointment, error) {
	doctor, err := handler.store.GetDoctor(appointment.Actor.ID)
	if err != nil {
		return api.Appointment{}, err
	}

	diagnosis, err := handler.store.GetDiagnosisForAppointment(appointment.ID)

	return api.Appointment{
		ID:        appointment.ID,
		Summary:   AppointmentSummary(appointment),
		Doctor:    PreferredLastName(*doctor),
		Diagnosis: DiagnosisText(*diagnosis),
	}, nil
}

func (handler patientsHandler) simplifyAppointments(appointments []Appointment) ([]api.Appointment, error) {
	var simpleAppointments []api.Appointment
	for _, appointment := range appointments {
		simpleAppointment, err := handler.makeAPIAppointment(appointment)
		if err != nil {
			return nil, err
		}
		simpleAppointments = append(simpleAppointments, simpleAppointment)
	}
	return simpleAppointments, nil
}
