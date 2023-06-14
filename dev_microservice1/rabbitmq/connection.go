package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Connection represents a RabbitMQ connection.
type Connection struct {
	conn *amqp.Connection
}

// NewConnection creates a new RabbitMQ connection.
func NewConnection(amqpURL string) (*Connection, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	return &Connection{conn: conn}, nil
}

// Close closes the RabbitMQ connection.
func (c *Connection) Close() error {
	if c.conn == nil {
		return nil
	}

	return c.conn.Close()
}

// Channel creates a new RabbitMQ channel.
func (c *Connection) Channel() (*amqp.Channel, error) {
	return c.conn.Channel()
}
