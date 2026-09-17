package mock

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDefinition(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr error
	}{
		{
			name: "valid definition",
			content: `
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
			name: "invalid yaml",
			content: `
name: users
routes:
  - this is invalid
`,
			wantErr: ErrInvalidDefinition,
		},
		{
			name: "no routes",
			content: `
name: users
routes: []
`,
			wantErr: ErrInvalidDefinition,
		},
		{
			name: "missing method",
			content: `
name: users
routes:
  - request:
      path: /users
    response:
      status: 200
`,
			wantErr: ErrInvalidDefinition,
		},
		{
			name: "missing path",
			content: `
name: users
routes:
  - request:
      method: GET
    response:
      status: 200
`,
			wantErr: ErrInvalidDefinition,
		},
		{
			name: "path does not start with slash",
			content: `
name: users
routes:
  - request:
      method: GET
      path: users
    response:
      status: 200
`,
			wantErr: ErrInvalidDefinition,
		},
		{
			name: "invalid response status",
			content: `
name: users
routes:
  - request:
      method: GET
      path: /users
    response:
      status: 700
`,
			wantErr: ErrInvalidDefinition,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "definition.yaml")

			if err := os.WriteFile(path, []byte(tt.content), 0o644); err != nil {
				t.Fatal(err)
			}

			definition, err := Load(path)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if definition == nil {
				t.Fatal("expected definition, got nil")
			}

			if definition.Name != "users" {
				t.Errorf("expected name %q, got %q", "users", definition.Name)
			}

			if len(definition.Routes) != 1 {
				t.Fatalf("expected 1 route, got %d", len(definition.Routes))
			}
		})
	}
}

func TestLoadDefinition_FileNotFound(t *testing.T) {
	_, err := Load(filepath.Join(t.TempDir(), "missing.yaml"))

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected os.ErrNotExist, got %v", err)
	}
}

