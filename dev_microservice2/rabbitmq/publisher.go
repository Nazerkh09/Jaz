package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

func PublishMessage(exchange, routingKey string, message []byte) error {
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

	// Publish the message to the exchange
	err = ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish RabbitMQ message: %v", err)
	}

	return nil
}
