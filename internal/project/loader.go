package project

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mishalalajmi/mimic/internal/mock"
	"gopkg.in/yaml.v3"
)

func Load(path string) (*Project, error) {
	project, err := findProject(path)
	if err != nil {
		return nil, err
	}

	definitions := make([]mock.Definition, 0, len(project.DefinitionFiles))
	for _, f := range project.DefinitionFiles {
		def, err := mock.Load(filepath.Join(path, f))
		if err != nil {
			return nil, fmt.Errorf("loading definition: %s: %v", f, err)
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
