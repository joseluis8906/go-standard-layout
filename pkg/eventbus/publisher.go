package eventbus

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"sync"
)

var (
	ErrNotFound = errors.New("event not found")
)

type (
	RawPublisher interface {
		Publish(context.Context, []byte) error
	}

	Publisher struct {
		client RawPublisher
		events map[string][]byte
		mux    sync.Mutex
	}
)

func NewPublisher(publisher RawPublisher) *Publisher {
	return &Publisher{
		client: publisher,
		events: map[string][]byte{},
	}
}

func NoopPublisher() *Publisher {
	return &Publisher{
		client: &NoopRawPublisher{},
		events: map[string][]byte{},
	}
}

func (ec *Publisher) Add(eventID string, event interface{}) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	ec.mux.Lock()
	defer ec.mux.Unlock()

	ec.events[eventID] = data
	return nil
}

func (ec *Publisher) Get(eventID string) ([]byte, error) {
	event, ok := ec.events[eventID]
	if !ok {
		return nil, ErrNotFound
	}

	ec.mux.Lock()
	defer ec.mux.Unlock()

	delete(ec.events, eventID)
	return event, nil
}

func (ec *Publisher) Discard(ctx context.Context, eventIDs []string) {
	ec.mux.Lock()
	defer ec.mux.Unlock()

	for _, eventID := range eventIDs {
		delete(ec.events, eventID)
	}
}

func (ec *Publisher) Publish(ctx context.Context, eventIDs []string) {
	for _, eventID := range eventIDs {
		event, err := ec.Get(eventID)
		if err != nil {
			log.Print(err)
			continue
		}

		err = ec.client.Publish(ctx, event)
		log.Print(err)
	}
}

var (
	publisher *Publisher
)

func SetPublisher(aPublisher *Publisher) {
	publisher = aPublisher
}

func Add(eventID string, event interface{}) {
	err := publisher.Add(eventID, event)
	if err != nil {
		log.Print(err)
	}
}

func Publish(ctx context.Context, eventIDs []string) {
	publisher.Publish(ctx, eventIDs)
}

func Discard(ctx context.Context, eventIDs []string) {
	publisher.Discard(ctx, eventIDs)
}
