package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/katboe/go-redis-task-queue/config"
	"github.com/katboe/go-redis-task-queue/task"
)

func main() {
	config.InitRedis()
	fmt.Println("Redis initialized")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Set a timeout
	defer cancel()                                                           // Ensure resources are cleaned up

	numtask, err := strconv.Atoi(os.Getenv("NUM_TASKS"))
	if err != nil {
		fmt.Println("Error converting NUM_TASKS:", err)
		return
	}

	delay, err := strconv.Atoi(os.Getenv("TASK_DELAY"))
	if err != nil {
		fmt.Println("Error converting TASK_DELAY:", err)
		return
	}

	maxRetries, err := strconv.Atoi(os.Getenv("MAX_RETRIES"))
	if err != nil {
		fmt.Println("Error converting MAX_RETRIES:", err)
		return
	}

	numWorkers, err := strconv.Atoi(os.Getenv("NUM_WORKERS"))
	if err != nil {
		fmt.Println("Error converting NUM_WORKERS:", err)
		return
	}

	go func() {
		for i := 1; i <= numtask; i++ {
			task.ProduceTask(fmt.Sprintf("Nap %d", i), rand.Intn(2), delay) // Generates priorities 0 or 1; 2 retries
		}
	}()

	task.ConsumeTask(ctx, maxRetries, numWorkers)

}
