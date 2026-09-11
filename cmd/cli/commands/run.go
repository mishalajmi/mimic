package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Start the mimic mock server",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Starting mimic...")
		return nil
	},
}

func init() {
	rootCommand.AddCommand(runCommand)
}
