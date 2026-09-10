package main

import (
	"log"

	"github.com/mishalalajmi/mimic/internal/server"
)

func main() {
	srv := server.New(":4000")

	log.Println("Start server on port 4000")

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
