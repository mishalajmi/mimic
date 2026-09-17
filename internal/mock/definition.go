package mock

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var ErrInvalidDefinition = errors.New("invalid definition")

type Definition struct {
	Name   string  `yaml:"name"`
	Routes []Route `yaml:"routes"`
}

func Load(path string) (*Definition, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading definition file: %w", err)
	}

	var definition Definition

	if err := yaml.Unmarshal(data, &definition); err != nil {
		return nil, fmt.Errorf("%w: parsing definition YAML: %v", ErrInvalidDefinition, err)
	}

	if err := definition.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidDefinition, err)
	}

	return &definition, nil
}

func (d *Definition) Validate() error {
	if len(d.Routes) == 0 {
		return fmt.Errorf("at least one route is required")
	}

	for i, r := range d.Routes {
		if err := r.Validate(i); err != nil {
			return err
		}
	}

	return nil
}
