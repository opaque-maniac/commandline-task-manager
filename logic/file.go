package logic

import (
	"fmt"
	"os"
	"strings"
)

const filename = "todo.txt"

func WriteData(data string) error {
	fmtData := fmt.Sprintf("%s\n", data)
	fp, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()

	if _, err := fp.WriteString(fmtData); err != nil {
		return err
	}

	return nil
}

func ReadData() ([]string, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []string{}, nil
	}

	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func RemoveData(item string) error {
	items, err := ReadData()
	if err != nil {
		return err
	}

	var updatedItems []string
	for _, v := range items {
		if v != item {
			updatedItems = append(updatedItems, v)
		}
	}

	return os.WriteFile(filename, []byte(strings.Join(updatedItems, "\n")+"\n"), 0644)
}

func RemoveAllData() error {
	data := "\n"
	return os.WriteFile(filename, []byte(data), 0644)
}
