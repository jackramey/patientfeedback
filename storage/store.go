package storage

import (
	"patientfeedback/domain"

	"github.com/pkg/errors"
)

type ResourceStore interface {
	GetAppointment(id string) (*domain.Appointment, error)
	GetDiagnosis(id string) (*domain.Diagnosis, error)
	GetDoctor(id string) (*domain.Doctor, error)
	GetFeedback(id string) (*domain.Feedback, error)
	GetPatient(id string) (*domain.Patient, error)

	GetAllPatients(id string) ([]domain.Patient, error)
	GetAppointmentsForPatient(patientId string) ([]domain.Appointment, error)
	GetResourceByReference(ref domain.Reference) (domain.Resource, error)

	WriteAppointment(appointment domain.Appointment) error
	WriteDiagnosis(diagnosis domain.Diagnosis) error
	WriteDoctor(doctor domain.Doctor) error
	WriteFeedback(feedback domain.Feedback) error
	WritePatient(patient domain.Patient) error
}

type memoryStore struct {
	Appointments map[string]domain.Appointment
	Diagnoses    map[string]domain.Diagnosis
	Doctors      map[string]domain.Doctor
	Feedback     map[string]domain.Feedback
	Patients     map[string]domain.Patient
}

func (m *memoryStore) GetAppointment(id string) (*domain.Appointment, error) {
	if len(m.Appointments) == 0 {
		return nil, nil
	}

	appointment, ok := m.Appointments[id]
	if !ok {
		return nil, nil
	}

	return &appointment, nil
}

func (m *memoryStore) GetDiagnosis(id string) (*domain.Diagnosis, error) {
	if len(m.Diagnoses) == 0 {
		return nil, nil
	}

	diagnosis, ok := m.Diagnoses[id]
	if !ok {
		return nil, nil
	}

	return &diagnosis, nil
}

func (m *memoryStore) GetDoctor(id string) (*domain.Doctor, error) {
	if len(m.Doctors) == 0 {
		return nil, nil
	}

	doctor, ok := m.Doctors[id]
	if !ok {
		return nil, nil
	}

	return &doctor, nil
}
func (m *memoryStore) GetFeedback(id string) (*domain.Feedback, error) {
	if len(m.Feedback) == 0 {
		return nil, nil
	}

	feedback, ok := m.Feedback[id]
	if !ok {
		return nil, nil
	}

	return &feedback, nil
}

func (m *memoryStore) GetPatient(id string) (*domain.Patient, error) {
	if len(m.Patients) == 0 {
		return nil, nil
	}

	patient, ok := m.Patients[id]
	if !ok {
		return nil, nil
	}

	return &patient, nil
}

func (m *memoryStore) GetAllPatients(id string) ([]domain.Patient, error) {
	var patients []domain.Patient
	for _, patient := range m.Patients {
		patients = append(patients, patient)
	}
	return patients, nil
}

func (m *memoryStore) GetAppointmentsForPatient(patientId string) ([]domain.Appointment, error) {
	// Terrible performance for large sets of appointments, but this gets the job done for a memory store
	// which we anticipate having a very small set. Ideally we'd use a DB like BuntDB if we anticipated needing
	// an in-memory database with complex functionality like secondary indexes
	var appointments []domain.Appointment
	for _, appointment := range m.Appointments {
		if appointment.Subject.ID == patientId {
			appointments = append(appointments, appointment)
		}
	}
	return appointments, nil
}

func (m *memoryStore) GetResourceByReference(ref domain.Reference) (domain.Resource, error) {
	switch ref.Type {
	case domain.AppointmentResType:
		return m.GetAppointment(ref.ID)
	case domain.DiagnosisResType:
		return m.GetDiagnosis(ref.ID)
	case domain.DoctorResType:
		return m.GetDoctor(ref.ID)
	case domain.FeedbackResType:
		return m.GetFeedback(ref.ID)
	case domain.PatientResType:
		return m.GetPatient(ref.ID)
	default:
		return nil, errors.Errorf("resource type unsupported by store: %s", ref.Type)
	}
}

func (m *memoryStore) WriteAppointment(appointment domain.Appointment) error {
	if m.Appointments == nil {
		m.Appointments = map[string]domain.Appointment{}
	}

	m.Appointments[appointment.ID] = appointment
	return nil
}

func (m *memoryStore) WriteDiagnosis(diagnosis domain.Diagnosis) error {
	if m.Diagnoses == nil {
		m.Diagnoses = map[string]domain.Diagnosis{}
	}

	m.Diagnoses[diagnosis.ID] = diagnosis
	return nil
}

func (m *memoryStore) WriteDoctor(doctor domain.Doctor) error {
	if m.Doctors == nil {
		m.Doctors = map[string]domain.Doctor{}
	}

	m.Doctors[doctor.ID] = doctor
	return nil
}

func (m *memoryStore) WriteFeedback(feedback domain.Feedback) error {
	if m.Feedback == nil {
		m.Feedback = map[string]domain.Feedback{}
	}

	m.Feedback[feedback.ID] = feedback
	return nil
}

func (m *memoryStore) WritePatient(patient domain.Patient) error {
	if m.Patients == nil {
		m.Patients = map[string]domain.Patient{}
	}

	m.Patients[patient.ID] = patient
	return nil
}
