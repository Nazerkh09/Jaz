package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Publisher represents a RabbitMQ message publisher.
type Publisher struct {
	channel *amqp.Channel
}

// NewPublisher creates a new RabbitMQ message publisher.
func NewPublisher(channel *amqp.Channel) *Publisher {
	return &Publisher{channel: channel}
}

// Publish publishes a message to the specified exchange with the given routing key.
func (p *Publisher) Publish(exchange, routingKey string, message []byte) error {
	err := p.channel.Publish(
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
		return fmt.Errorf("failed to publish message: %v", err)
	}

	return nil
}
