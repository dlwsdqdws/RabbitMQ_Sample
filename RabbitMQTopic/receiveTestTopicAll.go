package main

import RabbitMQ "test_rabbitmq/rabbitmq"

func main() {
	test1 := RabbitMQ.NewRabbitMQTopic("exTestTopic", "#")
	test1.ReceiveTopic()
}
