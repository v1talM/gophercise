package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gophercise/task/db"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("啊哦...好像出了一点问题:", err.Error())
			return
		}
		fmt.Printf("已添加\"%s\"到Todo list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
