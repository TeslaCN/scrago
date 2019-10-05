package test

import (
	"github.com/streadway/amqp"
	"log"
	"testing"
	"time"
)

func BenchmarkSend(b *testing.B) {
	channel, queue, e := connect()
	defer func() { _ = channel.Close() }()
	for i := 0; i < b.N; i++ {
		s := time.Now().String()
		log.Printf("<<< %s\n", s)
		e = channel.Publish("", queue.Name, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(s),
		})
		failOnError(e)
	}
}

func TestSend(t *testing.T) {
	channel, queue, e := connect()
	defer func() { _ = channel.Close() }()

	s := time.Now().String()
	log.Printf("<<< %s\n", s)
	e = channel.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(s),
	})
	failOnError(e)

}

func TestConsume(t *testing.T) {
	channel, queue, e := connect()
	defer func() { _ = channel.Close() }()

	deliveries, e := channel.Consume(queue.Name, "test_consumer", true, false, false, false, nil)
	failOnError(e)

	c := make(chan byte)

	go func() {
		for m := range deliveries {
			log.Printf(">>> %s\n", m.Body)
		}
	}()
	<-c
}

func connect() (*amqp.Channel, amqp.Queue, error) {
	connection, e := amqp.Dial("amqp://scrago:scrago@lo:5672/scrago")
	failOnError(e)
	channel, e := connection.Channel()
	failOnError(e)
	queue, e := channel.QueueDeclare("test_queue", false, false, false, false, nil)
	failOnError(e)
	return channel, queue, e
}

func failOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
