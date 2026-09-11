package matcher

import (
	"net/http"
	"testing"

	"github.com/mishalalajmi/mimic/internal/mock"
)

func TestMatch(t *testing.T) {
	routes := []mock.Route{
		{
			Name: "test mock",
			Request: mock.Request{
				Method: "GET",
				Path:   "/hello",
			},
			Response: mock.Response{
				Status:  200,
				Headers: nil,
				Body:    nil,
			},
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

	if route.Request.Path != "/hello" {
		t.Fatalf("expected /hello, got %s", route.Request.Path)
	}
}

func TestNoMatch(t *testing.T) {
	routes := []mock.Route{
		{
			Name: "test mock",
			Request: mock.Request{
				Method: "GET",
				Path:   "/hello",
			},
			Response: mock.Response{
				Status:  200,
				Headers: nil,
				Body:    nil,
			},
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
			Name: "test mock",
			Request: mock.Request{
				Method: "GET",
				Path:   "/hello",
			},
			Response: mock.Response{
				Status:  200,
				Headers: nil,
				Body:    nil,
			},
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
