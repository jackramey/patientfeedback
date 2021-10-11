package cli

import "github.com/manifoldco/promptui"

var appointmentSelectTemplates = &promptui.SelectTemplates{
	Label:    "{{ . }}:",
	Active:   "\U000027A4 {{ .Summary | cyan }} ({{ .Doctor | white }})",
	Inactive: "  {{ .Summary | cyan }} ({{ .Doctor | white }})",
	Selected: "\U000027A4 {{ .Summary | green }} {{\"with Dr\" | green}} {{ .Doctor | green }}",
	Details: `
--------- Appointment ----------
{{ "Summary:" | faint }}	{{ .Summary }}
{{ "Doctor:" | faint }}	{{ .Doctor }}`,
}

var ratingSelectTemplates = &promptui.SelectTemplates{
	Label:    "{{ . }}",
	Active:   "\U000027A4 {{ .| cyan | bold}}",
	Inactive:   "{{ .| faint }}",
	Selected: "Rating: {{ . | green | cyan }}",
}

var understoodSelectTemplates = &promptui.SelectTemplates{
	Label:    "{{ . }}",
	Active:   "\U000027A4 {{ .| cyan | bold}}",
	Inactive:   "{{ .| faint }}",
	Selected: "Understood diagnosis: {{ . | green | cyan }}",
}

var commentPromptTemplates = &promptui.PromptTemplates{
	Prompt:  "{{ . }} ",
	Valid:   "{{ . }} ",
	Success: "{{ . }} ",
}

var ratingPromptText = "Hi %s, on a scale of 1-10, would you recommend Dr %s to a friend or family member? 1 = Would not recommend, 10 = Would strongly recommend:"
var understoodPromptText = "Thank you. You were diagnosed with %s. Did Dr %s explain how to manage this diagnosis in a way you could understand?"
var commentPromptText = "We appreciate the feedback, one last question: how do you feel about being diagnosed with %s?"
