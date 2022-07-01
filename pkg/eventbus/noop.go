package eventbus

import "context"

type (
	NoopRawPublisher struct{}
)

func (n *NoopRawPublisher) Publish(context.Context, []byte) error { return nil }

type (
	NoopConsumerHandler struct{}
)

func (n *NoopConsumerHandler) Listen(context.Context) chan []byte { return nil }
