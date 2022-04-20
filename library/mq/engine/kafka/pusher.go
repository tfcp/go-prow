package kafka

import (
	"context"
	"errors"
	"prow/library/log"
	"github.com/Shopify/sarama"
	"github.com/gogf/gf/frame/g"
	"sync"
	"time"
)

var (
	once sync.Once
)

type PusherKafka struct {
	client sarama.SyncProducer
}

func NewKafkaPusher() *PusherKafka {
	return &PusherKafka{}
}

func (pusher *PusherKafka) Push(data []byte, topic string, blockKey ...string) error {
	var (
		err             error
		defaultBlockKey = ""
	)
	producer := pusher.getProducerClient(g.Config().GetStrings("mq.kafka.addr"))
	defer producer.Close()
	if len(blockKey) > 0 {
		defaultBlockKey = blockKey[0]
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(defaultBlockKey),
		Value: sarama.ByteEncoder(data),
	}

	interval := 200 * time.Millisecond
	tick := time.NewTicker(interval)
	defer tick.Stop()

	toCtx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

ForLoop:
	for {
		select {
		case <-tick.C:
			err = producer.SendMessages([]*sarama.ProducerMessage{msg})
			if err == nil {
				log.Logger.Info("send message success")

				break ForLoop
			}
		case <-toCtx.Done():
			err = errors.New("send message timeout")
			break ForLoop
		}
	}

	return err
}

func (pusher *PusherKafka) getProducerClient(addrs []string) sarama.SyncProducer {
	// 初始化push client
	once.Do(func() {
		config := sarama.NewConfig()
		config.Version = sarama.V0_10_2_0
		//config.Version = sarama.V2_1_0_0
		config.Producer.Return.Successes = true
		// 启动 producer
		producer, err := sarama.NewSyncProducer(addrs, config)
		if err != nil {
			log.Logger.Errorf("creating sync producer client failed: %v", err)
			return
		}
		pusher.client = producer
	})
	return pusher.client
}
