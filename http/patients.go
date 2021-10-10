package http

import (
	"net/http"

	"patientfeedback/api"
	. "patientfeedback/domain"
	"patientfeedback/storage"

	"github.com/labstack/echo/v4"
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

	return c.JSON(http.StatusOK, patient)
}

type GetPatientsResponse struct {
	Patients []Patient `json:"patients"`
}

func (handler patientsHandler) GetAllPatients(c echo.Context) error {
	patients, err := handler.store.GetAllPatients()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, GetPatientsResponse{
		Patients: patients,
	})
}

type GetAppointmentsResponse struct {
	Appointments []api.Appointment `json:"appointments"`
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

	return c.JSON(http.StatusOK, GetAppointmentsResponse{
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
