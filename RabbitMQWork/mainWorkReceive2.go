package main

import RabbitMQ "test_rabbitmq/rabbitmq"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "imoocSimple")
	rabbitmq.ConsumeSimple()
}
