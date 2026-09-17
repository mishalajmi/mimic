package commands

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [name]",
	Short: "Create a new Mimic project",
	Args:  cobra.MaximumNArgs(1),
	RunE:  initProject,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initProject(cmd *cobra.Command, args []string) error {
	projectDir := projectPath

	if len(args) == 1 {
		projectDir = filepath.Join(projectPath, args[0])
	}

	if err := os.MkdirAll(projectDir, 0o755); err != nil {
		return fmt.Errorf("creating project directory: %w", err)
	}

	projectFile := filepath.Join(projectDir, "mimic.yaml")

	if _, err := os.Stat(projectFile); err == nil {
		return fmt.Errorf("project already exists at %q", projectFile)
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("checking project file: %w", err)
	}

	projectName := filepath.Base(filepath.Clean(projectDir))

	mocksDir := filepath.Join(projectDir, "mocks")

	if err := os.MkdirAll(mocksDir, 0o755); err != nil {
		return fmt.Errorf("creating mocks directory: %w", err)
	}

	projectContent := fmt.Sprintf(`$schema: https://raw.githubusercontent.com/mishalajmi/mimic/main/schema/mimic-project-v1.json
name: %s
definitions:
  - mocks/example.yaml
`, projectName)

	definitionContent := `$schema: https://raw.githubusercontent.com/mishalajmi/mimic/main/schema/mimic-v1.json
name: example
routes:
  - name: hello
    request:
      method: GET
      path: /hello
    response:
      status: 200
      headers:
       Content-Type: "application/json"
       X-Mimic: "true"
      body:
       message: hello, mimic
`

	if err := os.WriteFile(projectFile, []byte(projectContent), 0o644); err != nil {
		return fmt.Errorf("writing project file: %w", err)
	}

	definitionFile := filepath.Join(mocksDir, "example.yaml")

	if err := os.WriteFile(definitionFile, []byte(definitionContent), 0o644); err != nil {
		return fmt.Errorf("writing example definition: %w", err)
	}

	fmt.Printf("Created Mimic project in %q\n", projectDir)

	return nil
}

