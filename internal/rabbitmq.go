package internal

import (
	"encoding/json"
	"mds/internal/message"

	"github.com/streadway/amqp"
)

type RabbitMQ interface {
	CreateQueue(queueName string) error
	DeleteQueue(queueName string) error
	Publish(queueName string, message message.Message) error
	Consume(queueName string) (string, error)
}

type rabbitmq struct {
	conn *amqp.Connection
}

func NewRabbitMQ(conn *amqp.Connection) RabbitMQ {
	return rabbitmq{
		conn: conn,
	}
}

func (rabbit rabbitmq) CreateQueue(userId string) error {
	ch, err := rabbit.conn.Channel()
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(userId, false, false, false, false, nil)
	if err != nil {
		return err
	}

	return nil
}

func (rabbit rabbitmq) DeleteQueue(userId string) error {
	ch, err := rabbit.conn.Channel()
	if err != nil {
		return err
	}

	_, err = ch.QueueDelete(userId, false, false, false)
	if err != nil {
		return err
	}

	return nil
}

func (rabbit rabbitmq) Publish(queueName string, message message.Message) error {
	ch, err := rabbit.conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	queue, err := ch.QueueInspect(queueName)
	if err != nil {
		return err
	}

	jsonString, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(jsonString),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (rabbit rabbitmq) Consume(queueName string) (string, error) {

	ch, err := rabbit.conn.Channel()
	if err != nil {
		return "", err
	}

	defer ch.Close()

	queue, err := ch.QueueInspect(queueName)
	if err != nil {
		return "", err
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return "", err
	}

	data := <-msgs

	return string(data.Body), nil
}
