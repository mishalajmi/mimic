package project

import (
	"fmt"
	"strings"

	"github.com/mishalalajmi/mimic/internal/mock"
)

type Project struct {
	Name            string            `yaml:"name"`
	Definition      []mock.Definition `yaml:"-"`
	DefinitionFiles []string          `yaml:"definitions"`
}

func (p *Project) Validate() error {
	if len(p.DefinitionFiles) == 0 {
		return fmt.Errorf("Project file does not contain mock defintions")
	}

	if strings.TrimSpace(p.Name) == "" {
		return fmt.Errorf("Project file must have a name")
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
