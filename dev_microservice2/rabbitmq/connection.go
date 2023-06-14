package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection

func InitRabbitMQ() error {
	var err error

	// Connect to RabbitMQ server
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	log.Println("Connected to RabbitMQ")

	return nil
}

func CloseRabbitMQ() {
	if conn != nil {
		conn.Close()
		log.Println("Closed RabbitMQ connection")
	}
}
