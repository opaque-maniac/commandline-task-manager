package main

import (
	"commandline-taskmanager/logic"
	"fmt"
	"os"
)

func main() {
	var command string
	var item string

	if len(os.Args) == 2 {
		command = os.Args[1]
		item = ""

		if command == "help" {
			logic.HelpMessage()
			os.Exit(0)
		}
	} else if len(os.Args) > 2 {
		command = os.Args[1]
		item = os.Args[2]
	} else {
		logic.HelpMessage()
		os.Exit(1)
	}

	if err := logic.ParseArgs(command, item); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Command executed successfully")
	}
}
