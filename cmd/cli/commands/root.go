package commands

import "github.com/spf13/cobra"

var projectPath string

var rootCommand = &cobra.Command{
	Use:           "mimic",
	Short:         "A modern API mocking utility",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCommand.PersistentFlags().StringVar(&projectPath, "path", ".", "Path to the Mimic project")
}

func Execute() error {
	return rootCommand.Execute()
}
