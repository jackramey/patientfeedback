package storage

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"patientfeedback/domain"

	"github.com/pkg/errors"
)

type Loader struct {
	resourceStore ResourceStore
}

func (l Loader) LoadData(reader io.Reader) error {
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
			if err := l.resourceStore.WriteAppointment(*resource.(*domain.Appointment)); err != nil {
				return err
			}
		case domain.DiagnosisResType:
			if err := l.resourceStore.WriteDiagnosis(*resource.(*domain.Diagnosis)); err != nil {
				return err
			}
		case domain.DoctorResType:
			if err := l.resourceStore.WriteDoctor(*resource.(*domain.Doctor)); err != nil {
				return err
			}
		case domain.PatientResType:
			if err := l.resourceStore.WritePatient(*resource.(*domain.Patient)); err != nil {
				return err
			}
		default:
			return errors.New("unsupported resource type found")
		}
	}

	return nil
}
