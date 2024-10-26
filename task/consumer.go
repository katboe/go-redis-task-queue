// consume tasks from redis queue
package task

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
	"github.com/katboe/go-redis-task-queue/config"
)

func processTask(task Task, maxRetries int) {
	err := performTask(task)
	if err != nil {
		log.Printf("Task %s failed: %v", task.Name, err)

		task.Retries++

		if task.Retries <= maxRetries {
			// Re-enqueue the task back to the original queue
			jsonTask, _ := json.Marshal(task)
			var queue string
			if task.Priority == 1 {
				queue = "high_priority_queue"
			} else if task.Priority == 0 {
				queue = "low_priority_queue"
			}

			_, err := config.Rdb.LPush(queue, jsonTask).Result()
			if err != nil {
				log.Printf("Error re-enqueuing task %s: %v", task.Name, err)
			}
		} else {
			// Move to failed queue
			jsonTask, _ := json.Marshal(task)
			_, err := config.Rdb.LPush("failed_queue", jsonTask).Result()
			if err != nil {
				log.Printf("Error moving task %s to failed queue: %v", task.Name, err)
			} else {
				log.Printf("Too many retires: moving task %s to failed queue", task.Name)
			}
		}
	} else {
		log.Printf("Task %s completed successfully", task.Name)
	}
}

func performTask(task Task) error {
	//Task = napping

	// Simulated task logic; replace with actual processing logic
	if rand.Float32() < 0.3 { // Simulate failure 50% of the time
		return fmt.Errorf("simulated failure")
	} else {
		time.Sleep(time.Duration(task.Delay) * time.Second)
		fmt.Printf("Priorirty %d Task completed: %s\n", task.Priority, task.Name)
		return nil
	}
}

func ConsumeTask(maxRetries int) {
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
			processTask(task, maxRetries)
		}
	}

}
