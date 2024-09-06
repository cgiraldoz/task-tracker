package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "task-cli",
		Short: "A CLI for managing your TODOs.",
		Long:  "task-cli is a CLI for managing your TODOs.",
	}
)

func init() {
	rootCmd.AddCommand(NewAddCmd())
	rootCmd.AddCommand(NewUpdateCmd())
	rootCmd.AddCommand(NewDeleteCmd())
	rootCmd.AddCommand(NewMarkCmd())
	rootCmd.AddCommand(NewListCmd())
}

func Execute() error {
	return rootCmd.Execute()
}
