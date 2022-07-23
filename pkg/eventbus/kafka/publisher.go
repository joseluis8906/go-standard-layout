package kafka

import (
	"context"

	stdkafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joseluis8906/go-standard-layout/pkg/log"
)

type (
	// Publisher ...
	Publisher struct {
		Client *stdkafka.Producer
		Topic  string
	}
)

// Publish ...
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

// NewProducer ...
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
		log.Logger().Fatalf("error trying to connect to kafka: %v", err)
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *stdkafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Logger().Errorf("delivery failed: %s", ev.TopicPartition)
					continue
				}

				log.Logger().Infof("delivered message to %s", ev.TopicPartition)
			}
		}
	}()

	return p
}
