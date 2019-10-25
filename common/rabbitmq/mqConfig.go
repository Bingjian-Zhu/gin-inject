package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
	"reflect"
)

/**
 * 排序执行方法
 * 交换机声明一定在前
 */
type MqConfig struct {
	Channel *amqp.Channel
}

func (m *MqConfig) AutoFunc() {
	typeOf := reflect.TypeOf(m)
	valueOf := reflect.ValueOf(m)

	numOfMethod := valueOf.NumMethod()
	for i := 0; i < numOfMethod; i++ {
		if typeOf.Method(i).Name == "AutoFunc" {
			continue
		}
		valueOf.Method(i).Call(nil)
	}
}

func (m *MqConfig) Exchange1() {
	err := m.Channel.ExchangeDeclare(
		//*
		"",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Exchange Declare [%s]: %v", "", err)
	}
	log.Println("[RabbitMQ] Exchange ->", "")
}

func (m *MqConfig) Queue1() amqp.Queue {
	queue, err := m.Channel.QueueDeclare(
		"",    //name
		false, //durable - 持久化
		false, //delete when unused - 队列清理
		false, //exclusive - 独立化
		false, //no-wait - 不堵塞
		nil,   //arguments - 其他参数
	)
	if err != nil {
		log.Fatalf("Queue Declare [%s]: %v", "", err)
	}
	err = m.Channel.QueueBind(
		queue.Name,
		"",
		"",
		false,
		nil,
	)
	log.Println("[RabbitMQ] Queue ->", "")
	return queue
}

func (m *MqConfig) Queue2() amqp.Queue {
	queue, err := m.Channel.QueueDeclare(
		"",    //name
		false, //durable - 持久化
		false, //delete when unused - 队列清理
		false, //exclusive - 独立化
		false, //no-wait - 不堵塞
		nil,   //arguments - 其他参数
	)
	if err != nil {
		log.Fatalf("Queue Declare [%s]: %v", "", err)
	}
	err = m.Channel.QueueBind(
		queue.Name,
		"",
		"",
		false,
		nil,
	)
	log.Println("[RabbitMQ] Queue ->", "")
	return queue
}
