package internal

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitClient struct {
	conn *amqp.Connection // it is a TCP connection used by client
	ch   *amqp.Channel    // multiplex over connection, it is used to process / send messages
}

func ConnectRabbitMQ(username, password, host, vhost string) (*amqp.Connection, error) {
	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, host, vhost))
}

func NewRabbitMQClient(conn *amqp.Connection) (RabbitClient, error) {
	ch, err := conn.Channel()
	if err != nil {
		return RabbitClient{}, err
	}

	if err := ch.Confirm(false); err != nil {
		return RabbitClient{}, err
	}

	return RabbitClient{
		conn: conn,
		ch:   ch,
	}, nil
}

func (rc RabbitClient) Close() error {
	return rc.ch.Close()
}

// // CreateQueue creates a new queue based on given cfgs
// func (rc RabbitClient) CreateQueue(queueName string, durable, autodelete bool) error {
// 	// durable = true, mantain the messages
// 	_, err := rc.ch.QueueDeclare(queueName, durable, autodelete, false, false, nil)
// 	return err
// }

func (rc RabbitClient) CreateQueue(queueName string, durable, autodelete bool) (amqp.Queue, error) {
	// durable = true, mantain the messages
	q, err := rc.ch.QueueDeclare(queueName, durable, autodelete, false, false, nil)
	if err != nil {
		return amqp.Queue{}, err
	}

	return q, nil
}

// CreateBinding binds the current channel to the given exchange using the routingkey provided
func (rc RabbitClient) CreateBinding(name, binding, exchange string) error {
	// leaving nowait false, having nowait set to false will make the channler return an if its binding fails
	return rc.ch.QueueBind(name, binding, exchange, false, nil)
}

// Send is used to public payload onto an exchange with the given routingKey
func (rc RabbitClient) Send(ctx context.Context, exchange, routingKey string, options amqp.Publishing) error {
	confirmation, err := rc.ch.PublishWithDeferredConfirmWithContext(
		ctx,
		exchange,
		routingKey,
		// Mandatory is used to determine if an error should be returned upon failure
		true,
		false,
		options,
	)

	if err != nil {
		return err
	}

	log.Println(confirmation.Wait())
	return nil
}

// func (rc RabbitClient) Send(ctx context.Context, exchange, routingKey string, options amqp.Publishing) error {
// 	return rc.ch.PublishWithContext(
// 		ctx,
// 		exchange,
// 		routingKey,
// 		// Mandatory is used to determine if an error should be returned upon failure
// 		true,
// 		false,
// 		options,
// 	)
// }

// Consume is used to consume a queue
func (rc RabbitClient) Consume(queue, consumer string, autoAck bool) (<-chan amqp.Delivery, error) {
	return rc.ch.Consume(queue, consumer, autoAck, false, false, false, nil)
}

// ApplyQos
// prefetch count - an integer on how many unacknowledged messages the server can deliver
// prefetch size - an integer of how many bytes
// global - determines if the rule should be applied globally or not
func (rc RabbitClient) ApplyQos(count, size int, global bool) error {
	return rc.ch.Qos(count, size, global)
}
