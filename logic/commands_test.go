package logic

import (
	"testing"
)

func TestValidAddCommand(t *testing.T) {
	err := ParseArgs("add", "item")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestInvalidAddCommand(t *testing.T) {
	err := ParseArgs("add", "")
	if err == nil {
		t.Errorf("Error: %v", err)
	}
}

func TestValidListCommand(t *testing.T) {
	err := ParseArgs("list", "")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestInvalidListCommand(t *testing.T) {
	err := ParseArgs("list", "item")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestValidRemoveCommand(t *testing.T) {
	err := ParseArgs("remove", "item")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestInvalidRemoveCommand(t *testing.T) {
	err := ParseArgs("remove", "")
	if err == nil {
		t.Errorf("Error: %v", err)
	}
}

func TestValidRemoveAllCommand(t *testing.T) {
	err := ParseArgs("remove-all", "")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestInvalidRemoveAllCommand(t *testing.T) {
	err := ParseArgs("remove-all", "item")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
