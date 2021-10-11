package storage

import (
	"patientfeedback/internal/domain"

	"github.com/pkg/errors"
)

type ResourceStore interface {
	GetAppointment(id string) (*domain.Appointment, error)
	GetDiagnosis(id string) (*domain.Diagnosis, error)
	GetDoctor(id string) (*domain.Doctor, error)
	GetFeedback(id string) (*domain.Feedback, error)
	GetPatient(id string) (*domain.Patient, error)

	GetAllPatients() ([]domain.Patient, error)
	GetAppointmentsForPatient(patientId string) ([]domain.Appointment, error)
	GetDiagnosisForAppointment(appointmentId string) (*domain.Diagnosis, error)
	GetFeedbackForAppointment(appointmendId string) ([]domain.Feedback, error)

	GetResourceByReference(ref domain.Reference) (domain.Resource, error)

	WriteAppointment(appointment domain.Appointment) error
	WriteDiagnosis(diagnosis domain.Diagnosis) error
	WriteDoctor(doctor domain.Doctor) error
	WriteFeedback(feedback domain.Feedback) error
	WritePatient(patient domain.Patient) error

	DumpDB() DBDump
}

type MemoryStore struct {
	Appointments map[string]domain.Appointment
	Diagnoses    map[string]domain.Diagnosis
	Doctors      map[string]domain.Doctor
	Feedback     map[string]domain.Feedback
	Patients     map[string]domain.Patient
}

type DBDump struct {
	Appointments []domain.Appointment `json:"appointments"`
	Diagnoses    []domain.Diagnosis   `json:"diagnoses"`
	Doctors      []domain.Doctor      `json:"doctors"`
	Feedback     []domain.Feedback    `json:"feedback"`
	Patients     []domain.Patient     `json:"patients"`
}

func (m *MemoryStore) GetAppointment(id string) (*domain.Appointment, error) {
	if len(m.Appointments) == 0 {
		return nil, nil
	}

	appointment, ok := m.Appointments[id]
	if !ok {
		return nil, nil
	}

	return &appointment, nil
}

func (m *MemoryStore) GetDiagnosis(id string) (*domain.Diagnosis, error) {
	if len(m.Diagnoses) == 0 {
		return nil, nil
	}

	diagnosis, ok := m.Diagnoses[id]
	if !ok {
		return nil, nil
	}

	return &diagnosis, nil
}

func (m *MemoryStore) GetDoctor(id string) (*domain.Doctor, error) {
	if len(m.Doctors) == 0 {
		return nil, nil
	}

	doctor, ok := m.Doctors[id]
	if !ok {
		return nil, nil
	}

	return &doctor, nil
}
func (m *MemoryStore) GetFeedback(id string) (*domain.Feedback, error) {
	if len(m.Feedback) == 0 {
		return nil, nil
	}

	feedback, ok := m.Feedback[id]
	if !ok {
		return nil, nil
	}

	return &feedback, nil
}

func (m *MemoryStore) GetPatient(id string) (*domain.Patient, error) {
	if len(m.Patients) == 0 {
		return nil, nil
	}

	patient, ok := m.Patients[id]
	if !ok {
		return nil, nil
	}

	return &patient, nil
}

func (m *MemoryStore) GetAllPatients() ([]domain.Patient, error) {
	var patients []domain.Patient
	for _, patient := range m.Patients {
		patients = append(patients, patient)
	}
	return patients, nil
}

func (m *MemoryStore) GetAppointmentsForPatient(patientId string) ([]domain.Appointment, error) {
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

func (m *MemoryStore) GetDiagnosisForAppointment(appointmentId string) (*domain.Diagnosis, error) {
	for _, diagnosis := range m.Diagnoses {
		// Short circuit and exit since we're only allowing one feedback per appointment
		if diagnosis.Appointment.ID == appointmentId {
			return &diagnosis, nil
		}
	}
	return nil, nil
}

func (m *MemoryStore) GetFeedbackForAppointment(appointmendId string) ([]domain.Feedback, error) {
	var feedbacks []domain.Feedback
	for _, feedback := range m.Feedback {
		if feedback.Appointment.ID == appointmendId {
			feedbacks = append(feedbacks, feedback)
		}
	}
	return feedbacks, nil
}

func (m *MemoryStore) GetResourceByReference(ref domain.Reference) (domain.Resource, error) {
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

func (m *MemoryStore) WriteAppointment(appointment domain.Appointment) error {
	if m.Appointments == nil {
		m.Appointments = map[string]domain.Appointment{}
	}

	m.Appointments[appointment.ID] = appointment
	return nil
}

func (m *MemoryStore) WriteDiagnosis(diagnosis domain.Diagnosis) error {
	if m.Diagnoses == nil {
		m.Diagnoses = map[string]domain.Diagnosis{}
	}

	m.Diagnoses[diagnosis.ID] = diagnosis
	return nil
}

func (m *MemoryStore) WriteDoctor(doctor domain.Doctor) error {
	if m.Doctors == nil {
		m.Doctors = map[string]domain.Doctor{}
	}

	m.Doctors[doctor.ID] = doctor
	return nil
}

func (m *MemoryStore) WriteFeedback(feedback domain.Feedback) error {
	if m.Feedback == nil {
		m.Feedback = map[string]domain.Feedback{}
	}

	m.Feedback[feedback.ID] = feedback
	return nil
}

func (m *MemoryStore) WritePatient(patient domain.Patient) error {
	if m.Patients == nil {
		m.Patients = map[string]domain.Patient{}
	}

	m.Patients[patient.ID] = patient
	return nil
}

func (m *MemoryStore) DumpDB() DBDump {
	var appointments []domain.Appointment
	var diagnoses []domain.Diagnosis
	var doctors []domain.Doctor
	var feedback []domain.Feedback
	var patients []domain.Patient

	for _, val := range m.Appointments {
		appointments = append(appointments, val)
	}
	for _, val := range m.Diagnoses {
		diagnoses = append(diagnoses, val)
	}
	for _, val := range m.Doctors {
		doctors = append(doctors, val)
	}
	for _, val := range m.Feedback {
		feedback = append(feedback, val)
	}
	for _, val := range m.Patients {
		patients = append(patients, val)
	}

	return DBDump{
		Appointments: appointments,
		Diagnoses:    diagnoses,
		Doctors:      doctors,
		Feedback:     feedback,
		Patients:     patients,
	}
}
