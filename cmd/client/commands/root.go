package commands

import "github.com/spf13/cobra"

func Root() *cobra.Command {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(feedbackCmd)
	rootCmd.AddCommand(eventCmd)
	return rootCmd
}
