package cmd

import (
	"github.com/cgiraldoz/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a new task to your TODO list",
		Long:  "Add a new task to your TODO list with the specified description.",
		Run: func(cmd *cobra.Command, args []string) {
			task.AddTask(args)
		},
	}
}
