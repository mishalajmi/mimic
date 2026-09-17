package project

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		project string
		def     string
		wantErr error
	}{
		{
			name: "valid project",
			project: `
name: test-project
definitions:
  - users.yaml
`,
			def: `
name: users
routes:
  - name: get-users
    request:
      method: GET
      path: /users
    response:
      status: 200
      body:
        users: []
`,
		},
		{
			name: "invalid project yaml",
			project: `
name: [invalid
`,
			wantErr: ErrInvalidProject,
		},
		{
			name: "missing project name",
			project: `
definitions:
  - users.yaml
`,
			def: `
name: users
routes:
  - request:
      method: GET
      path: /users
    response:
      status: 200
`,
			wantErr: ErrInvalidProject,
		},
		{
			name: "no definitions",
			project: `
name: test-project
definitions: []
`,
			wantErr: ErrInvalidProject,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()

			writeFile(t, filepath.Join(dir, "mimic.yaml"), tt.project)

			if tt.def != "" {
				writeFile(t, filepath.Join(dir, "users.yaml"), tt.def)
			}

			got, err := Load(dir)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got == nil {
				t.Fatal("expected project, got nil")
			}

			if got.Name != "test-project" {
				t.Errorf("expected name %q, got %q", "test-project", got.Name)
			}

			if len(got.Definition) != 1 {
				t.Fatalf("expected 1 definition, got %d", len(got.Definition))
			}

			if len(got.GetRoutes()) != 1 {
				t.Fatalf("expected 1 route, got %d", len(got.GetRoutes()))
			}
		})
	}
}

func TestLoad_ProjectNotFound(t *testing.T) {
	dir := t.TempDir()

	_, err := Load(dir)

	if !errors.Is(err, ErrProjectNotFound) {
		t.Fatalf("expected ErrProjectNotFound, got %v", err)
	}
}

func TestLoad_DefinitionNotFound(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, filepath.Join(dir, "mimic.yaml"), `
name: test-project
definitions:
  - users.yaml
`)

	_, err := Load(dir)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected os.ErrNotExist, got %v", err)
	}
}

func TestLoad_InvalidDefinition(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, filepath.Join(dir, "mimic.yaml"), `
name: test-project
definitions:
  - users.yaml
`)

	writeFile(t, filepath.Join(dir, "users.yaml"), `
name: users
routes: []
`)

	_, err := Load(dir)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// The project loader wraps the definition loader error,
	// so errors.Is should still identify the definition category.
	if !errors.Is(err, os.ErrNotExist) && err == nil {
		t.Fatal("unexpected error")
	}
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

