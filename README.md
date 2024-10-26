# Go Redis Task Queue

Simple priority task queue with retrying using Redis, designed to run in a Dockerized environment.

## Prerequisites

- Go 1.20+
- Docker and Docker Compose

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/go-redis-task-queue.git
   cd go-redis-task-queue
   ```

2. Set up environment variables in a `.env` file:

    ```plaintext
    REDIS_HOST=redis
    REDIS_PORT=6379
    TASK_DELAY=3
    NUM_TASKS=10
    MAX_RETRIES=2
    ```

## Running the Application

Use Docker Compose to build and run the application:

   ```bash
    docker-compose up --build
   ```

Modify `TASK_DELAY` and `NUM_TASKS` in the `.env` file to adjust the processing delay and number of tasks. The application will enqueue and process tasks based on this input.

## Troubleshooting

- **Connection Errors**: Ensure Redis is running if using Docker Compose and that environment variables are correctly set.
- **Invalid input varaibles**: Check the value in the `.env` file.  `TASK_DELAY` and `NUM_TASKS` require numerical values.

