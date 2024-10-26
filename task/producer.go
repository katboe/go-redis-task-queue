// generate tasks & push to redis queue
package task

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/katboe/go-redis-task-queue/config"
)

func ProduceTask(name string, priority int, delay int) error {
	task := Task{
		ID:       uuid.New().String(),
		Name:     name,
		Priority: priority,
		Retries:  0,
		Delay:    delay,
	}

	var queue string
	if priority == 1 {
		queue = "high_priority_queue"
	} else if priority == 0 {
		queue = "low_priority_queue"
	} else {
		fmt.Errorf("Priority queue %d unavailable", priority)
	}

	jsonTask, err := json.Marshal(task)
	if err != nil {
		return err
	}

	_, err = config.Rdb.LPush(queue, jsonTask).Result()
	if err != nil {
		return fmt.Errorf("error adding task: %v", err)
	}
	fmt.Printf("Priorirty %d Task added: %s\n", task.Priority, task.Name)
	return nil
}
