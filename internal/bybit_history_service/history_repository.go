package bybit_history_service

import (
	db "github.com/bxcodec/go-clean-arch/db/postgres"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models"
	"github.com/bxcodec/go-clean-arch/util"
)

type HistoryRepository interface {
	FindBySymbol(userId int64, symbol string, pageIndex int, pageSize int) ([]models.BybitFutureOrderHistory, error)
	FindById(id int64) (models.BybitFutureOrderHistory, error)
	FindByBetweenCreatedTime(userId int64, startTime string, endTime string, pageIndex int, pageSize int) ([]models.BybitFutureOrderHistory, error)
	FindByBetweenCreatedTimeAndSymbol(userId int64, symbol string, startTime string, endTime string, pageIndex int, pageSize int) ([]models.BybitFutureOrderHistory, error)
}
type HistoryRepositoryImpl struct {
	db *db.PostgresDB
}

func NewHistory(db *db.PostgresDB) *HistoryRepositoryImpl {
	return &HistoryRepositoryImpl{
		db: db,
	}
}

func (s *HistoryRepositoryImpl) FindBySymbol(userId int64, symbol string, pageIndex int, pageSize int) ([]models.BybitFutureOrderHistory, error) {
	var historys []models.BybitFutureOrderHistory
	tx := s.db.Conn().Where("symbol = ? AND user_id = ?", symbol, userId).
		Limit(pageSize).Offset(pageIndex).Find(&historys)
	if tx.Error != nil {
		return []models.BybitFutureOrderHistory{}, tx.Error
	}
	return historys, nil
}

func (s *HistoryRepositoryImpl) FindById(id int64) (models.BybitFutureOrderHistory, error) {
	var historys models.BybitFutureOrderHistory
	tx := s.db.Conn().Where("id = ?", id).First(&historys)
	if tx.Error != nil {
		return models.BybitFutureOrderHistory{}, tx.Error
	}
	return historys, nil
}
func (s *HistoryRepositoryImpl) FindByBetweenCreatedTime(userId int64, startTime string, endTime string, pageIndex int, pageSize int) ([]models.BybitFutureOrderHistory, error) {
	start, _ := util.DecodeCursor(startTime)
	end, _ := util.DecodeCursor(endTime)

	var historys []models.BybitFutureOrderHistory
	tx := s.db.Conn().Where("user_id = ? AND created_at>= ? AND created_at<= ?",
		userId, start, end).
		Limit(pageSize).Offset(pageIndex).Find(&historys)
	if tx.Error != nil {
		return []models.BybitFutureOrderHistory{}, tx.Error
	}
	return historys, nil
}
func (s *HistoryRepositoryImpl) FindByBetweenCreatedTimeAndSymbol(userId int64, symbol string, startTime string, endTime string, pageIndex int, pageSize int) ([]models.BybitFutureOrderHistory, error) {
	start, _ := util.DecodeCursor(startTime)
	end, _ := util.DecodeCursor(endTime)
	var historys []models.BybitFutureOrderHistory
	tx := s.db.Conn().Where("user_id = ? AND symbol = ? AND created_at>= ? AND created_at<= ?",
		userId, symbol, start, end).
		Limit(pageSize).Offset(pageIndex).Find(&historys)
	if tx.Error != nil {
		return []models.BybitFutureOrderHistory{}, tx.Error
	}
	return historys, nil
}
