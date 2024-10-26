package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/katboe/go-redis-task-queue/config"
	"github.com/katboe/go-redis-task-queue/task"
)

func main() {
	config.InitRedis()
	fmt.Println("Redis initialized")

	numtask, err := strconv.Atoi(os.Getenv("NUM_TASKS"))
	if err != nil {
		fmt.Println("Error converting MY_INTEGER:", err)
		return
	}

	delay, err := strconv.Atoi(os.Getenv("TASK_DELAY"))
	if err != nil {
		fmt.Println("Error converting MY_INTEGER:", err)
		return
	}

	go func() {
		for i := 1; i <= numtask; i++ {
			task.ProduceTask(fmt.Sprintf("Nap %d", i), rand.Intn(2)) // Generates priorities 0 or 1
		}
	}()

	task.ConsumeTask(delay)

}
