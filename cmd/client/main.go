package main

import (
	"fmt"
	"os"

	"patientfeedback/cmd/client/cli"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please supply the member ID of the patient you wish to run the CLI as")
	}
	app, err := cli.NewApp(os.Args[1])
	if err != nil {
		panic(err)
	}
	app.Run()
}
