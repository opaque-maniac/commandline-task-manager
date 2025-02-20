package logic

import (
	"encoding/json"
	"fmt"
	"os"
)

const filename = "todo-data.json"

func AddTask(data string) error {
	if data == "" {
		return fmt.Errorf("Task cannot be empty")
	}

	var items []Todo
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		items = []Todo{}
	} else {
		items, err = ReadData()
		if err != nil {
			return err
		}
	}

	for _, v := range items {
		if v.Task == data {
			return nil
		}
	}

	newItem := Todo{Task: data, Completed: false}
	items = append(items, newItem)
	jsonStr, err := json.Marshal(items)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonStr, 0644)
}

func ReadData() ([]Todo, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []Todo{}, nil
	}

	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var items []Todo
	err = json.Unmarshal(data, &items)

	if err != nil {
		return nil, err
	}

	return items, nil
}

func ListTasks() error {
	items, err := ReadData()

	if err != nil {
		return err
	}

	if len(items) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	for i, v := range items {
		check := " "
		if v.Completed {
			check = "X"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, check, v.Task)
	}
	return nil
}

func RemoveTask(item string) error {
	if item == "" {
		return fmt.Errorf("item cannot be empty")
	}

	items, err := ReadData()
	if err != nil {
		return err
	}

	var updatedItems []Todo
	for _, v := range items {
		if v.Task != item {
			updatedItems = append(updatedItems, v)
		}
	}

	jsonStr, err := json.Marshal(updatedItems)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonStr, 0644)
}

func RemoveAllTasks() error {
	data := []Todo{}
	jsonStr, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonStr, 0644)
}

func UpdateTask(olditem, newitem string) error {
	if newitem == "" {
		return fmt.Errorf("new item cannot be empty")
	}

	if olditem == "" {
		return fmt.Errorf("item cannot be empty")
	}

	items, err := ReadData()

	if err != nil {
		return err
	}

	for i := range items {
		if items[i].Task == olditem {
			items[i].Task = newitem
		}
	}

	jsonStr, err := json.Marshal(items)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonStr, 0644)
}

func CompleteTask(task string) error {
	if task == "" {
		return fmt.Errorf("task cannot be empty")
	}

	data, err := ReadData()

	if err != nil {
		return err
	}

	for i := range data {
		if data[i].Task == task {
			data[i].Completed = true
		}
	}

	jsonStr, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonStr, 0644)
}

func UnCompleteTask(task string) error {
	if task == "" {
		return fmt.Errorf("task cannot be empty")
	}

	data, err := ReadData()

	if err != nil {
		return err
	}

	for i := range data {
		if data[i].Task == task {
			data[i].Completed = false
		}
	}

	jsonStr, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonStr, 0644)
}
