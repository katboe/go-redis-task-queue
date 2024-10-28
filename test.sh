#!/bin/bash

# Start Docker Compose in detached mode
docker-compose up -d

# Wait for services to be healthy (optional)
sleep 5

# Run Go tests
go test ./...

# Stop and remove Docker Compose services
docker-compose down
