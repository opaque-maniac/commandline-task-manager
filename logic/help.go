package logic

import "fmt"

func HelpMessage() {
	fmt.Println("Usage: todo <command> <item>")
	fmt.Println("Commands:")
	fmt.Println("  help          Show this help message")
	fmt.Println("  add <item>    Add an item to the list")
	fmt.Println("  list          List all items")
	fmt.Println("  remove <item> Remove an item from the list")
	fmt.Println("  remove-all    Remove all items from the list")
}
