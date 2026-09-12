package commands

import (
	"fmt"

	"github.com/mishalalajmi/mimic/internal/project"
	"github.com/mishalalajmi/mimic/internal/server"
	"github.com/spf13/cobra"
)

var (
	projectPath string
	port        int
)

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Start the mimic mock server",
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := project.Load(projectPath)
		if err != nil {
			return fmt.Errorf("loading project: %w", err)
		}

		addr := fmt.Sprintf(":%d", port)
		srv := server.New(addr, p)
		fmt.Printf("Starting Mimic project %q on port %s...\n", p.Name, addr)

		return srv.Start()
	},
}

func init() {
	rootCommand.AddCommand(runCommand)

	runCommand.Flags().StringVar(&projectPath, "path", ".", "Path to project and mock definitions")
	runCommand.Flags().IntVar(&port, "port", 4010, "Port to run the mimic mock server on")
}
