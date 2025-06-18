package data

import (
	"DemoApp/ent"
	"DemoApp/internal/biz"
	"DemoApp/internal/conf"
	"context"

	//	event "DemoApp/internal/event"
	"fmt"
	"os"

	"entgo.io/ent/dialect"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	_ "github.com/lib/pq"
)

func ProvideKafkaBrokers() []string {
	return []string{"localhost:9092"}
}

func ProvideKafkaTopic() string {
	return "history-topic"
}

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData,
	NewRedisClient,
	NewRedisConfig,
	NewGreeterRepo,
	NewKafkaProducer,
	NewKafkaHistoryPublisher,
	NewHistoryHelper,
	NewHistoryRepo,
	ProvideKafkaBrokers,
	ProvideKafkaTopic,
)

// Data .
type Data struct {
	// TODO wrapped database client
	DB    *ent.Client
	Redis *redis.Client
	Kafka *KafkaProducer
	// HistoryHelper *HistoryHelper
}

func getPostgresSourceFromEnv() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if sslmode == "" {
		sslmode = "disable"
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, pass, dbname, port, sslmode)
}

// NewData .
func NewData(c *conf.Data, redisClient *redis.Client, logger log.Logger, kafkaProducer *KafkaProducer) (*Data, func(), error) {
	dsn := getPostgresSourceFromEnv()
	client, err := ent.Open(dialect.Postgres, dsn, ent.Debug())

	if err != nil {
		return nil, nil, fmt.Errorf("failed to open database connection: %w", err)
	}
	log.NewHelper(logger).Info("Successfully connected to db", "dsn", dsn)

	// repo := event.NewHistoryRepo(&Data{DB: client})
	// event.StartKafkaConsumer(
	// 	[]string{"localhost:9092"}, // Replace with your Kafka brokers
	// 	"history-topic",            // Replace with your Kafka topic
	// 	repo,
	// )
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		client.Close()
		redisClient.Close()
		if err := kafkaProducer.Close(); err != nil {
			log.NewHelper(logger).Error(err)
		}
	}
	return &Data{
		DB:    client,
		Redis: redisClient,
		Kafka: kafkaProducer,
		// HistoryHelper: historyHelper,
	}, cleanup, nil
}

type KafkaHistoryPublisher struct {
	producer *KafkaProducer
}

// NewKafkaHistoryPublisher creates a new Kafka history publisher
func NewKafkaHistoryPublisher(producer *KafkaProducer) biz.HistoryEventPublisher {
	return &KafkaHistoryPublisher{
		producer: producer,
	}
}

// PublishHistoryEvent implements HistoryEventPublisher interface
func (khp *KafkaHistoryPublisher) PublishHistoryEvent(ctx context.Context, event *biz.HistoryEvent) error {
	// Convert biz.HistoryEvent to data.HistoryEvent
	dataEvent := &HistoryEvent{
		ID:        event.ID,
		TableName: event.TableName,
		RecordID:  event.RecordID,
		Action:    event.Action,
		OldData:   event.OldData,
		NewData:   event.NewData,
		UserID:    event.UserID,
		Timestamp: event.Timestamp,
		Metadata:  event.Metadata,
	}
	return khp.producer.PublishHistoryEvent(ctx, dataEvent)
}
