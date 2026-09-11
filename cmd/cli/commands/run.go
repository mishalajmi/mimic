package commands

import (
	"fmt"

	"github.com/mishalalajmi/mimic/internal/project"
	"github.com/mishalalajmi/mimic/internal/server"
	"github.com/spf13/cobra"
)

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Start the mimic mock server",
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := project.Load("examples/example-project/")
		if err != nil {
			return fmt.Errorf("loading project: %w", err)
		}

		srv := server.New(":4000", p)
		fmt.Printf("Starting Mimic project %q...\n", p.Name)
		return srv.Start()
	},
}

func init() {
	rootCommand.AddCommand(runCommand)
}
