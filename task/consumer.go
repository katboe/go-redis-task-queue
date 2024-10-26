// consume tasks from redis queue
package task

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/katboe/go-redis-task-queue/config"
)

func processTask(task Task, delay int) {
	//Task = napping
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Printf("Priorirty %d Task completed: %s\n", task.Priority, task.Name)
}

func ConsumeTask(delay int) {
	for {
		// First, try to pop a task from the high-priority queue
		taskJSON, err := config.Rdb.RPop("high_priority_queue").Result()
		if err != nil && err != redis.Nil {
			log.Printf("Error retrieving high-priority task: %v", err)
			continue
		}

		// If no high-priority task, check the low-priority queue
		if taskJSON == "" {
			taskJSON, err = config.Rdb.RPop("low_priority_queue").Result()
			if err != nil && err != redis.Nil {
				log.Printf("Error retrieving low-priority task: %v", err)
				continue
			} else if err == redis.Nil {
				fmt.Println("No tasks in low priority queue, taking a short nap...")
				time.Sleep(2 * time.Second)
				continue
			}
		}

		if taskJSON != "" {
			var task Task
			if err := json.Unmarshal([]byte(taskJSON), &task); err != nil {
				log.Printf("Error unmarshalling task: %v", err)
				continue
			}
			// Process the task
			processTask(task, delay)
		}
	}

}
