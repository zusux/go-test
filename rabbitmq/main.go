package main

import (
	"log"
	"github.com/streadway/amqp"
)

func main()  {
	HelloWorldPublish()
	//HelloWorldConsume()
}

func HelloWorldPublish()  {
	conn,err := GetConnect()
	failOnError(err, "Failed to connect to RabbitMQ")
	ch,err := CreateChannel(conn)
	failOnError(err, "Failed to create channel")

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World! 22"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

}

func HelloWorldConsume()  {
	conn,err := GetConnect()
	failOnError(err, "Failed to connect to RabbitMQ")
	ch,err := CreateChannel(conn)
	failOnError(err, "Failed to create channel")

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
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

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

func CreateChannel(conn *amqp.Connection)  (*amqp.Channel,error){
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch,err
}

func GetConnect() (*amqp.Connection ,error){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	return conn,err
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}