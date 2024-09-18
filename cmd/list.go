package cmd

import (
	"github.com/cgiraldoz/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all of your incomplete tasks",
		Long:  "List all of your incomplete tasks",
		Run: func(cmd *cobra.Command, args []string) {
			task.ListTasks()
		},
	}
}
