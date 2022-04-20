package mq

import "prow/library/mq/engine/kafka"

type Producer interface {
	// producer
	Push(data []byte, topic string, blockKey ...string) error
}

func NewPusher(pusherTypes ...string) Producer {
	var producer Producer
	if len(pusherTypes) == 0 {
		pusherTypes = []string{"kafka"}
	}
	switch pusherTypes[0] {
	case "kafka":
		producer = kafka.NewKafkaPusher()
		break
	}
	return producer
}
