package commands

import (
	"patientfeedback/cmd/client/cli"

	"github.com/spf13/cobra"
)

var feedbackCmd = &cobra.Command{
	Use: "feedback",
	Short: "Give feedback for an appointment",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NewApp().Run()
	},
}
