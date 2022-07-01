package mongo

import (
	"context"
	"fmt"

	"github.com/joseluis8906/go-standard-layout/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(ctx context.Context, opts ...OptionFunc) *mongo.Client {
	conf := &config{}
	for _, opt := range opts {
		opt(conf)
	}

	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/%s",
		conf.user,
		conf.passwd,
		conf.host,
		conf.port,
		conf.db,
	)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("error trying to connect to mongo: %v", err)
	}

	return client
}
