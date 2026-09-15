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
	project, err := findProject(path)
	if err != nil {
		return nil, err
	}

	definitions := make([]mock.Definition, 0, len(project.DefinitionFiles))
	for _, f := range project.DefinitionFiles {
		def, err := mock.Load(filepath.Join(path, f))
		if err != nil {
			return nil, fmt.Errorf("%s: %v", f, err)
		}

		definitions = append(definitions, *def)
	}
	project.Definition = definitions

	return project, nil
}

func findProject(path string) (*Project, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		data, err := os.ReadFile(filepath.Join(path, entry.Name()))
		if err != nil {
			return nil, err
		}

		var p Project
		if err := yaml.Unmarshal(data, &p); err != nil {
			continue // not a valid project file
		}
		if err := p.Validate(); err != nil {
			return nil, fmt.Errorf("%w, %v", ErrInvalidProject, err)
		}

		return &p, nil
	}

	return nil, ErrProjectNotFound
}
