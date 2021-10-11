package domain

import (
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

func (e *Entry) UnmarshalJSON(bytes []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	resourceJSON, ok := raw["resource"]
	if !ok {
		return errors.New("entry has no resource attribute")
	}

	resource, err := unmarshalResource(resourceJSON)
	if err != nil {
		return err
	}

	e.Resource = resource
	return nil
}

func unmarshalResource(bytes []byte) (Resource, error) {
	var header ResourceHeader
	if err := json.Unmarshal(bytes, &header); err != nil {
		return nil, err
	}

	var out Resource
	switch header.Type {
	case AppointmentResType:
		out = &Appointment{}
	case DiagnosisResType:
		out = &Diagnosis{}
	case DoctorResType:
		out = &Doctor{}
	case PatientResType:
		out = &Patient{}
	default:
		return nil, errors.New("unknown resource type")
	}

	if err := json.Unmarshal(bytes, out); err != nil {
		return nil, err
	}

	return out, nil
}

func (r Reference) MarshalJSON() ([]byte, error) {
	ref := strings.Join([]string{r.Type.String(), r.ID}, "/")
	return json.Marshal(ref)
}

func (r *Reference) UnmarshalJSON(bytes []byte) error {
	type wrapper struct {
		Reference string `json:"reference"`
	}
	var ref wrapper
	if err := json.Unmarshal(bytes, &ref); err != nil {
		return err
	}

	elements := strings.Split(ref.Reference, "/")
	r.Type = ResourceType(elements[0]) // TODO check for valid type
	r.ID = elements[1]

	return nil
}

