package mock

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()

	path := filepath.Join(dir, "mock.yaml")

	content := `
routes:
  - name: test mock
    request:
      method: GET
      path: /hello
    response:
      status: 201
      headers:
        Content-Type: application/json
      body: '{"hello":"world"}'
`

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	definition, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}

	if len(definition.Routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(definition.Routes))
	}

	route := definition.Routes[0]

	if route.Response.Status != 201 {
		t.Fatalf("expected status 201, got %d", route.Response.Status)
	}

	if route.Response.Headers["Content-Type"] != "application/json" {
		t.Fatal("expected Content-Type header")
	}

	if route.Response.Body != `{"hello":"world"}` {
		t.Fatalf("unexpected body: %s", route.Response.Body)
	}
}

func TestValidate(t *testing.T) {
	definition := Definition{
		Routes: []Route{
			{
				Name: "Test Mock Definition",
				Request: Request{
					Method: "GET",
					Path:   "/hello",
				},
				Response: Response{
					Status: 200,
					Body:   "Hello",
				},
			},
		},
	}

	if err := definition.Validate(); err != nil {
		t.Fatalf("expected definition to be valid: %v", err)
	}
}

func TestValidateRequiresRoutes(t *testing.T) {
	definition := Definition{}

	if err := definition.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidateRequiresMethod(t *testing.T) {
	definition := Definition{
		Routes: []Route{
			{
				Name: "Test Mock Definition",
				Request: Request{
					Path: "/hello",
				},
				Response: Response{
					Status: 200,
				},
			},
		},
	}

	if err := definition.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidateRequiresPath(t *testing.T) {
	definition := Definition{
		Routes: []Route{
			{
				Name: "Test Mock Definition",
				Request: Request{
					Method: "GET",
				},
				Response: Response{
					Status: 200,
				},
			},
		},
	}

	if err := definition.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidatePathMustStartWithSlash(t *testing.T) {
	definition := Definition{
		Routes: []Route{
			{
				Name: "Test mock definition",
				Request: Request{
					Method: "GET",
					Path:   "hello",
				},
				Response: Response{
					Status: 200,
				},
			},
		},
	}

	if err := definition.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidateStatus(t *testing.T) {
	definition := Definition{
		Routes: []Route{
			{
				Name: "Test mock definition",
				Request: Request{
					Method: "GET",
					Path:   "/hello",
				},
				Response: Response{
					Status: 700,
				},
			},
		},
	}

	if err := definition.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}
