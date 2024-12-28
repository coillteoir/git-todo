/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"strings"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an item to your todo list",
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(".git/todo"); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("file %v does not exist, run git-todo init", "./.git/todo")
		} else if err != nil {
			return fmt.Errorf("%v",err)
		}

		if len(args) == 0 {
			return fmt.Errorf("no message given\n")
		}

		data := []byte(strings.Join(args, " ") + "\n")
		return os.WriteFile(".git/todo/01.todo", data, 0644)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
