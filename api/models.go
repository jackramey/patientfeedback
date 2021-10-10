package api

type Appointment struct {
	ID        string `json:"id"`
	Summary   string `json:"summary"`
	Doctor    string `json:"doctor"`
	Diagnosis string `json:"diagnosis"`
}

type Patient struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Feedback struct {
	ID         string `json:"id"`
	Rating     int    `json:"rating"`
	Understood bool   `json:"understood"`
	Comment    string `json:"comment"`
}
