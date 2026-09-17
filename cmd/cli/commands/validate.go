package commands

import (
	"errors"
	"fmt"

	"github.com/mishalalajmi/mimic/internal/project"
	"github.com/spf13/cobra"
)

var validateCommand = &cobra.Command{
	Use:   "validate",
	Short: "Validate a Mimic project",
	RunE:  validate,
}

func validate(cmd *cobra.Command, args []string) error {
	p, err := project.Load(projectPath)
	if errors.Is(err, project.ErrProjectNotFound) {
		return fmt.Errorf(
			"project not found at %q: use --path to specify the project directory",
			projectPath,
		)
	}

	fmt.Printf("Project %q is valid\n", p.Name)
	return nil
}

func init() {
	rootCommand.AddCommand(validateCommand)
}
