package biz

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

// HistoryEvent represents a history event
type HistoryEvent struct {
	ID        string                 `json:"id"`
	TableName string                 `json:"table_name"`
	RecordID  string                 `json:"record_id"`
	Action    string                 `json:"action"`
	OldData   map[string]interface{} `json:"old_data,omitempty"`
	NewData   map[string]interface{} `json:"new_data,omitempty"`
	UserID    string                 `json:"user_id,omitempty"`
	Timestamp time.Time              `json:"timestamp"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// History represents a history record
type History struct {
	ID        string                 `json:"id"`
	TableName string                 `json:"table_name"`
	RecordID  string                 `json:"record_id"`
	Action    string                 `json:"action"`
	OldData   map[string]interface{} `json:"old_data,omitempty"`
	NewData   map[string]interface{} `json:"new_data,omitempty"`
	UserID    string                 `json:"user_id,omitempty"`
	CreatedAt time.Time              `json:"created_at"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// HistoryRepo defines the interface for history repository
type HistoryRepo interface {
	CreateHistory(ctx context.Context, history *History) (*History, error)
	GetHistoryByRecordID(ctx context.Context, tableName, recordID string) ([]*History, error)
	GetHistoryByTableName(ctx context.Context, tableName string) ([]*History, error)
}

// HistoryEventPublisher defines the interface for publishing history events
type HistoryEventPublisher interface {
	PublishHistoryEvent(ctx context.Context, event *HistoryEvent) error
}

// HistoryUsecase handles history business logic
type HistoryUsecase struct {
	repo      HistoryRepo
	publisher HistoryEventPublisher
	log       *log.Helper
}

// NewHistoryUsecase creates a new history usecase
func NewHistoryUsecase(repo HistoryRepo, publisher HistoryEventPublisher, logger log.Logger) *HistoryUsecase {
	return &HistoryUsecase{
		repo:      repo,
		publisher: publisher,
		log:       log.NewHelper(logger),
	}
}

// PublishHistoryEvent publishes a history event to Kafka
func (hu *HistoryUsecase) PublishHistoryEvent(ctx context.Context, tableName, recordID, action string, oldData, newData map[string]interface{}, userID string, metadata map[string]interface{}) error {
	event := &HistoryEvent{
		ID:        uuid.New().String(),
		TableName: tableName,
		RecordID:  recordID,
		Action:    action,
		OldData:   oldData,
		NewData:   newData,
		UserID:    userID,
		Timestamp: time.Now(),
		Metadata:  metadata,
	}

	return hu.publisher.PublishHistoryEvent(ctx, event)
}

// CreateHistoryFromEvent creates a history record from an event
func (hu *HistoryUsecase) CreateHistoryFromEvent(ctx context.Context, eventData []byte) error {
	var event HistoryEvent
	if err := json.Unmarshal(eventData, &event); err != nil {
		hu.log.Errorf("Failed to unmarshal history event: %v", err)
		return err
	}

	history := &History{
		ID:        event.ID,
		TableName: event.TableName,
		RecordID:  event.RecordID,
		Action:    event.Action,
		OldData:   event.OldData,
		NewData:   event.NewData,
		UserID:    event.UserID,
		CreatedAt: event.Timestamp,
		Metadata:  event.Metadata,
	}

	_, err := hu.repo.CreateHistory(ctx, history)
	if err != nil {
		hu.log.Errorf("Failed to create history record: %v", err)
		return err
	}

	hu.log.Infof("History record created successfully for table: %s, record: %s", history.TableName, history.RecordID)
	return nil
}

// GetHistoryByRecordID retrieves history records by record ID
func (hu *HistoryUsecase) GetHistoryByRecordID(ctx context.Context, tableName, recordID string) ([]*History, error) {
	return hu.repo.GetHistoryByRecordID(ctx, tableName, recordID)
}

// GetHistoryByTableName retrieves history records by table name
func (hu *HistoryUsecase) GetHistoryByTableName(ctx context.Context, tableName string) ([]*History, error) {
	return hu.repo.GetHistoryByTableName(ctx, tableName)
}
