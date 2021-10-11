package storage

import (
	"testing"

	"patientfeedback/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_memoryStore_GetAllPatients(t *testing.T) {
	t.Run("should return all patients", func(t *testing.T) {
		memStore := MemoryStore{}
		require.NoError(t, memStore.WritePatient(domain.Patient{
			ResourceHeader: domain.ResourceHeader{
				ID: "1",
			},
		}))
		require.NoError(t, memStore.WritePatient(domain.Patient{
			ResourceHeader: domain.ResourceHeader{
				ID: "2",
			},
		}))
		patients, err := memStore.GetAllPatients()
		require.NoError(t, err)
		assert.Len(t, patients, 2)
	})
}

func TestMemoryStore_GetAppointmentsForPatient(t *testing.T) {
	store := MemoryStore{}
	patientId := uuid.New().String()
	appointment := domain.Appointment{
		ResourceHeader: domain.ResourceHeader{
			ID:   uuid.New().String(),
			Type: domain.AppointmentResType,
		},
		Subject: domain.Reference{
			Type: domain.PatientResType,
			ID:   patientId,
		},
	}
	require.NoError(t, store.WriteAppointment(appointment))

	t.Run("gets appointment when subject is patientId", func(t *testing.T) {
		result, err := store.GetAppointmentsForPatient(patientId)
		require.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, result[0], appointment)
	})
	t.Run("returns an empty set if no appointments exist for patient", func(t *testing.T) {
		result, err := store.GetAppointmentsForPatient("nobody")
		require.NoError(t, err)
		assert.Empty(t, result)
	})
}

func TestMemoryStore_GetResourceByReference(t *testing.T) {
	m := MemoryStore{}
	appointment := domain.Appointment{
		ResourceHeader: domain.ResourceHeader{
			ID:   uuid.New().String(),
			Type: domain.AppointmentResType,
		},
	}
	require.NoError(t, m.WriteAppointment(appointment))
	diagnosis := domain.Diagnosis{
		ResourceHeader: domain.ResourceHeader{
			ID:   uuid.New().String(),
			Type: domain.DiagnosisResType,
		},
	}
	require.NoError(t, m.WriteDiagnosis(diagnosis))
	doctor := domain.Doctor{
		ResourceHeader: domain.ResourceHeader{
			ID:   uuid.New().String(),
			Type: domain.DoctorResType,
		},
	}
	require.NoError(t, m.WriteDoctor(doctor))
	feedback := domain.Feedback{
		ResourceHeader: domain.ResourceHeader{
			ID:   uuid.New().String(),
			Type: domain.FeedbackResType,
		},
	}
	require.NoError(t, m.WriteFeedback(feedback))
	patient := domain.Patient{
		ResourceHeader: domain.ResourceHeader{
			ID:   uuid.New().String(),
			Type: domain.PatientResType,
		},
	}
	require.NoError(t, m.WritePatient(patient))

	tests := []struct {
		name     string
		ref      domain.Reference
		wantType interface{}
		wantErr  bool
	}{
		{
			name: "Get Appointment by reference",
			ref: domain.Reference{
				Type: domain.AppointmentResType,
				ID:   appointment.ID,
			},
			wantType: &domain.Appointment{},
			wantErr:  false,
		},
		{
			name: "Get Diagnosis by reference",
			ref: domain.Reference{
				Type: domain.DiagnosisResType,
				ID:   diagnosis.ID,
			},
			wantType: &domain.Diagnosis{},
			wantErr:  false,
		},
		{
			name: "Get Doctor by reference",
			ref: domain.Reference{
				Type: domain.DoctorResType,
				ID:   doctor.ID,
			},
			wantType: &domain.Doctor{},
			wantErr:  false,
		},
		{
			name: "Get Feedback by reference",
			ref: domain.Reference{
				Type: domain.FeedbackResType,
				ID:   feedback.ID,
			},
			wantType: &domain.Feedback{},
			wantErr:  false,
		},
		{
			name: "Get Patient by reference",
			ref: domain.Reference{
				Type: domain.PatientResType,
				ID:   patient.ID,
			},
			wantType: &domain.Patient{},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.GetResourceByReference(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetResourceByReference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.IsType(t, tt.wantType, got) {
				t.Errorf("GetResourceByReference() got = %v, want %v", got, tt.wantType)
			}
		})
	}
}
