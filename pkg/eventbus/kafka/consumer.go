package kafka

import (
	"context"
	"time"

	stdkafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joseluis8906/go-standard-layout/pkg/log"
)

type (
	Consumer struct {
		Client      *stdkafka.Consumer
		MsgCount    int
		ReadTimeout time.Duration
	}
)

func (c *Consumer) Listen(ctx context.Context) chan []byte {
	msgs := make(chan []byte, c.MsgCount)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(msgs)
				return
			default:
				msg, err := c.Client.ReadMessage(c.ReadTimeout)
				if err != nil {
					kafkaErr, ok := err.(stdkafka.Error)
					if ok && kafkaErr.Code() != stdkafka.ErrTimedOut {
						log.Logger().Errorf("Consumer error: %v", err)
						continue
					}
				}

				if msg == nil {
					continue
				}

				msgs <- msg.Value

			}
		}
	}()

	return msgs
}

func NewConsumer(opts ...OptionFunc) *stdkafka.Consumer {
	conf := &config{}
	for _, opt := range opts {
		opt(conf)
	}

	confMap := &stdkafka.ConfigMap{
		"bootstrap.servers": conf.bootstrapServers,
		"group.id":          conf.groupID,
		"auto.offset.reset": conf.autoOffsetReset,
	}

	c, err := stdkafka.NewConsumer(confMap)
	if err != nil {
		log.Logger().Fatalf("error trying to connect to kafka: %v", err)
	}

	err = c.SubscribeTopics(conf.consumerTopics, nil)
	if err != nil {
		log.Logger().Fatalf("error trying to subscribe to kafka: %v", err)
	}

	return c
}
