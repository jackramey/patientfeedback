package api

type Appointment struct {
	ID        string `json:"id"`
	Summary   string `json:"summary"`
	Doctor    string `json:"doctor"`
	Diagnosis string `json:"diagnosis"`
}

type Patient struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
}

type Feedback struct {
	ID         string `json:"id"`
	Rating     int    `json:"rating"`
	Understood bool   `json:"understood"`
	Comment    string `json:"comment"`
}

type GetAppointmentsResponse struct {
	Appointments []Appointment `json:"appointments"`
}

type GetPatientsResponse struct {
	Patients []Patient `json:"patients"`
}

type CreateFeedbackRequest struct {
	Rating     int    `json:"rating"`
	Understood bool   `json:"understood"`
	Comment    string `json:"comment"`
}

type GetFeedbackResponse struct {
	Feedback []Feedback `json:"feedback"`
}
