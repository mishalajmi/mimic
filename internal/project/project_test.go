package project

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	dir := t.TempDir()

	projectFile := `
name: test-project
definitions:
  - mock.yaml
`

	definitionFile := `
routes:
  - name: Hello Mock
    request:
      method: GET
      path: /hello
    response:
      status: 200
      headers:
        Content-Type: text/plain
        X-Mimic: "true"
      body: "Hello, from mimic"
`

	if err := os.WriteFile(
		filepath.Join(dir, "mimic.yaml"),
		[]byte(projectFile),
		0o644,
	); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(
		filepath.Join(dir, "mock.yaml"),
		[]byte(definitionFile),
		0o644,
	); err != nil {
		t.Fatal(err)
	}

	p, err := Load(dir)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if p.Name != "test-project" {
		t.Fatalf("project name = %q, want %q", p.Name, "test-project")
	}

	if len(p.Definition) != 1 {
		t.Fatalf("definitions = %d, want 1", len(p.Definition))
	}
}

func TestLoad_NoProject(t *testing.T) {
	dir := t.TempDir()

	_, err := Load(dir)
	if err == nil {
		t.Fatal("Load() expected an error")
	}
}

func TestLoad_InvalidDefinition(t *testing.T) {
	dir := t.TempDir()

	projectFile := `
name: test-project
definitions:
  - mock.yaml
`

	definitionFile := `
routes:
  - name: Hello Mock
    request:
      path: /hello
    response:
      status: 200
      headers:
        Content-Type: text/plain
        X-Mimic: "true"
      body: "Hello, from mimic"
`
	if err := os.WriteFile(
		filepath.Join(dir, "mimic.yaml"),
		[]byte(projectFile),
		0o644,
	); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(
		filepath.Join(dir, "mock.yaml"),
		[]byte(definitionFile),
		0o644,
	); err != nil {
		t.Fatal(err)
	}

	_, err := Load(dir)
	if err == nil {
		t.Fatal("Load() expected an error")
	}
}
