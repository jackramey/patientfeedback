package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"patientfeedback/api"
)

const serviceHost = "http://localhost:1323"

type PatientClient struct {
	doer      Doer
}

func NewPatientClient() PatientClient {
	return PatientClient{
		doer:      http.DefaultClient,
	}
}

func (c PatientClient) GetAllPatients() ([]api.Patient, error) {
	req, err := http.Get(serviceHost + "/patients/")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var resp api.GetPatientsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		panic(err)
	}

	return resp.Patients, nil
}

func (c PatientClient) GetPatientInfo(patientId string) (api.Patient, error) {
	req, err := http.Get(serviceHost + "/patients/" + patientId)
	if err != nil {
		fmt.Printf("did you forget to start the server?")
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var patient api.Patient
	if err := json.Unmarshal(body, &patient); err != nil {
		panic(err)
	}

	return patient, nil
}

func (c PatientClient) GetAppointmentsForPatient(patientId string) ([]api.Appointment, error) {
	req, err := http.NewRequest(http.MethodGet, serviceHost+"/patients/"+patientId+"/appointments", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.doer.Do(req)
	if err != nil {
		return nil, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var out api.GetAppointmentsResponse
	if err := json.Unmarshal(respBytes, &out); err != nil {
		return nil, err
	}

	return out.Appointments, nil
}

func (c PatientClient) WriteFeedback(appointementId string, f api.Feedback) error {
	body, err := json.Marshal(f)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		serviceHost+"/appointments/"+appointementId+"/feedback",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	_, err = c.doer.Do(req)
	if err != nil {
		return err
	}

	return nil
}

type Doer interface {
	Do(r *http.Request) (*http.Response, error)
}
