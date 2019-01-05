package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gophercise/task/db"
	"strconv"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "To mark the list you have done.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("啊哦...好像出了一点问题:", err.Error())
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("无效的id")
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("标记任务\"%d\"为完成状态失败，error:%s\n", id, err)
			} else {
				fmt.Printf("已标记任务\"%d\"为完成状态\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
