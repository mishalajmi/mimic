package matcher

import (
	"net/http"
	"testing"

	"github.com/mishalalajmi/mimic/internal/mock"
)

func TestMatch(t *testing.T) {
	routes := []mock.Route{
		{
			Method: "GET",
			Path:   "/hello",
		},
	}

	matcher := New(routes)
	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	route := matcher.Match(req)
	if route == nil {
		t.Fatal("Expected the route to match")
	}

	if route.Path != "/hello" {
		t.Fatalf("expected /hello, got %s", route.Path)
	}
}

func TestNoMatch(t *testing.T) {
	routes := []mock.Route{
		{
			Method: "GET",
			Path:   "/hello",
		},
	}

	matcher := New(routes)
	req, err := http.NewRequest(http.MethodGet, "/goodbye", nil)
	if err != nil {
		t.Fatal(err)
	}

	route := matcher.Match(req)
	if route != nil {
		t.Fatal("Expected the route to not match")
	}
}

func TestMethodMustMatch(t *testing.T) {
	routes := []mock.Route{
		{
			Method: "GET",
			Path:   "/hello",
		},
	}

	matcher := New(routes)
	req, err := http.NewRequest(http.MethodPost, "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	route := matcher.Match(req)
	if route != nil {
		t.Fatal("Expected the route to not match")
	}
}
