package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	taskName string
	date     string
	priority string
)

func init() {
	rootCmd.AddCommand(addCmd)

	// Add flags to the add command
	addCmd.Flags().StringVarP(&taskName, "task", "t", "", "Task name (required)")
	addCmd.Flags().StringVarP(&date, "date", "d", "", "Due date (required)")
	addCmd.Flags().StringVarP(&priority, "priority", "p", "", "Priority level (required)")

	// Mark all flags as required
	addCmd.MarkFlagRequired("task")
	addCmd.MarkFlagRequired("date")
	addCmd.MarkFlagRequired("priority")
}

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new task",
	Long:    "Add a new task to the list with specified name, date, and priority",
	Example: "add -t 'Buy groceries' -d '2025-01-01' -p 'high'",
	Aliases: []string{"a"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Task '%s' with priority '%s' due on '%s' has been added\n", taskName, priority, date)
	},
}
