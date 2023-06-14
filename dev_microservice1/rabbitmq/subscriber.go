package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Subscriber represents a RabbitMQ message subscriber.
type Subscriber struct {
	channel  *amqp.Channel
	queue    *amqp.Queue
	messages <-chan amqp.Delivery
}

// NewSubscriber creates a new RabbitMQ message subscriber.
func NewSubscriber(channel *amqp.Channel, exchange, routingKey string) (*Subscriber, error) {
	queue, err := channel.QueueDeclare(
		"",    // queue name
		false, // durable
		true,  // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %v", err)
	}

	err = channel.QueueBind(
		queue.Name, // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind queue: %v", err)
	}

	messages, err := channel.Consume(
		queue.Name, // queue name
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume messages: %v", err)
	}

	return &Subscriber{
		channel:  channel,
		queue:    &queue,
		messages: messages,
	}, nil
}

// ConsumeMessages starts consuming messages and calls the provided callback function for each message received.
func (s *Subscriber) ConsumeMessages(callback func([]byte)) {
	for message := range s.messages {
		callback(message.Body)
	}
}

// Close closes the subscriber.
func (s *Subscriber) Close() error {
	if s.channel != nil {
		if err := s.channel.Cancel("", false); err != nil {
			return fmt.Errorf("failed to cancel consumer: %v", err)
		}
	}

	return nil
}
