package project

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mishalalajmi/mimic/internal/mock"
	"gopkg.in/yaml.v3"
)

var canonicalProjectFileName = "mimic.yaml"
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

func Load(path string) (*Project, error) {
	projectPath := filepath.Join(path, canonicalProjectFileName)

	data, err := os.ReadFile(projectPath)
	if err != nil {
		return nil, ErrProjectNotFound
	}

	var project Project
	if err := yaml.Unmarshal(data, &project); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidProject, err)
	}

	if err := project.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidProject, err)
	}

	definitions := make([]mock.Definition, 0, len(project.DefinitionFiles))

	for _, f := range project.DefinitionFiles {
		def, err := mock.Load(filepath.Join(path, f))
		if err != nil {
			return nil, fmt.Errorf("loading definition %s: %w", f, err)
		}

		definitions = append(definitions, *def)
	}
	project.Definition = definitions

	return &project, nil
}
