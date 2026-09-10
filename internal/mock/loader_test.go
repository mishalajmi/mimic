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
  - method: GET
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
