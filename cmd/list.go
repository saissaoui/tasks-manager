package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all tasks",
	Long:    "List all tasks with their details",
	Example: "list",
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing all tasks")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
