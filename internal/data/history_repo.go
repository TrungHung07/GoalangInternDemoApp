package data

import (
	"DemoApp/ent"
	"DemoApp/ent/history"
	"DemoApp/internal/biz"
	"context"
)

type historyRepo struct {
	data *Data
}

// NewHistoryRepo creates and returns a new HistoryRepo instance using the provided data source.
func NewHistoryRepo(data *Data) biz.HistoryRepo {
	return &historyRepo{data: data}
}

func (r *historyRepo) CreateHistory(ctx context.Context, h *biz.History) (*biz.History, error) {
	_, err := r.data.DB.History.Create().
		SetID(h.ID). // h.ID is already string (uuid string)
		SetTableName(h.TableName).
		SetRecordID(h.RecordID).
		SetAction(h.Action).
		SetOldData(h.OldData).
		SetNewData(h.NewData).
		SetUserID(h.UserID).
		SetCreatedAt(h.CreatedAt).
		SetMetadata(h.Metadata).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	return h, nil
}

func (r *historyRepo) GetHistoryByRecordID(ctx context.Context, tableName, recordID string) ([]*biz.History, error) {
	rows, err := r.data.DB.History.Query().
		Where(
			history.TableNameEQ(tableName),
			history.RecordIDEQ(recordID),
		).
		Order(ent.Desc(history.FieldCreatedAt)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return r.toBiz(rows), nil
}

func (r *historyRepo) GetHistoryByTableName(ctx context.Context, tableName string) ([]*biz.History, error) {
	rows, err := r.data.DB.History.Query().
		Where(history.TableNameEQ(tableName)).
		Order(ent.Desc(history.FieldCreatedAt)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return r.toBiz(rows), nil
}

// Helper để chuyển từ ent -> biz
func (r *historyRepo) toBiz(rows []*ent.History) []*biz.History {
	var result []*biz.History
	for _, row := range rows {
		result = append(result, &biz.History{
			ID:        row.ID,
			TableName: row.TableName,
			RecordID:  row.RecordID,
			Action:    row.Action,
			OldData:   row.OldData,
			NewData:   row.NewData,
			UserID:    row.UserID,
			CreatedAt: row.CreatedAt,
			Metadata:  row.Metadata,
		})
	}
	return result
}
