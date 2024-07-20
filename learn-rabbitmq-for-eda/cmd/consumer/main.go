package main

import (
	"context"
	"log"
	"time"

	"github.com/arturbaccarin/learn-rabbitmq-for-eda/internal"
	"github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/errgroup"
)

func main() {
	conn, err := internal.ConnectRabbitMQ("percy", "secret", "localhost:5672", "customers")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	publishConn, err := internal.ConnectRabbitMQ("percy", "secret", "localhost:5672", "customers")
	if err != nil {
		panic(err)
	}
	defer publishConn.Close()

	client, err := internal.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	publishClient, err := internal.NewRabbitMQClient(publishConn)
	if err != nil {
		panic(err)
	}
	defer publishClient.Close()

	// messageBus, err := client.Consume("customers_created", "email-service", false)
	// if err != nil {
	// 	panic(err)
	// }

	queue, err := client.CreateQueue("", true, true)
	if err != nil {
		panic(err)
	}

	if err := client.CreateBinding(queue.Name, "", "customer_events"); err != nil {
		panic(err)
	}

	messageBus, err := client.Consume(queue.Name, "email-service", false)
	if err != nil {
		panic(err)
	}

	// set a timeout for 15 secs
	ctx := context.Background()

	var blocking chan struct{}

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	g, _ := errgroup.WithContext(ctx)

	// apply a hard limit on the server
	if err := client.ApplyQos(10, 0, true); err != nil {
		panic(err)
	}

	// errgroup allows us concurrent tasks
	g.SetLimit(10)

	go func() {
		for message := range messageBus {
			// Spawn a worker
			msg := message
			g.Go(func() error {
				log.Printf("New Message : %v", msg)
				time.Sleep(10 * time.Second)
				if err := msg.Ack(false); err != nil {
					log.Println("Acke message failed")
					return err
				}

				if err := publishClient.Send(ctx, "customer_callbacks", msg.ReplyTo, amqp091.Publishing{
					ContentType:   "text/plain",
					DeliveryMode:  amqp091.Persistent,
					Body:          []byte("RPC COMPLETE"),
					CorrelationId: msg.CorrelationId,
				}); err != nil {
					panic(err)
				}

				log.Printf("Acknowledged message: %v", msg.MessageId)
				return nil
			})
		}
	}()

	log.Println("Consuming, use CTRL+C to exit")
	<-blocking

	/*
		var blocking chan struct{}

		go func() {
			for message := range messageBus {
				log.Printf("New Message : %v", message)

				if !message.Redelivered {
					message.Nack(false, true)
					continue
				}

				if err := message.Ack(false); err != nil {
					log.Println("Failed to ack message")
					continue
				}

				// if err := message.Ack(false); err != nil {
				// 	log.Println("Acknowledged message: ", err)
				// 	continue
				// }
				// log.Printf("Acknowledged message: %v", message.MessageId)
			}
		}()

		log.Println("Consuming, to close the programa press CTRL+C")

		<-blocking
	*/
}
