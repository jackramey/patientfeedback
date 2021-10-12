package storage

import (
	. "patientfeedback/internal/domain"

	"github.com/pkg/errors"
)

type ResourceStore interface {
	GetAppointment(id string) (*Appointment, error)
	GetDiagnosis(id string) (*Diagnosis, error)
	GetDoctor(id string) (*Doctor, error)
	GetFeedback(id string) (*Feedback, error)
	GetPatient(id string) (*Patient, error)

	GetAllPatients() ([]Patient, error)
	GetAppointmentsForPatient(patientId string) ([]Appointment, error)
	GetDiagnosisForAppointment(appointmentId string) (*Diagnosis, error)
	GetFeedbackForAppointment(appointmendId string) ([]Feedback, error)

	GetResourceByReference(ref Reference) (Resource, error)

	WriteAppointment(appointment Appointment) error
	WriteDiagnosis(diagnosis Diagnosis) error
	WriteDoctor(doctor Doctor) error
	WriteFeedback(feedback Feedback) error
	WritePatient(patient Patient) error

	DumpDB() Bundle
}

type MemoryStore struct {
	Appointments map[string]Appointment
	Diagnoses    map[string]Diagnosis
	Doctors      map[string]Doctor
	Feedback     map[string]Feedback
	Patients     map[string]Patient
}

type DBDump struct {
	Appointments []Appointment `json:"appointments"`
	Diagnoses    []Diagnosis   `json:"diagnoses"`
	Doctors      []Doctor      `json:"doctors"`
	Feedback     []Feedback    `json:"feedback"`
	Patients     []Patient     `json:"patients"`
}

func (m *MemoryStore) GetAppointment(id string) (*Appointment, error) {
	if len(m.Appointments) == 0 {
		return nil, nil
	}

	appointment, ok := m.Appointments[id]
	if !ok {
		return nil, nil
	}

	return &appointment, nil
}

func (m *MemoryStore) GetDiagnosis(id string) (*Diagnosis, error) {
	if len(m.Diagnoses) == 0 {
		return nil, nil
	}

	diagnosis, ok := m.Diagnoses[id]
	if !ok {
		return nil, nil
	}

	return &diagnosis, nil
}

func (m *MemoryStore) GetDoctor(id string) (*Doctor, error) {
	if len(m.Doctors) == 0 {
		return nil, nil
	}

	doctor, ok := m.Doctors[id]
	if !ok {
		return nil, nil
	}

	return &doctor, nil
}
func (m *MemoryStore) GetFeedback(id string) (*Feedback, error) {
	if len(m.Feedback) == 0 {
		return nil, nil
	}

	feedback, ok := m.Feedback[id]
	if !ok {
		return nil, nil
	}

	return &feedback, nil
}

func (m *MemoryStore) GetPatient(id string) (*Patient, error) {
	if len(m.Patients) == 0 {
		return nil, nil
	}

	patient, ok := m.Patients[id]
	if !ok {
		return nil, nil
	}

	return &patient, nil
}

func (m *MemoryStore) GetAllPatients() ([]Patient, error) {
	var patients []Patient
	for _, patient := range m.Patients {
		patients = append(patients, patient)
	}
	return patients, nil
}

func (m *MemoryStore) GetAppointmentsForPatient(patientId string) ([]Appointment, error) {
	// Terrible performance for large sets of appointments, but this gets the job done for a memory store
	// which we anticipate having a very small set. Ideally we'd use a DB like BuntDB if we anticipated needing
	// an in-memory database with complex functionality like secondary indexes
	var appointments []Appointment
	for _, appointment := range m.Appointments {
		if appointment.Subject.ID == patientId {
			appointments = append(appointments, appointment)
		}
	}
	return appointments, nil
}

func (m *MemoryStore) GetDiagnosisForAppointment(appointmentId string) (*Diagnosis, error) {
	for _, diagnosis := range m.Diagnoses {
		// Short circuit and exit since we're only allowing one feedback per appointment
		if diagnosis.Appointment.ID == appointmentId {
			return &diagnosis, nil
		}
	}
	return nil, nil
}

func (m *MemoryStore) GetFeedbackForAppointment(appointmendId string) ([]Feedback, error) {
	var feedbacks []Feedback
	for _, feedback := range m.Feedback {
		if feedback.Appointment.ID == appointmendId {
			feedbacks = append(feedbacks, feedback)
		}
	}
	return feedbacks, nil
}

func (m *MemoryStore) GetResourceByReference(ref Reference) (Resource, error) {
	switch ref.Type {
	case AppointmentResType:
		return m.GetAppointment(ref.ID)
	case DiagnosisResType:
		return m.GetDiagnosis(ref.ID)
	case DoctorResType:
		return m.GetDoctor(ref.ID)
	case FeedbackResType:
		return m.GetFeedback(ref.ID)
	case PatientResType:
		return m.GetPatient(ref.ID)
	default:
		return nil, errors.Errorf("resource type unsupported by store: %s", ref.Type)
	}
}

func (m *MemoryStore) WriteAppointment(appointment Appointment) error {
	if m.Appointments == nil {
		m.Appointments = map[string]Appointment{}
	}

	m.Appointments[appointment.ID] = appointment
	return nil
}

func (m *MemoryStore) WriteDiagnosis(diagnosis Diagnosis) error {
	if m.Diagnoses == nil {
		m.Diagnoses = map[string]Diagnosis{}
	}

	m.Diagnoses[diagnosis.ID] = diagnosis
	return nil
}

func (m *MemoryStore) WriteDoctor(doctor Doctor) error {
	if m.Doctors == nil {
		m.Doctors = map[string]Doctor{}
	}

	m.Doctors[doctor.ID] = doctor
	return nil
}

func (m *MemoryStore) WriteFeedback(feedback Feedback) error {
	if m.Feedback == nil {
		m.Feedback = map[string]Feedback{}
	}

	m.Feedback[feedback.ID] = feedback
	return nil
}

func (m *MemoryStore) WritePatient(patient Patient) error {
	if m.Patients == nil {
		m.Patients = map[string]Patient{}
	}

	m.Patients[patient.ID] = patient
	return nil
}

func (m *MemoryStore) DumpDB() Bundle {
	bundle := NewBundle()

	for _, val := range m.Appointments {
		bundle.Entries = append(bundle.Entries, Entry{Resource: val})
	}
	for _, val := range m.Diagnoses {
		bundle.Entries = append(bundle.Entries, Entry{Resource: val})
	}
	for _, val := range m.Doctors {
		bundle.Entries = append(bundle.Entries, Entry{Resource: val})
	}
	for _, val := range m.Feedback {
		bundle.Entries = append(bundle.Entries, Entry{Resource: val})
	}
	for _, val := range m.Patients {
		bundle.Entries = append(bundle.Entries, Entry{Resource: val})
	}

	return bundle
}
