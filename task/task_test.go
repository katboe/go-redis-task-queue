package task

import (
	"testing"
)

func TestPerformTask(t *testing.T) {
	// Create a sample task
	task := Task{
		Name:     "Sample Task",
		Priority: 0,
		Retries:  0,
		Delay:    1,
	}

	// Call performTask multiple times to check for success/failure
	for i := 0; i < 10; i++ {
		err := performTask(task)
		if err != nil {
			t.Logf("Task %s failed on attempt %d: %v", task.Name, i+1, err)
		} else {
			t.Logf("Task %s succeeded on attempt %d", task.Name, i+1)
		}
	}
}
