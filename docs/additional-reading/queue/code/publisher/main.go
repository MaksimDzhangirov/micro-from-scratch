package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/streadway/amqp"
)

type addTask struct {
	Number1 int
	Number2 int
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	amqpHost := os.Getenv("AMQP_HOST")
	amqpPort := os.Getenv("AMQP_PORT")
	amqpUser := os.Getenv("AMQP_USER")
	amqpPassword := os.Getenv("AMQP_PASSWORD")
	amqpConnectionURL := fmt.Sprintf("amqp://%s:%s@%s:%s", amqpUser, amqpPassword, amqpHost, amqpPort)
	conn, err := amqp.Dial(amqpConnectionURL)
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")

	rand.Seed(time.Now().UnixNano())

	addTask := addTask{Number1: rand.Intn(999), Number2: rand.Intn(999)}
	body, err := json.Marshal(addTask)
	if err != nil {
		handleError(err, "Error encoding JSON")
	}

	err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}

	log.Printf("AddTask: %d+%d", addTask.Number1, addTask.Number2)
}
