package matcher

import (
	"net/http"

	"github.com/mishalalajmi/mimic/internal/mock"
)

type Matcher struct {
	routes []mock.Route
}

func New(routes []mock.Route) *Matcher {
	return &Matcher{
		routes,
	}
}

func (m *Matcher) Match(r *http.Request) *mock.Route {
	for i := range m.routes {
		route := m.routes[i]

		if route.Path == r.URL.Path && string(route.Method) == r.Method {
			return &route
		}
	}
	return nil
}
