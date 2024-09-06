package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update a task on your TODO list",
		Long:  "Update a task on your TODO list",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("update called")
		},
	}
}
