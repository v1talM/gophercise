package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "add a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		var task = strings.Join(args, " ")
		fmt.Printf("you add \"%s\" to the list.", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
