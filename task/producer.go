// generate tasks & push to redis queue
package task

import (
	"fmt"

	"github.com/katboe/go-redis-task-queue/config"
)

func ProduceTask(task string) error {
	err := config.Rdb.LPush("task_queue", task).Err()
	if err != nil {
		return fmt.Errorf("error adding task: %v", err)
	}
	fmt.Printf("Task added: %s\n", task)
	return nil
}
