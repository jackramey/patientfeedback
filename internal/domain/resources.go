package domain

import (
	"time"
)

const (
	AppointmentResType ResourceType = "Appointment"
	BundleResType      ResourceType = "Bundle"
	DiagnosisResType   ResourceType = "Diagnosis"
	DoctorResType      ResourceType = "Doctor"
	FeedbackResType    ResourceType = "Feedback"
	PatientResType     ResourceType = "Patient"
)

type (
	Resource interface {
		GetID() string
		GetResourceType() ResourceType
	}

	ResourceType string

	ResourceHeader struct {
		ID   string       `json:"id"`
		Type ResourceType `json:"resourceType"`
	}

	Entry struct {
		Resource Resource `json:"resource"`
	}

	Bundle struct {
		ResourceHeader
		Timestamp time.Time `json:"timestamp"` // RFC3339 formatted Datetime
		Entries   []Entry   `json:"entry"`
	}

	Appointment struct {
		ResourceHeader
		Status  string            `json:"status"`
		Type    []AppointmentType `json:"type"`
		Actor   Reference         `json:"actor"`
		Subject Reference         `json:"subject"`
	}

	AppointmentType struct {
		Text string `json:"text"`
	}

	Diagnosis struct {
		ResourceHeader
		Status      string    `json:"status"`
		Appointment Reference `json:"appointment"`
		Code        Code      `json:"code"`
	}

	Code struct {
		Codings []Coding `json:"coding"`
	}

	Coding struct {
		Name   string `json:"name"`
		System string `json:"system"`
		Code   string `json:"code"`
	}

	Patient struct {
		ResourceHeader
		Active bool         `json:"active"`
		Names  []Name       `json:"name"`
		Gender   string     `json:"gender"`
		Contacts  []Contact `json:"contact"`
		Addresses []Address `json:"address"`
	}

	Doctor struct {
		ResourceHeader
		Names []Name `json:"name"`
	}

	Name struct {
		Text   string   `json:"text"`
		Family string   `json:"family"`
		Given  []string `json:"given"`
	}

	Contact struct {
		System string `json:"system"`
		Value  string `json:"value"`
		Use    string `json:"use"`
	}

	Address struct {
		Line []string `json:"line"`
		Use  string   `json:"use"`
	}

	Feedback struct {
		ResourceHeader
		Rating      int       `json:"rating"`
		Understood  bool      `json:"understood"`
		Comment     string    `json:"comment"`
		Appointment Reference `json:"appointment"`
	}

	Reference struct {
		Type ResourceType
		ID   string
	}
)

func (r ResourceType) String() string {
	return string(r)
}

func (r ResourceHeader) GetID() string {
	return r.ID
}

func (r ResourceHeader) GetResourceType() ResourceType {
	return r.Type
}

// Convenience mechanisms for getting values for script

func AppointmentSummary(a Appointment) string {
	// Assume first appointment type
	return a.Type[0].Text
}

func PatientPreferredFirstName(p Patient) string {
	// Assume first Name and Given is preferred
	if len(p.Names) == 0 || len(p.Names[0].Given) == 0 {
		return "Patient"
	}

	return p.Names[0].Given[0]
}

func PatientPreferredFullName(p Patient) string {
	// Assume first Name and Given is preferred
	if len(p.Names) == 0 || len(p.Names[0].Text) == 0 {
		return "Patient"
	}

	return p.Names[0].Text
}

func PreferredLastName(d Doctor) string {
	// Assume first Name is preferred
	if len(d.Names) == 0 {
		// Return an empty string. Anything else feels odd
		return ""
	}

	return d.Names[0].Family
}

func DiagnosisText(d Diagnosis) string {
	// TODO Right now we unsafely assume that if there is a code, we have at minimum one coding. Rather than complicating
	// this return with an error return, we're just going to return the name of the first coding in the diagnosis
	return d.Code.Codings[0].Name
}


