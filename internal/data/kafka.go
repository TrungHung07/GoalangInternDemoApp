package data

import (
	"context"
	"encoding/json"
	"time"

	"github.com/IBM/sarama" //thư viện client Kafka
	"github.com/go-kratos/kratos/v2/log"
)

// HistoryEvent represents the event structure sent to Kafka
type HistoryEvent struct {
	ID        string                 `json:"id"`
	TableName string                 `json:"table_name"`
	RecordID  string                 `json:"record_id"`
	Action    string                 `json:"action"`
	OldData   map[string]interface{} `json:"old_data,omitempty"`
	NewData   map[string]interface{} `json:"new_data,omitempty"`
	UserID    string                 `json:"user_id,omitempty"` // omitempty nếu giá trị rỗng thì sẽ không đưa vào json
	Timestamp time.Time              `json:"timestamp"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// KafkaProducer handles producing messages to Kafka
type KafkaProducer struct {
	producer sarama.SyncProducer //đối tượng producrer đồng bộ của sarama để gửi message lên kafka
	topic    string              //tên topic Kafka để gửi message vào topic này
	log      *log.Helper         //helper để ghi log
}

// NewKafkaProducer tạo mới một Kafka producer để gử message lên kafka
// brokers : là danh sách các broker kafka mà producer sẽ kết nối tới
// topic : là tên topic mà producer sẽ gửi message vào
func NewKafkaProducer(brokers []string, topic string, logger log.Logger) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Compression = sarama.CompressionGZIP

	producer, err := sarama.NewSyncProducer(brokers, config) //Tạo một producer đồng bộ mới với danh sách broker và cấu hình vừa thiết lập.
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		producer: producer,
		topic:    topic,
		log:      log.NewHelper(logger),
	}, nil
}

// PublishHistoryEvent gửi một event history lên Kafka
func (kp *KafkaProducer) PublishHistoryEvent(ctx context.Context, event *HistoryEvent) error {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		kp.log.Errorf("Failed to marshal history event: %v", err)
		return err
	}

	msg := &sarama.ProducerMessage{ //tạo một message
		Topic: kp.topic,
		Key:   sarama.StringEncoder(event.RecordID), //khóa của mesage là recored id , ép kiểu string recorder
		Value: sarama.ByteEncoder(eventBytes),       //nội dung của message là chuỗi json vừa tọa , ép kiểu byte encoder
		Headers: []sarama.RecordHeader{ // danh sách header của message
			{
				Key:   []byte("table_name"), // tên bảng
				Value: []byte(event.TableName),
			},
			{
				Key:   []byte("action"), // insert , update ,delete
				Value: []byte(event.Action),
			},
		},
	}

	partition, offset, err := kp.producer.SendMessage(msg) // gửi message lên kafka với SendMessage
	// SendMessage sẽ trả về partition và offset của message vừa gửi
	//offset là vị trí của message trong partition
	if err != nil {
		kp.log.Errorf("Failed to publish history event: %v", err)
		return err
	}

	kp.log.Infof("History event published successfully - Partition: %d, Offset: %d", partition, offset)
	return nil
}

// Close producer để giải phóng tài nguyên
func (kp *KafkaProducer) Close() error {
	return kp.producer.Close()
}
