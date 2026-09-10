package server

import (
	"fmt"
	"net/http"

	"github.com/mishalalajmi/mimic/internal/matcher"
	"github.com/mishalalajmi/mimic/internal/mock"
)

type Server struct {
	addr    string
	matcher *matcher.Matcher
}

func New(addr string, definition *mock.Definition) *Server {
	return &Server{
		addr,
		matcher.New(definition.Routes),
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handleFunc)

	return http.ListenAndServe(s.addr, mux)
}

func (s *Server) handleFunc(w http.ResponseWriter, r *http.Request) {
	route := s.matcher.Match(r)

	if route == nil {
		http.NotFound(w, r)
		return
	}

	for k, v := range route.Response.Headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(route.Response.Status)
	fmt.Fprint(w, route.Response.Body)
}
