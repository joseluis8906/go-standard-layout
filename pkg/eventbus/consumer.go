package eventbus

import (
	"context"
	"sync"
)

type (
	ConsumerFunc func(context.Context, []byte)

	ConsumerHandler interface {
		Listen(context.Context) chan []byte
	}

	Consumer struct {
		consumerFuncs   []ConsumerFunc
		consumerHandler ConsumerHandler
	}
)

func NewConsumer(consumerHandler ConsumerHandler) *Consumer {
	return &Consumer{
		consumerHandler: consumerHandler,
		consumerFuncs:   []ConsumerFunc{},
	}
}

func NoopConsumer() *Consumer {
	return &Consumer{
		consumerHandler: &NoopConsumerHandler{},
		consumerFuncs:   []ConsumerFunc{},
	}
}

func (er *Consumer) AddConsumerFunc(r ConsumerFunc) {
	er.consumerFuncs = append(er.consumerFuncs, r)
}

func (er *Consumer) Listen(ctx context.Context) {
	wg := sync.WaitGroup{}
	for msg := range er.consumerHandler.Listen(ctx) {
		for _, r := range er.consumerFuncs {
			wg.Add(1)

			go func(r ConsumerFunc) {
				defer wg.Done()
				r(ctx, msg)
			}(r)
		}

		wg.Wait()
	}
}

var (
	consumer *Consumer
)

func SetConsumer(c *Consumer) {
	consumer = c
}

func Listen(ctx context.Context) {
	consumer.Listen(ctx)
}

func AddConsumerFunc(r ConsumerFunc) {
	consumer.AddConsumerFunc(r)
}
