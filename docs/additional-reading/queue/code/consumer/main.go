package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

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

	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)

			addTask := &addTask{}

			err := json.Unmarshal(d.Body, addTask)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			log.Printf("Result of %d + %d is : %d", addTask.Number1, addTask.Number2, addTask.Number1+addTask.Number2)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}
		}
	}()

	// ждём данных из канала, чтобы программа не завершилась
	<-stopChan
}
