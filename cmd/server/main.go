package main

import (
	"os"

	"patientfeedback/http"
	"patientfeedback/storage"
)

func main() {
	memStore := storage.MemoryStore{}
	fileReader, err := os.Open("data/patient-feedback-raw-data.json")
	if err != nil {
		panic(err)
	}

	if err := storage.LoadData(fileReader, &memStore); err != nil {
		panic(err)
	}
	e := http.NewEchoServer(&memStore)
	e.Logger.Fatal(e.Start(":1323"))
}
