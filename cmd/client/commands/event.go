package commands

import (
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var eventCmd = &cobra.Command{
	Use: "triggerEvent [DATAFILE]",
	Short: "simulate an event hitting the backend",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(args[0])
		checkErr(err)

		req, err := http.NewRequest(http.MethodPost, "http://localhost:1323/secured/bundle", file)
		checkErr(err)

		_, err = http.DefaultClient.Do(req)
		checkErr(err)
	},
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}