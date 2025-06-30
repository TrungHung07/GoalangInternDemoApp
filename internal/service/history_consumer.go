package service

import (
	"context"
	"sync"

	"DemoApp/internal/biz"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
)
// KafkaBrokers represents a list of Kafka broker addresses used for connecting to the Kafka cluster.
type KafkaBrokers []string

// KafkaGroupID represents the identifier of a Kafka consumer group.
type KafkaGroupID string

// KafkaTopic represents the name of a Kafka topic used for publishing or consuming messages.
type KafkaTopic string

// HistoryConsumer handles consuming history events from Kafka
type HistoryConsumer struct {
	consumer       sarama.ConsumerGroup
	historyUsecase *biz.HistoryUsecase
	topic          string
	groupID        string
	log            *log.Helper
	ctx            context.Context
	cancel         context.CancelFunc
	wg             sync.WaitGroup
}

// NewHistoryConsumer creates a new Kafka consumer for history events
func NewHistoryConsumer(brokers KafkaBrokers, groupID KafkaGroupID, topic KafkaTopic, historyUsecase *biz.HistoryUsecase, logger log.Logger) (*HistoryConsumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Session.Timeout = 10000000000   // 10 seconds
	config.Consumer.Group.Heartbeat.Interval = 3000000000 // 3 seconds

	consumer, err := sarama.NewConsumerGroup([]string(brokers), string(groupID), config)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &HistoryConsumer{
		consumer:       consumer,
		historyUsecase: historyUsecase,
		topic:          string(topic),
		groupID:        string(groupID),
		log:            log.NewHelper(logger),
		ctx:            ctx,
		cancel:         cancel,
	}, nil
}

// Start starts the consumer
// chạy consumer trong một goroutine riêng biệt
// để có thể liên tục lắng nghe các message từ Kafka
// và xử lý chúng bằng phương thức ConsumeClaim
// Lưu ý: consumer sẽ chạy mãi cho đến khi context bị hủy
func (hc *HistoryConsumer) Start() {
	hc.wg.Add(1)
	go func() {
		defer hc.wg.Done()
		for {
			select {
			case <-hc.ctx.Done():
				hc.log.Info("History consumer context cancelled")
				return
			default:
				err := hc.consumer.Consume(hc.ctx, []string{hc.topic}, hc)
				if err != nil {
					hc.log.Errorf("Error consuming messages: %v", err)
					return
				}
			}
		}
	}()

	hc.log.Infof("History consumer started for topic: %s, group: %s", hc.topic, hc.groupID)
}

// Stop stops the consumer
func (hc *HistoryConsumer) Stop() {
	hc.log.Info("Stopping history consumer...")
	hc.cancel()
	hc.wg.Wait()

	if err := hc.consumer.Close(); err != nil {
		hc.log.Errorf("Error closing consumer: %v", err)
	}

	hc.log.Info("History consumer stopped")
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (hc *HistoryConsumer) Setup(sarama.ConsumerGroupSession) error {
	hc.log.Info("History consumer session setup")
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (hc *HistoryConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	hc.log.Info("History consumer session cleanup")
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages()
func (hc *HistoryConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			if message == nil {
				return nil
			}

			hc.log.Infof("Processing history event from partition %d, offset %d", message.Partition, message.Offset)

			// Process the message
			if err := hc.historyUsecase.CreateHistoryFromEvent(session.Context(), message.Value); err != nil {
				hc.log.Errorf("Failed to process history event: %v", err)
				// Depending on your error handling strategy, you might want to skip or retry
				// For now, we'll mark it as processed to avoid infinite retries
			}

			// Mark message as processed
			session.MarkMessage(message, "")

		case <-session.Context().Done():
			hc.log.Info("Consumer claim context cancelled")
			return nil
		}
	}
}
