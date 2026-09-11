package main

import (
	"log"

	"github.com/mishalalajmi/mimic/internal/project"
	"github.com/mishalalajmi/mimic/internal/server"
)

func main() {
	loadedProject, err := project.Load("./examples/example-project")
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New(":4000", loadedProject)

	log.Println("mimic listening on http://localhost:4000")

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
