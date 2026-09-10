package mock

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*Definition, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var definition Definition

	if err := yaml.Unmarshal(data, &definition); err != nil {
		return nil, err
	}

	if err := definition.Validate(); err != nil {
		return nil, err
	}

	return &definition, nil
}
