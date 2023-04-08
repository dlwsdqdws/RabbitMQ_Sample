package main

import RabbitMQ "test_rabbitmq/rabbitmq"

func main() {
	test1 := RabbitMQ.NewRabbitMQRouting("exTest", "test_one")
	test1.ReceiveRouting()
}
