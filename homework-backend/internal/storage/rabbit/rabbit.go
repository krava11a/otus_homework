package rabbit

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RQueue struct {
	conn *amqp.Connection
}

func failOnError(err error, msg string) error {
	if err != nil {
		return fmt.Errorf("Rabbit error -  %s: %s", msg, err)
	}
	return nil
}

func New(connectionString string) (*RQueue, error) {
	conn, err := amqp.Dial(connectionString)
	failOnError(err, "Failed to connect to RabbitMQ")
	return &RQueue{conn: conn}, err
	// defer conn.Close()
}

func (rq *RQueue) PublishTo(name, message string) error {

	ch, err := rq.conn.Channel()
	if err != nil {
		return failOnError(err, "Failed to open channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		name,
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		return failOnError(err, "Failed to declare a queue")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// body := "Hello World"
	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] sent %s\n", message)

	return nil
}

func (rq *RQueue) ReadFrom(name string) (mesages []string) {
	ch, err := rq.conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		true,  // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			mesages = append(mesages, fmt.Sprintf("%s", d.Body))
		}
	}()

	return mesages
}
