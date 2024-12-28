package cmd

import (
	"fmt"
	"path/filepath"
	"io/fs"
	"os"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the current items in your todo list",
	Run: func(cmd *cobra.Command, args []string) {
		filepath.Walk(".git/todo/", func (path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			fmt.Println(path, string(data))
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
