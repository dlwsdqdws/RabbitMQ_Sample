package main

import (
	"fmt"
	"strconv"
	RabbitMQ "test_rabbitmq/rabbitmq"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("" + "newProduct")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishPub("PubSub No." + strconv.Itoa(i))
		fmt.Println("PubSub No." + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
