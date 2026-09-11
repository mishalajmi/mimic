package server

import (
	"net/http"

	"github.com/mishalalajmi/mimic/internal/matcher"
	"github.com/mishalalajmi/mimic/internal/project"
)

type Server struct {
	addr    string
	project *project.Project
	matcher *matcher.Matcher
}

func New(addr string, p *project.Project) *Server {
	return &Server{
		addr,
		p,
		matcher.New(p.GetRoutes()),
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

	if err := WriteResponse(w, route.Response); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
