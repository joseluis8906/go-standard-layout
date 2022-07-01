package kafka

import (
	"context"

	stdkafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joseluis8906/go-standard-layout/pkg/log"
)

type (
	Publisher struct {
		Client *stdkafka.Producer
		Topic  string
	}
)

func (p *Publisher) Publish(_ context.Context, msj []byte) error {
	err := p.Client.Produce(&stdkafka.Message{
		TopicPartition: stdkafka.TopicPartition{
			Topic:     &p.Topic,
			Partition: stdkafka.PartitionAny,
		},
		Value: msj,
	}, nil)

	return err
}

func NewProducer(opts ...OptionFunc) *stdkafka.Producer {
	conf := &config{}
	for _, opt := range opts {
		opt(conf)
	}

	confMap := &stdkafka.ConfigMap{
		"bootstrap.servers": conf.bootstrapServers,
	}

	p, err := stdkafka.NewProducer(confMap)
	if err != nil {
		log.Fatalf("error trying to connect to kafka: %v", err)
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *stdkafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Errorf("delivery failed: %s", ev.TopicPartition)
					continue
				}

				log.Info("delivered message to %s", ev.TopicPartition)
			}
		}
	}()

	return p
}
