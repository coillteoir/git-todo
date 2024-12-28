package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an item from your todo list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
