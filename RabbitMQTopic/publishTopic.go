package main

import (
	"fmt"
	"strconv"
	RabbitMQ "test_rabbitmq/rabbitmq"
	"time"
)

func main() {
	test1 := RabbitMQ.NewRabbitMQTopic("exTestTopic", "test.topic.one")
	test2 := RabbitMQ.NewRabbitMQTopic("exTestTopic", "test.topic.two")

	for i := 0; i < 10; i++ {
		test1.PublishTopic("Hello test one!" + strconv.Itoa(i))
		test2.PublishTopic("Hello test two!" + strconv.Itoa(i))
		fmt.Println("Sending No. " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
