package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/mishalalajmi/mimic/internal/matcher"
	"github.com/mishalalajmi/mimic/internal/project"
)

var ErrServerClosed = errors.New("server closed")

type Server struct {
	addr    string
	project *project.Project
	matcher *matcher.Matcher
	server  *http.Server
}

func New(addr string, p *project.Project) *Server {
	return &Server{
		addr,
		p,
		matcher.New(p.GetRoutes()),
		nil,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handleFunc)

	s.server = &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	err := s.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return ErrServerClosed
	}
	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.server == nil {
		return nil
	}

	return s.server.Shutdown(ctx)
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
