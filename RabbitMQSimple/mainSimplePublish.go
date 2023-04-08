package main

import (
	"fmt"
	RabbitMQ "test_rabbitmq/rabbitmq"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple(
		"imoocSimple",
	)
	rabbitmq.PublishSimple("hello world!")
	fmt.Println("send successfully")
}
