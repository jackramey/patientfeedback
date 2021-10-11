package cli

import (
	"fmt"
	"strconv"

	"patientfeedback/api"
	"patientfeedback/client"

	"github.com/manifoldco/promptui"
)

type App struct {
	client client.PatientClient

	patient     api.Patient
	appointment api.Appointment
	feedback    api.Feedback
}

func NewApp() App {
	return App{
		client: client.NewPatientClient(),
	}
}

func (a App) Run() error {
	a.patient = a.selectPatient()
	a.appointment = a.selectAppointment()
	a.feedback = a.promptForFeedback()
	a.showFeedback()
	return a.client.WriteFeedback(a.appointment.ID, a.feedback)
}

func (a App) selectPatient() api.Patient {
	patients, err := a.client.GetAllPatients()
	checkErr(err)

	prompt := promptui.Select{
		Label:     "Login as:",
		Items:     patients,
		Templates: patientSelectTemplates,
	}

	i, _, err := prompt.Run()
	checkErr(err)

	return patients[i]
}

func (a App) selectAppointment() api.Appointment {
	appointments, err := a.client.GetAppointmentsForPatient(a.patient.ID)
	checkErr(err)

	prompt := promptui.Select{
		Label:     "Select appointment to leave feedback for",
		Items:     appointments,
		Templates: appointmentSelectTemplates,
	}

	i, _, err := prompt.Run()
	checkErr(err)

	return appointments[i]
}

func (a App) promptForFeedback() api.Feedback {
	return api.Feedback{
		Rating:     a.getRatingInput(),
		Understood: a.getUnderstoodInput(),
		Comment:    a.getCommentInput(),
	}
}

func (a App) getRatingInput() int {
	searcher := func(input string, index int) bool {
		val, err := strconv.Atoi(input)
		if err != nil {
			return false
		}
		return val == index
	}

	prompt := promptui.Select{
		Label:     fmt.Sprintf(ratingPromptText, a.patient.Name, a.appointment.Doctor),
		Items:     makeRange(0, 10),
		CursorPos: 5,
		Size:      11,
		Templates: ratingSelectTemplates,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	checkErr(err)
	return i
}

func (a App) getUnderstoodInput() bool {
	prompt := promptui.Select{
		Label:     fmt.Sprintf(understoodPromptText, a.appointment.Diagnosis, a.appointment.Doctor),
		Items:     []string{"Yes", "No"},
		Templates: understoodSelectTemplates,
	}

	_, result, err := prompt.Run()
	checkErr(err)

	return result == "Yes"
}

func (a App) getCommentInput() string {
	prompt := promptui.Prompt{
		Label:     fmt.Sprintf(commentPromptText, a.appointment.Diagnosis),
		Templates: commentPromptTemplates,
	}

	result, err := prompt.Run()
	checkErr(err)

	return result
}

func (a App) showFeedback() {
	fmt.Println("Thanks again! Here’s what we heard:")
	fmt.Printf("Clinician rating: %d\n", a.feedback.Rating)
	fmt.Printf("Understood diagnosis: %s\n", boolToString(a.feedback.Understood))
	fmt.Printf("Your feelings: %s\n", a.feedback.Comment)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func boolToString(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}
