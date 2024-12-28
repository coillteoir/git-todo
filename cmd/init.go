package cmd

import (
	//"errors"
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the todo list in your repositiory",
	RunE: func(cmd *cobra.Command, args []string) error {
		statInfo, err := os.Stat("./.git/todo")			
		if statInfo == nil && err != nil {
			err = os.Mkdir("./.git/todo", 0750)
			if err != nil {
				return err
			}
		}	
		fmt.Println(statInfo)

		return err
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
