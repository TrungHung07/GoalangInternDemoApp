package service

import (
	"context"

	"DemoApp/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// HistoryService provides history-related services
type HistoryService struct {
	historyUsecase *biz.HistoryUsecase
	log            *log.Helper
}

// NewHistoryService creates a new history service
func NewHistoryService(historyUsecase *biz.HistoryUsecase, logger log.Logger) *HistoryService {
	return &HistoryService{
		historyUsecase: historyUsecase,
		log:            log.NewHelper(logger),
	}
}

// GetHistoryByRecordID retrieves history records by record ID
func (hs *HistoryService) GetHistoryByRecordID(ctx context.Context, tableName, recordID string) ([]*biz.History, error) {
	return hs.historyUsecase.GetHistoryByRecordID(ctx, tableName, recordID)
}

// GetHistoryByTableName retrieves history records by table name
func (hs *HistoryService) GetHistoryByTableName(ctx context.Context, tableName string) ([]*biz.History, error) {
	return hs.historyUsecase.GetHistoryByTableName(ctx, tableName)
}
