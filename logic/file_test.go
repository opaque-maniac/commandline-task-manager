package logic

import "testing"

func TestWriteData(t *testing.T) {
	err := WriteData("item")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestReadData(t *testing.T) {
	_, err := ReadData()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestRemoveData(t *testing.T) {
	err := RemoveData("item")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
