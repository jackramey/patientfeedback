package main

import (
	"os"

	"patientfeedback/service"
	"patientfeedback/storage"
)

func main() {
	store := storage.MemoryStore{}
	fileReader, err := os.Open("data/patient-feedback-raw-data.json")
	if err != nil {
		panic(err)
	}

	if err := storage.LoadData(fileReader, &store); err != nil {
		panic(err)
	}
	e := service.NewEchoServer(&store)
	e.Logger.Fatal(e.Start(":1323"))
}
