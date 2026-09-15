package commands

import "github.com/spf13/cobra"

var rootCommand = &cobra.Command{
	Use:   "mimic",
	Short: "A modern API mocking utility",
}

func Execute() error {
	return rootCommand.Execute()
}
