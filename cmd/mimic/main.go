package main

import (
	"log"

	"github.com/mishalalajmi/mimic/internal/mock"
	"github.com/mishalalajmi/mimic/internal/server"
)

func main() {
	definition, err := mock.Load("examples/mock.yaml")
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New(":4000", definition)

	log.Println("mimic listening on http://localhost:4000")

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
