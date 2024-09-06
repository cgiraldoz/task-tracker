package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a task from your TODO list",
		Long:  "Delete a task from your TODO list",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("delete called")
		},
	}
}
