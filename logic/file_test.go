package logic

import "testing"

func TestValidWriteData(t *testing.T) {
	err := AddTask("item")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestInvalidWriteData(t *testing.T) {
	err := AddTask("")
	if err == nil {
		t.Errorf("Error: Task cannot be empty")
	}
}

func TestReadData(t *testing.T) {
	_, err := ReadData()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestListTasks(t *testing.T) {
	err := ListTasks()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestValidUpdateTask(t *testing.T) {
	err := UpdateTask("item", "new item")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestInvalidUpdateTask(t *testing.T) {
	err := UpdateTask("item", "")
	if err == nil {
		t.Errorf("Error: new item cannot be empty")
	}

	err = UpdateTask("", "new item")
	if err == nil {
		t.Errorf("Error: item cannot be empty")
	}
}

func TestRemoveData(t *testing.T) {
	err := RemoveTask("item")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestInvalidRemoveData(t *testing.T) {
	err := RemoveTask("")
	if err == nil {
		t.Errorf("Error: %v", err)
	}
}

func TestRemoveAllData(t *testing.T) {
	err := RemoveAllTasks()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestValidCompleteTask(t *testing.T) {
	err := CompleteTask("new item")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestInvalidCompleteTask(t *testing.T) {
	err := CompleteTask("")
	if err == nil {
		t.Errorf("Error: item cannot be empty")
	}
}
