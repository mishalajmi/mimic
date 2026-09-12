package project

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mishalalajmi/mimic/internal/mock"
)

var (
	ErrProjectNotFound = errors.New("project not found")
	ErrInvalidProject  = errors.New("invalid project")
)

type Project struct {
	Name            string            `yaml:"name"`
	Definition      []mock.Definition `yaml:"-"`
	DefinitionFiles []string          `yaml:"definitions"`
}

func (p *Project) Validate() error {
	if len(p.DefinitionFiles) == 0 {
		return fmt.Errorf("no mock definitions are configured")
	}

	if strings.TrimSpace(p.Name) == "" {
		return fmt.Errorf("project name is required")
	}

	return nil
}

func (p *Project) GetRoutes() []mock.Route {
	var routes []mock.Route
	for _, d := range p.Definition {
		routes = append(routes, d.Routes...)
	}
	return routes
}
