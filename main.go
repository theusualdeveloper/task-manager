package main

import (
	"fmt"
	"os"

	"github.com/theusualdeveloper/task-manager/task"
)

func main() {
	ts := task.NewTaskStore()
	ts.Add("First Task")
	ts.Add("Second Task")
	ts.Add("Third Task")

	err := ts.Complete(2)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(ts.List())
}
