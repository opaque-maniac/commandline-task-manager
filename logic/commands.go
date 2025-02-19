package logic

import "fmt"

func ParseArgs(command, item string) error {
	var err error

	switch command {
	case "add":
		if item == "" {
			err = fmt.Errorf("Item cannot be empty")
		} else {
			err = WriteData(item)
		}
		// Add item
	case "list":
		data, err_ := ReadData()
		if err_ != nil {
			err = err_
		} else {
			fmt.Println("Todo List:")
			for k, v := range data {
				fmt.Printf("%d. %s\n", k+1, v)
			}
		}
		// List items
	case "remove":
		if item == "" {
			err = fmt.Errorf("Item cannot be empty")
		} else {
			err_ := RemoveData(item)
			if err_ != nil {
				err = err_
			}
		}
		// Remove item
	case "remove-all":
		err_ := RemoveAllData()
		if err_ != nil {
			err = err_
		}
	default:
		err = fmt.Errorf("Unknown command: %s", command)
	}

	return err
}
