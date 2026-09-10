package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	addr string
}

func New(addr string) *Server {
	return &Server{
		addr,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, from mimic!")
	})

	return http.ListenAndServe(s.addr, mux)
}
