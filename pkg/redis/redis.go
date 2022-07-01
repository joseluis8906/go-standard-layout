package redis

import (
	"context"
	"fmt"

	stdredis "github.com/go-redis/redis/v8"
)

func NewClient(ctx context.Context, opts ...OptionFunc) *stdredis.Client {
	conf := &config{}
	for _, opt := range opts {
		opt(conf)
	}

	client := stdredis.NewClient(&stdredis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.host, conf.port),
		Password: conf.passwd,
		DB:       conf.db,
	})

	pong := client.Ping(ctx)
	if err := pong.Err(); err != nil {
		panic(fmt.Sprintf("error trying to connect redis: %v", err))
	}

	return client
}