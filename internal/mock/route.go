package mock

import (
	"fmt"
	"strings"
)

type Request struct {
	Method string `yaml:"method"`
	Path   string `yaml:"path"`
}

type Response struct {
	Status  int               `yaml:"status"`
	Headers map[string]string `yaml:"headers"`
	Body    any               `yaml:"body"`
}

type Route struct {
	Name     string   `yaml:"name"`
	Request  Request  `yaml:"request"`
	Response Response `yaml:"response"`
}

func (r *Route) Validate(index int) error {
	if strings.TrimSpace(r.Request.Method) == "" {
		return fmt.Errorf("route %d: method is not defined", index)
	}

	if strings.TrimSpace(r.Request.Path) == "" {
		return fmt.Errorf("route %d: path is required", index)
	}

	if !strings.HasPrefix(r.Request.Path, "/") {
		return fmt.Errorf("route %d: path must start with a /", index)
	}

	if r.Response.Status < 100 || r.Response.Status > 599 {
		return fmt.Errorf("route %d: response status must be between 100 and 599", index)
	}

	return nil
}
