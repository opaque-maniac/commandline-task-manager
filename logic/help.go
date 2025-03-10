package logic

import "fmt"

// HelpMessage displays available commands and their usage.
func HelpMessage() {
	fmt.Println("Todo is a CLI task manager")
	fmt.Println()
	fmt.Println("*todo* is a placeholder for the binary name you have chosen")
	fmt.Println("Usage: todo <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  help                   Show this help message")
	fmt.Println("  add <item>             Add an item to the list")
	fmt.Println("  list                   List all items")
	fmt.Println("  remove <item>          Remove an item from the list")
	fmt.Println("  remove-all             Remove all items from the list")
	fmt.Println("  update <old> <new>     Rename an item in the list")
	fmt.Println("  complete <item>        Mark an item as completed")
	fmt.Println("  un-complete <item>     Mark an item as un-completed")
	fmt.Println("  web <port>             Start the web server on http://localhost:<port>")

	fmt.Println("\nExamples:")
	fmt.Println("  todo add 'Buy milk'         # Adds 'Buy milk' to the list")
	fmt.Println("  todo list                   # Lists all tasks")
	fmt.Println("  todo remove 'Buy milk'      # Removes 'Buy milk' from the list")
	fmt.Println("  todo update 'Buy milk' 'Buy coffee'  # Renames a task")
	fmt.Println("  todo complete 'Buy coffee'  # Marks a task as completed")
	fmt.Println("  todo un-complete 'Buy coffee'  # Marks a task as un-completed")
	fmt.Println("  todo remove-all             # Removes all tasks")
	fmt.Println("  todo web 8080               # Starts the web server on http://localhost:8080")
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
	case "un-complete": // Explicit case for un-complete
		message = "Task marked as un-completed"
	case "list": // Prevent "list" from printing an unintended message
		return
	default:
		return // Default does nothing if the command is unknown
	}

	fmt.Println(message)
}
