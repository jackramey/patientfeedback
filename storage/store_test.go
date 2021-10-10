package storage

import (
	"testing"

	"patientfeedback/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_memoryStore_GetResourceByReference(t *testing.T) {
	store := memoryStore{}
	t.Run("get appointment by reference", func(t *testing.T) {
		appointmentID := uuid.New().String()
		err := store.WriteAppointment(domain.Appointment{
			ResourceHeader: domain.ResourceHeader{
				ID: appointmentID,
			},
		})
		require.NoError(t, err)

		result, err := store.GetResourceByReference(domain.Reference{
			Type: domain.AppointmentResType,
			ID:   appointmentID,
		})
		require.NoError(t, err)
		assert.IsType(t, &domain.Appointment{}, result)

	})
}
