package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gophercise/task/db"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists your todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("啊哦...好像出了一点问题:", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("你还没有创建过Todo项目哦，赶快创建一个吧")
			return
		}
		fmt.Println("以下是你创建的Todo List:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
