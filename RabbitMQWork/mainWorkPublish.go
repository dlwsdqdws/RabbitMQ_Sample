package main

import (
	"fmt"
	"strconv"
	RabbitMQ "test_rabbitmq/rabbitmq"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "imoocSimple")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello world!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
