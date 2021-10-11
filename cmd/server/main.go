package main

import (
	"patientfeedback/service"
	"patientfeedback/storage"
)

func main() {
	store := storage.MemoryStore{}
	e := service.NewEchoServer(&store)
	e.Logger.Fatal(e.Start(":1323"))
}
