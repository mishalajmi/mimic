package server

import (
	"fmt"
	"net/http"

	"github.com/mishalalajmi/mimic/internal/mock"
)

type Server struct {
	addr       string
	definition *mock.Definition
}

func New(addr string, definition *mock.Definition) *Server {
	return &Server{
		addr,
		definition,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handleFunc)

	return http.ListenAndServe(s.addr, mux)
}

func (s *Server) handleFunc(w http.ResponseWriter, r *http.Request) {
	for _, route := range s.definition.Routes {
		if string(route.Method) == r.Method && route.Path == r.URL.Path {
			w.WriteHeader(route.Response.Status)
			fmt.Fprint(w, route.Response.Body)
			return
		}
	}
	http.NotFound(w, r)
}
