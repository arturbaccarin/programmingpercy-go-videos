package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/arturbaccarin/learn-rabbitmq-for-eda/internal"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := internal.ConnectRabbitMQ("percy", "secret", "localhost:5672", "customers")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// all consuming will be done on this Connection
	consumeConn, err := internal.ConnectRabbitMQ("percy", "secret", "localhost:5672", "customers")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client, err := internal.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// consumeClient
	consumeClient, err := internal.NewRabbitMQClient(consumeConn)
	if err != nil {
		panic(err)
	}
	defer consumeClient.Close()

	queue, err := consumeClient.CreateQueue("", true, true)
	if err != nil {
		panic(err)
	}

	if err := consumeClient.CreateBinding(queue.Name, queue.Name, "customer_callbacks"); err != nil {
		panic(err)
	}

	messageBus, err := consumeClient.Consume(queue.Name, "customer-api", true)
	if err != nil {
		panic(err)
	}

	go func() {
		for message := range messageBus {
			log.Printf("Message Callback %s\n", message.CorrelationId)
		}
	}()

	// if err := client.CreateQueue("customers_created", true, false); err != nil {
	// 	panic(err)
	// }

	// if err := client.CreateQueue("customers_test", false, true); err != nil {
	// 	panic(err)
	// }

	// if err := client.CreateBinding("customers_created", "customers.created.*", "customer_events"); err != nil {
	// 	panic(err)
	// }

	// if err := client.CreateBinding("customers_test", "customers.*", "customer_events"); err != nil {
	// 	panic(err)
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < 10; i++ {
		if err := client.Send(ctx, "customer_events", "customers.created.us", amqp.Publishing{
			ContentType:   "text/plain",
			DeliveryMode:  amqp.Persistent,
			ReplyTo:       queue.Name,
			CorrelationId: fmt.Sprintf("customer_created_%d", i),
			Body:          []byte(`An cool message between services`),
		}); err != nil {
			panic(err)
		}

		// sending a transient message
		// if err := client.Send(ctx, "customer_events", "customers.test", amqp.Publishing{
		// 	ContentType:  "text/plain",
		// 	DeliveryMode: amqp.Transient,
		// 	Body:         []byte(`An uncool undurable message`),
		// }); err != nil {
		// 	panic(err)
		// }
	}

	// time.Sleep(10 * time.Second)

	var blocking chan struct{}
	<-blocking
	// log.Println(client)
}
