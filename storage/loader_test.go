package storage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoader_LoadData(t *testing.T) {
	reader, err := os.Open("../data/patient-feedback-raw-data.json")
	require.NoError(t, err)
	memStore := &memoryStore{}

	loader := Loader{
		resourceStore: memStore,
	}

	err = loader.LoadData(reader)
	require.NoError(t, err)

	assert.Len(t, memStore.Patients, 1)
	assert.Len(t, memStore.Doctors, 1)
	assert.Len(t, memStore.Appointments, 1)
	assert.Len(t, memStore.Diagnoses, 1)
}
