package main

import (
	"commandline-taskmanager/logic"
	"commandline-taskmanager/web"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		logic.HelpMessage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "help":
		logic.HelpMessage()
	case "list":
		if err := logic.ListTasks(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: add <taskname>")
			os.Exit(1)
		}
		task := os.Args[2]
		if err := logic.AddTask(task); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: remove <taskname>")
			os.Exit(1)
		}
		task := os.Args[2]
		if err := logic.RemoveTask(task); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "remove-all":
		if err := logic.RemoveAllTasks(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: update <oldTask> <newTask>")
			os.Exit(1)
		}
		oldTask, newTask := os.Args[2], os.Args[3]
		if err := logic.UpdateTask(oldTask, newTask); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: complete <taskname>")
			os.Exit(1)
		}
		task := os.Args[2]
		if err := logic.CompleteTask(task); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "un-complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: un-complete <taskname>")
			os.Exit(1)
		}
		task := os.Args[2]
		if err := logic.UnCompleteTask(task); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "web":
		if len(os.Args) != 2 {
			fmt.Println("Usage: web")
			os.Exit(1)
		}
		web.Start()
	default:
		fmt.Println("Unknown command:", command)
		logic.HelpMessage()
		os.Exit(1)
	}

	logic.SuccessMessage(command)
}
