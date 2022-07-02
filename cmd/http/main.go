package main

import (
	"context"
	"fmt"
	"math/rand"
	stdhttp "net/http"
	"time"

	"github.com/joseluis8906/go-standard-layout/internal/application/commands"
	"github.com/joseluis8906/go-standard-layout/internal/application/queries"
	"github.com/joseluis8906/go-standard-layout/internal/infrastructure/inmemory"
	"github.com/joseluis8906/go-standard-layout/pkg/config"
	"github.com/joseluis8906/go-standard-layout/pkg/eventbus"
	"github.com/joseluis8906/go-standard-layout/pkg/eventbus/kafka"
	"github.com/joseluis8906/go-standard-layout/pkg/http"
	"github.com/joseluis8906/go-standard-layout/pkg/log"
	"github.com/joseluis8906/go-standard-layout/pkg/mongo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

const (
	port = 8000
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	config.InitViper("./configs", "go-standard-layout", "yml")

	logger := log.NewLogrus(
		log.Caller(viper.GetBool("log.caller")),
		log.Env(viper.GetString("environment")),
		log.Formatter(viper.GetString("log.formatter")),
		log.Level(viper.GetInt("log.level")),
	)

	log.SetLogger(logger)

	fluentHook := log.NewFluentHook(
		log.FluentdHost(viper.GetString("log.fluentd.host")),
		log.FluentdPort(viper.GetInt("log.fluentd.port")),
		log.FluentdAsync(true),
	)

	defer fluentHook.Close()
	logger.Logger.AddHook(fluentHook)

	mongodbClient := mongo.NewMongo(
		ctx,
		mongo.DB(viper.GetString("mongo.db")),
		mongo.Host(viper.GetString("mongo.host")),
		mongo.Passwd(viper.GetString("mongo.passwd")),
		mongo.Port(viper.GetInt("mongo.port")),
		mongo.User(viper.GetString("mongo.user")),
	)

	defer func() {
		if err := mongodbClient.Disconnect(ctx); err != nil {
			logger.Error(err)
		}
	}()

	kafkaProducer := kafka.NewProducer(
		kafka.BootstrapServers(viper.GetString("kafka.bootstrap.servers")),
	)

	defer kafkaProducer.Close()

	eventbus.SetPublisher(eventbus.NewPublisher(&kafka.Publisher{
		Client: kafkaProducer,
		Topic:  viper.GetString("kafka.producer.topic"),
	}))

	r := http.NewChiRouter()
	r.Handle("/metrics", promhttp.Handler())

	postRepo := inmemory.NewPostRepository()

	addPost := commands.AddPostHandler{
		PostPersistor: postRepo,
	}
	r.HandleFunc("/addPost", addPost.HandleFunc)

	getAllPosts := queries.GetAllPostHandler{
		PostFinder: postRepo,
	}
	r.HandleFunc("/getAllPosts", getAllPosts.HandleFunc)

	bind := fmt.Sprintf("%s:%d", viper.GetString("http.server.address"), viper.GetInt("http.server.port"))
	log.Info("http server is listening on: %s", bind)
	log.Fatal(stdhttp.ListenAndServe(bind, r))
}
