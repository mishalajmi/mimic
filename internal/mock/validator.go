package mock

import (
	"fmt"
	"strings"
)

func (d *Definition) Validate() error {
	if len(d.Routes) == 0 {
		return fmt.Errorf("at least one route is required")
	}

	for i, r := range d.Routes {
		if err := r.validate(i); err != nil {
			return fmt.Errorf("route validation failed: %v", err)
		}
	}

	return nil
}

func (r *Route) validate(index int) error {
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
