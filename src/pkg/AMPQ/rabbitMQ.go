package AMPQ

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Send2AMPQ(metricJsonChan <-chan string, send bool) {
	var strMsg string

	if send == false { // no need to go any further

		return
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	log.Printf(" Connected to RabbitMQ ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Fenway", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare Fenway queue")

	log.Printf("AMQP Sending Data to Fenway Queue ")

	for {
		strMsg = <-metricJsonChan

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(strMsg),
			})
		log.Printf("AMQP  %s", strMsg)
		failOnError(err, "Failed to publish a message")

	}

}
