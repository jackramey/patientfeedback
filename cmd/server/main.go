package main

import (
	"patientfeedback/service"
	"patientfeedback/storage"
)

func main() {
	server := service.NewServer(service.Config{
		Store: &storage.MemoryStore{},
		Port: 1323,
	})
	server.Logger.Fatal(server.Run())
}
