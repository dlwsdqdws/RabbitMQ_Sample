package main

import RabbitMQ "test_rabbitmq/rabbitmq"

func main() {
	test1 := RabbitMQ.NewRabbitMQRouting("exTest", "test_two")
	test1.ReceiveRouting()
}
