package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"gophercise/task/cmd"
	"gophercise/task/db"
	"os"
	"path/filepath"
)

func main() {
	homeDir, _ := homedir.Dir()
	dbPath := filepath.Join(homeDir, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}