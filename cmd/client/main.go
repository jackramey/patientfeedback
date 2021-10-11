package main

import (
	"fmt"
	"os"

	"patientfeedback/cmd/client/commands"
)

func main() {
	if err := commands.Root().Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
	//checkErr(cli.NewApp().Run())
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
