package qutils

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

//check the return value for each amqp call
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel")

	return conn, ch
}

func getQueue(name string, ch *amqp.Channel) *amqp.Queue {
	q, err := ch.QueueDeclare(
		name,
		false, //durable,
		false, //autoDelete,
		false, //exclusive,
		false, //noWait,
		nil)

	failOnError(err, "Failed to declare queue")

	return &q
}
