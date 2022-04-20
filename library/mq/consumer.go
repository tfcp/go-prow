package mq

import "prow/library/mq/engine/kafka"

type Consumer interface {
	// @param queueName 监听的队列名
	// @param group 事件处理组，同 group 对于同一个事件只处理一次
	// @param handler 事件处理函数
	// consumer
	Listen()
}

func NewConsumer(pusherTypes ...string) Consumer {
	var consumer Consumer
	if len(pusherTypes) == 0 {
		pusherTypes = []string{"kafka"}
	}
	switch pusherTypes[0] {
	case "kafka":
		consumer = kafka.NewKafkaConsumer()
		break
	}
	return consumer
}
