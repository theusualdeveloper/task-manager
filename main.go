package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/theusualdeveloper/task-manager/task"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		showCommandGuidelines()
		os.Exit(1)
	}
	command := args[1]
	ts, err := task.NewTaskStore("tasks.json")
	if err != nil {
		log.Fatal(err)
	}
	switch command {
	case "add":
		cs := args[2:]
		if len(cs) == 0 {
			fmt.Println("Please provide task content")
			os.Exit(1)
		}
		c := strings.Join(cs, " ")
		task, err := ts.Add(c)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Task added: [%d] %s\n", task.ID, task.Title)
	case "list":
		tasks := ts.List()
		var done string
		for _, task := range tasks {
			if task.Done {
				done = "done"
			} else {
				done = "pending"
			}
			fmt.Printf("[%d] %s    (%s)\n", task.ID, task.Title, done)
		}
	case "done":
		cs := args[2:]
		if len(cs) == 0 {
			fmt.Println("Please provide task id")
			os.Exit(1)
		}
		id, err := strconv.Atoi(cs[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = ts.Complete(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Task %d marked as done.\n", id)
	case "delete":
		cs := args[2:]
		if len(cs) == 0 {
			fmt.Println("Please provide task id")
			os.Exit(1)
		}
		id, err := strconv.Atoi(cs[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = ts.Delete(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Task %d deleted.\n", id)
	default:
		showCommandGuidelines()
	}
}

func showCommandGuidelines() {
	fmt.Println("Usage:")
	fmt.Println(" add <title>")
	fmt.Println(" list")
	fmt.Println(" done <id>")
	fmt.Println(" delete <id>")
}
