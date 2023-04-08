package main

import RabbitMQ "test_rabbitmq/rabbitmq"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("" + "newProduct")
	rabbitmq.ReceivePub()
}
