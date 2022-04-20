package kafka

import (
	"prow/library/log"
	"github.com/Shopify/sarama"
)

type ConsumerKafka struct {
	client sarama.ConsumerGroup
}

func NewKafkaConsumer() *ConsumerKafka {
	return &ConsumerKafka{}
}

func (consumer *ConsumerKafka) Listen() {

	return
}

func (consumer *ConsumerKafka) getConsumerClient(addrs []string, group string) sarama.ConsumerGroup {
	// 初始化consumer client
	once.Do(func() {
		config := sarama.NewConfig()
		config.Version = sarama.V0_10_2_0
		//config.Version = sarama.V2_1_0_0
		config.Producer.Return.Successes = true
		// 启动 consumer
		client, err := sarama.NewConsumerGroup(addrs, group, config)
		if err != nil {
			log.Logger.Errorf("creating consumer client failed: %v", err)
			return
		}
		consumer.client = client
	})
	return consumer.client
}
