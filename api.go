package patientfeedback

import (
	"context"

	. "patientfeedback/domain"
	. "patientfeedback/storage"
)

type Controller struct {
	store ResourceStore
}

func (c Controller) GetPatients(ctx context.Context) ([]Patient, error) {

	return nil, nil
}
