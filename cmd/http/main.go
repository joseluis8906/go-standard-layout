package main

import (
	"context"
	"fmt"
	"math/rand"
	stdhttp "net/http"
	"time"

	"github.com/joseluis8906/go-standard-layout/internal/application/commands"
	appcmdhttp "github.com/joseluis8906/go-standard-layout/internal/application/commands/http"
	"github.com/joseluis8906/go-standard-layout/internal/application/queries"
	appqtyhttp "github.com/joseluis8906/go-standard-layout/internal/application/queries/http"
	"github.com/joseluis8906/go-standard-layout/internal/infrastructure/inmemory"
	"github.com/joseluis8906/go-standard-layout/pkg/config"
	"github.com/joseluis8906/go-standard-layout/pkg/eventbus"
	"github.com/joseluis8906/go-standard-layout/pkg/eventbus/kafka"
	"github.com/joseluis8906/go-standard-layout/pkg/http"
	"github.com/joseluis8906/go-standard-layout/pkg/log"
	"github.com/joseluis8906/go-standard-layout/pkg/log/fluentd"
	"github.com/joseluis8906/go-standard-layout/pkg/log/logrus"
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

	logger := logrus.New(
		logrus.Env(viper.GetString("environment")),
		logrus.Formatter(viper.GetString("log.formatter")),
		logrus.Level(viper.GetInt("log.level")),
	)

	log.SetLogger(logger)

	fluentHook := fluentd.NewHook(
		fluentd.Host(viper.GetString("log.fluentd.host")),
		fluentd.Port(viper.GetInt("log.fluentd.port")),
		fluentd.Async(true),
	)

	defer fluentHook.Close()
	logger.AddHook(fluentHook)

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

	addPost := appcmdhttp.AddPost{
		Handler: commands.AddPostHandler{
			PostPersistor: postRepo,
		},
	}

	r.Post("/addPost", addPost.HandleFunc)

	getAllPosts := appqtyhttp.GetAllPost{
		Handler: queries.GetAllPostHandler{
			PostFinder: postRepo,
		},
	}

	r.Get("/allPosts", getAllPosts.HandleFunc)

	getNextID := appqtyhttp.GetNextID{
		Handler: queries.GetNextIDHandler{},
	}

	r.Get("/nextID", getNextID.HandleFunc)

	getPost := appqtyhttp.GetPost{
		Handler: queries.GetPostHandler{
			PostFinder: postRepo,
		},
	}

	r.Get("/post/{id}", getPost.HandleFunc)

	bind := fmt.Sprintf("%s:%d", viper.GetString("http.server.address"), viper.GetInt("http.server.port"))
	log.Infof("http server is listening on: %s", bind)
	log.Fatal(stdhttp.ListenAndServe(bind, r))
}
