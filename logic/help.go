package logic

import "fmt"

// HelpMessage displays available commands and their usage.
func HelpMessage() {
	fmt.Println("Usage: todo <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  help                   Show this help message")
	fmt.Println("  add <item>             Add an item to the list")
	fmt.Println("  list                   List all items")
	fmt.Println("  remove <item>          Remove an item from the list")
	fmt.Println("  remove-all             Remove all items from the list")
	fmt.Println("  update <old> <new>     Rename an item in the list")
	fmt.Println("  complete <item>        Mark an item as completed")

	fmt.Println("\nExamples:")
	fmt.Println("  todo add 'Buy milk'         # Adds 'Buy milk' to the list")
	fmt.Println("  todo list                   # Lists all tasks")
	fmt.Println("  todo remove 'Buy milk'      # Removes 'Buy milk' from the list")
	fmt.Println("  todo update 'Buy milk' 'Buy coffee'  # Renames a task")
	fmt.Println("  todo complete 'Buy coffee'  # Marks a task as completed")
}

func SuccessMessage(command string) {
	var message string

	switch command {
	case "add":
		message = "Task added successfully"
	case "remove":
		message = "Task removed successfully"
	case "remove-all":
		message = "All tasks removed successfully"
	case "update":
		message = "Task updated successfully"
	case "complete":
		message = "Task marked as completed"
	default:
		return
	}

	fmt.Println(message)
}
