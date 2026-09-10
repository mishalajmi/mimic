package main

import (
	"fmt"
	"log"

	"github.com/mishalalajmi/mimic/internal/mock"
)

func main() {
	definition, err := mock.Load("examples/mock.yaml")
	if err != nil {
		log.Fatal(err)
	}

	for _, route := range definition.Routes {
		fmt.Printf("%s %s -> %d\n", route.Method, route.Path, route.Response.Status)
	}
}
