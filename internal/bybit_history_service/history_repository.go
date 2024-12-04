package bybit_history_service

import (
	"context"
	"fmt"
	db "github.com/bxcodec/go-clean-arch/db/postgres"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models/ent"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models/ent/bybithistory"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models/ent/bybituser"
	"github.com/bxcodec/go-clean-arch/util"
	"strconv"
)

type HistoryRepository interface {
	FindById(ctx context.Context, id int64) (ent.ByBitHistory, error)
	FindBySymbol(ctx context.Context, userId int64, symbol string) ([]ent.ByBitHistory, error)
	FindByBetweenCreatedTime(ctx context.Context, userId int64, startTime int32, endTime int32) ([]ent.ByBitHistory, error)
	FindByBetweenCreatedTimeAndSymbol(ctx context.Context, userId int64, symbol string, startTime int32, endTime int32) ([]ent.ByBitHistory, error)
}
type HistoryRepositoryImpl struct {
	db *db.PostgresDB
}

func NewHistory(db *db.PostgresDB) *HistoryRepositoryImpl {
	return &HistoryRepositoryImpl{
		db: db,
	}
}
func (s *HistoryRepositoryImpl) FindById(ctx context.Context, id int64) (ent.ByBitHistory, error) {
	found, err := s.db.Conn().ByBitHistory.
		Query().Where(bybithistory.ID(id)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ent.ByBitHistory{}, err
		}
	}
	return *found, nil
}
func (s *HistoryRepositoryImpl) FindBySymbol(ctx context.Context, userId int64, symbol string) ([]ent.ByBitHistory, error) {

	founds, err := s.db.Conn().ByBitHistory.
		Query().Where(bybithistory.Symbol(symbol), bybithistory.UserID(userId)).All(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return []ent.ByBitHistory{}, err
		}
	}
	fmt.Println(len(founds))
	foundSlice := make([]ent.ByBitHistory, len(founds))
	for i, found := range founds {
		foundSlice[i] = *found
	}
	return foundSlice, nil
}
func (s *HistoryRepositoryImpl) FindByBetweenCreatedTime(ctx context.Context, userId int64, startTime int32, endTime int32) ([]ent.ByBitHistory, error) {
	start, _ := util.DecodeCursor(strconv.Itoa(int(startTime)))
	end, _ := util.DecodeCursor(strconv.Itoa(int(endTime)))

	founds, err := s.db.Conn().ByBitHistory.
		Query().
		WithUser(func(query *ent.ByBitUserQuery) {
			query.Where(bybituser.ID(userId))
		}).
		Where(bybithistory.CreatedTimeGTE(start), bybithistory.CreatedTimeLTE(end)).All(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return []ent.ByBitHistory{}, err
		}
	}
	foundSlice := make([]ent.ByBitHistory, len(founds))
	for i, found := range founds {
		foundSlice[i] = *found
	}

	return foundSlice, nil
}
func (s *HistoryRepositoryImpl) FindByBetweenCreatedTimeAndSymbol(ctx context.Context, userId int64, symbol string, startTime int32, endTime int32) ([]ent.ByBitHistory, error) {
	start, _ := util.DecodeCursor(string(startTime))
	end, _ := util.DecodeCursor(string(endTime))

	founds, err := s.db.Conn().ByBitHistory.
		Query().
		WithUser(func(query *ent.ByBitUserQuery) {
			query.Where(bybituser.ID(userId))
		}).
		Where(bybithistory.Symbol(symbol),
			bybithistory.CreatedTimeGTE(start),
			bybithistory.CreatedTimeLTE(end)).All(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return []ent.ByBitHistory{}, err
		}
	}
	foundSlice := make([]ent.ByBitHistory, len(founds))
	for i, found := range founds {
		foundSlice[i] = *found
	}

	return foundSlice, nil
}
