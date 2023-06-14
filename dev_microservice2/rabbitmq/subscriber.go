package rabbitmq

import (
	"fmt"
	"log"
)

func ConsumeMessages(exchange, queueName, routingKey string, handleMessage func([]byte) error) error {
	// Create a new RabbitMQ channel
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a RabbitMQ channel: %v", err)
	}
	defer ch.Close()

	// Declare the exchange
	err = ch.ExchangeDeclare(
		exchange, // exchange name
		"direct", // exchange type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare RabbitMQ exchange: %v", err)
	}

	// Declare the queue
	q, err := ch.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // auto-deleted
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare RabbitMQ queue: %v", err)
	}

	// Bind the queue to the exchange
	err = ch.QueueBind(
		q.Name,     // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to bind RabbitMQ queue to exchange: %v", err)
	}

	// Start consuming messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to consume RabbitMQ messages: %v", err)
	}

	// Handle incoming messages
	go func() {
		for msg := range msgs {
			err := handleMessage(msg.Body)
			if err != nil {
				log.Printf("failed to handle RabbitMQ message: %v", err)
			}
		}
	}()

	return nil
}
