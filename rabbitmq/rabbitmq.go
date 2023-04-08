package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// amqp://account:psw@rabbitmq_address/vhost
const MQURL = "amqp://test_user:test_user@127.0.0.1:5672/test"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// list name
	QueueName string
	// exchange
	Exchange string
	// key
	Key string
	// connect
	Mqurl string
}

// create RabbitMQ struct instance
func newRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "fail to connect rabbitmq")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "fail to open a channel")
	return rabbitmq
}

func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMQ) failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

// step1: create rabbitmq in simple version
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return newRabbitMQ(queueName, "", "")
}

// step2:producer
func (r *RabbitMQ) PublishSimple(msg string) {
	// 1. apply queue, if no such queue, create
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 2. send message to queue
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}

func (r *RabbitMQ) ConsumeSimple() {
	// 1. apply queue, if no such queue, create
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 2. accept message
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 3. consume message
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("[*]Waiting for message. To exit, press CTRL+C")
	<-forever
}

func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	rabbitmq := newRabbitMQ("", exchangeName, "")
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "fail to connect rabbitmq")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "fail to open a channel")

	return rabbitmq
}

func (r *RabbitMQ) PublishPub(msg string) {
	// 1. create exchange
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Fail to declare exchange")

	// 2. send message
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}

func (r *RabbitMQ) ReceivePub() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Fail to declare an exchange")

	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Fail to declare a queue")

	err = r.channel.QueueBind(
		q.Name,
		"",
		r.Exchange,
		false,
		nil,
	)

	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message %s", d.Body)
		}
	}()

	fmt.Println("[*]Waiting for message. To exit, press CTRL+C")
	<-forever
}

func NewRabbitMQRouting(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := newRabbitMQ("", exchangeName, routingKey)
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "fail to connect rabbitmq")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "fail to open a channel")

	return rabbitmq
}

func (r *RabbitMQ) PublishRouting(msg string) {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Fail to declare exchange")

	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}

func (r *RabbitMQ) ReceiveRouting() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Fail to declare an exchange")

	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Fail to declare a queue")

	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)

	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message %s", d.Body)
		}
	}()

	fmt.Println("[*]Waiting for message. To exit, press CTRL+C")
	<-forever
}

func NewRabbitMQTopic(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := newRabbitMQ("", exchangeName, routingKey)
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "fail to connect rabbitmq")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "fail to open a channel")

	return rabbitmq
}

func (r *RabbitMQ) PublishTopic(msg string) {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Fail to declare exchange")

	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}

// * -> one word => test.* for test.hello
// # -> several words => test.# for test.hellp.one
func (r *RabbitMQ) ReceiveTopic() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Fail to declare an exchange")

	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Fail to declare a queue")

	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)

	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message %s", d.Body)
		}
	}()

	fmt.Println("[*]Waiting for message. To exit, press CTRL+C")
	<-forever
}
