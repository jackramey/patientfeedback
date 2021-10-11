//go:build integration

package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"patientfeedback/api"
	"patientfeedback/storage"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var patientId = "6739ec3e-93bd-11eb-a8b3-0242ac130003"

func Test_patientsHandler_GetAppointmentsForPatient(t *testing.T) {
	store := storage.MemoryStore{}
	fileReader, err := os.Open("../data/bundle.json")
	require.NoError(t, err)

	err = storage.LoadData(fileReader, &store)
	require.NoError(t, err)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/patients/"+patientId+"/appointments", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/patients/:patientId/appointments")
	c.SetParamNames("patientId")
	c.SetParamValues(patientId)

	handler := patientsHandler{store: &store}
	err = handler.GetAppointmentsForPatient(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response api.GetAppointmentsResponse
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
	assert.Len(t, response.Appointments, 1)

	appt := response.Appointments[0]
	assert.Equal(t, appt.ID, "be142dc6-93bd-11eb-a8b3-0242ac130003")
	assert.Equal(t, appt.Doctor, "Careful")
	assert.Equal(t, appt.Diagnosis, "Diabetes without complications")
}
