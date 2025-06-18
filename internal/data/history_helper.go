package data

import (
	"context"
	"encoding/json"

	"DemoApp/internal/biz"
)

// HistoryHelper provides utility functions for tracking data changes
type HistoryHelper struct {
	historyUsecase *biz.HistoryUsecase
}

// NewHistoryHelper creates a new history helper
func NewHistoryHelper(historyUsecase *biz.HistoryUsecase) *HistoryHelper {
	return &HistoryHelper{
		historyUsecase: historyUsecase,
	}
}

// TrackInsert tracks an insert operation
func (hh *HistoryHelper) TrackInsert(ctx context.Context, tableName, recordID string, newData interface{}, userID string) error {
	newDataMap, err := hh.structToMap(newData)
	if err != nil {
		return err
	}

	return hh.historyUsecase.PublishHistoryEvent(
		ctx,
		tableName,
		recordID,
		"INSERT",
		nil,
		newDataMap,
		userID,
		nil,
	)
}

// TrackUpdate tracks an update operation
func (hh *HistoryHelper) TrackUpdate(ctx context.Context, tableName, recordID string, oldData, newData interface{}, userID string) error {
	oldDataMap, err := hh.structToMap(oldData)
	if err != nil {
		return err
	}

	newDataMap, err := hh.structToMap(newData)
	if err != nil {
		return err
	}

	return hh.historyUsecase.PublishHistoryEvent(
		ctx,
		tableName,
		recordID,
		"UPDATE",
		oldDataMap,
		newDataMap,
		userID,
		nil,
	)
}

// TrackDelete tracks a delete operation
func (hh *HistoryHelper) TrackDelete(ctx context.Context, tableName, recordID string, oldData interface{}, userID string) error {
	oldDataMap, err := hh.structToMap(oldData)
	if err != nil {
		return err
	}

	return hh.historyUsecase.PublishHistoryEvent(
		ctx,
		tableName,
		recordID,
		"DELETE",
		oldDataMap,
		nil,
		userID,
		nil,
	)
}

// structToMap converts a struct to map[string]interface{}
func (hh *HistoryHelper) structToMap(data interface{}) (map[string]interface{}, error) {
	if data == nil {
		return nil, nil
	}

	// Convert to JSON and back to map to handle nested structs and custom types
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Example usage in your greeter repository
// func (gr *greeterRepo) CreateGreeter(ctx context.Context, greeter *biz.Greeter) (*biz.Greeter, error) {
//     // Create the record
//     result, err := gr.data.db.Greeter.Create().
//         SetName(greeter.Name).
//         SetMessage(greeter.Message).
//         Save(ctx)
//     if err != nil {
//         return nil, err
//     }

//     // Convert to biz entity
//     bizGreeter := convertEntGreeterToBiz(result)

//     // Track the insert operation
//     if gr.historyHelper != nil {
//         if err := gr.historyHelper.TrackInsert(ctx, "greeter", result.ID, bizGreeter, getUserIDFromContext(ctx)); err != nil {
//             gr.log.Errorf("Failed to track insert history: %v", err)
//             // Don't fail the main operation, just log the error
//         }
//     }

//     return bizGreeter, nil
// }
