package storage

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"

	"patientfeedback/internal/domain"
)

func LoadData(reader io.Reader, store ResourceStore) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.Wrap(err, "error reading data")
	}

	var bundle domain.Bundle
	if err := json.Unmarshal(data, &bundle); err != nil {
		return errors.Wrap(err, "unable to unmarshal json data input")
	}

	for _, entry := range bundle.Entries {
		entry := entry
		resource := entry.Resource
		switch resource.GetResourceType() {
		case domain.AppointmentResType:
			if err := store.WriteAppointment(*resource.(*domain.Appointment)); err != nil {
				return err
			}
		case domain.DiagnosisResType:
			if err := store.WriteDiagnosis(*resource.(*domain.Diagnosis)); err != nil {
				return err
			}
		case domain.DoctorResType:
			if err := store.WriteDoctor(*resource.(*domain.Doctor)); err != nil {
				return err
			}
		case domain.PatientResType:
			if err := store.WritePatient(*resource.(*domain.Patient)); err != nil {
				return err
			}
		default:
			return errors.New("unsupported resource type found")
		}
	}

	return nil
}
