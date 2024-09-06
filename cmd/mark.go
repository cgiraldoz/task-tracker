package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewMarkCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mark",
		Short: "Mark a task on your TODO list as complete",
		Long:  "Mark a task on your TODO list as complete",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("task called")
		},
	}
}
