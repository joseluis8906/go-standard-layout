package grpc

import (
	"fmt"

	"github.com/joseluis8906/go-standard-layout/pkg/log"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	stdgrpc "google.golang.org/grpc"
)

func NewClientConn(opts ...OptionFunc) *stdgrpc.ClientConn {
	conf := &config{}
	for _, opt := range opts {
		opt(conf)
	}

	addr := fmt.Sprintf("%s:%d", conf.host, conf.port)
	conn, err := stdgrpc.Dial(addr,
		stdgrpc.WithInsecure(),
		stdgrpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		stdgrpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if err != nil {
		log.Fatalf("error trying to open grpc sales client: %v", err)
	}

	return conn
}
