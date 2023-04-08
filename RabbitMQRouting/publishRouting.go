package main

import (
	"fmt"
	"strconv"
	RabbitMQ "test_rabbitmq/rabbitmq"
	"time"
)

func main() {
	test1 := RabbitMQ.NewRabbitMQRouting("exTest", "test_one")
	test2 := RabbitMQ.NewRabbitMQRouting("exTest", "test_two")

	for i := 0; i < 10; i++ {
		test1.PublishRouting("Hello test one!" + strconv.Itoa(i))
		test2.PublishRouting("Hello test two!" + strconv.Itoa(i))
		fmt.Println("Sending No. " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
