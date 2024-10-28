FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o task-queue-app

CMD ["./task-queue-app"]

