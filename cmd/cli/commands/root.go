package commands

import "github.com/spf13/cobra"

var projectPath string

var rootCmd = &cobra.Command{
	Use:           "mimic",
	Short:         "A modern API mocking utility",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&projectPath, "path", ".", "Path to the Mimic project")
}

func Execute() error {
	return rootCmd.Execute()
}
