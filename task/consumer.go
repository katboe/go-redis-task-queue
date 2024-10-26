// consume tasks from redis queue
package task

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/katboe/go-redis-task-queue/config"
)

func ConsmeTask(delay int) {
	for {
		task, err := config.Rdb.RPop("task_queue").Result()
		if err == redis.Nil {
			fmt.Println("No tasks in queue, taking a short nap...")
			time.Sleep(2 * time.Second)
			continue
		} else if err != nil {
			fmt.Printf("Error retrieving task: %v\n", err)
			continue
		}
		fmt.Printf("Processing task %s\n", task)

		//Task = napping
		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Printf("Task completed: %s\n", task)
	}

}
