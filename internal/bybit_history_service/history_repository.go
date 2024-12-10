package bybit_history_service

import (
	"context"
	"errors"
	db "github.com/bxcodec/go-clean-arch/db/postgres"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models"
	"github.com/bxcodec/go-clean-arch/util"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type HistoryRepository interface {
	FindBySymbol(ctx context.Context, collectionName string, userId string, symbol string, pageIndex int, pageSize int) ([]models.BybitFutureOrderHistory, error)
	FindById(ctx context.Context, collectionName string, id string) (models.BybitFutureOrderHistory, error)
	FindByBetweenCreatedTime(ctx context.Context, collectionName string, userId string, startTime string, endTime string, pageIndex int, pageSize int) ([]models.BybitFutureOrderHistory, error)
	FindByBetweenCreatedTimeAndSymbol(ctx context.Context, collectionName string, userId string, symbol string, startTime string, endTime string, pageIndex int, pageSize int) ([]models.BybitFutureOrderHistory, error)
}
type HistoryRepositoryImpl struct {
	db *db.PostgresDB
}

func NewHistory(db *db.PostgresDB) *HistoryRepositoryImpl {
	return &HistoryRepositoryImpl{
		db: db,
	}
}

func (s *HistoryRepositoryImpl) FindBySymbol(ctx context.Context, collectionName string, userId string, symbol string, pageIndex int, pageSize int) ([]models.BybitFutureOrderHistory, error) {
	var historys []models.BybitFutureOrderHistory
	collection := s.db.MongoConn().Collection(collectionName)

	filter := bson.D{{"user_id", userId},
		{"symbol", symbol}}

	findOptions := options.Find().SetLimit(int64(pageSize)).SetSkip(int64(pageIndex))

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var user models.BybitFutureOrderHistory
		if err := cursor.Decode(user); err != nil {
			logrus.Fatal(err)
			return nil, err
		}
		historys = append(historys, user)
	}
	cursor.Close(ctx)
	return historys, nil
}

func (s *HistoryRepositoryImpl) FindById(ctx context.Context, collectionName string, id string) (models.BybitFutureOrderHistory, error) {
	var historys models.BybitFutureOrderHistory
	collection := s.db.MongoConn().Collection(collectionName)

	//filter := bson.D{{"_id", id}}
	filter := bson.D{{"_id", bson.D{{"$eq", id}}}}

	err := collection.FindOne(ctx, filter).Decode(historys)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return models.BybitFutureOrderHistory{}, err
	}
	return historys, nil
}
func (s *HistoryRepositoryImpl) FindByBetweenCreatedTime(ctx context.Context, collectionName string, userId string, startTime string, endTime string, pageIndex int, pageSize int) ([]models.BybitFutureOrderHistory, error) {
	start, _ := util.DecodeCursor(startTime)
	end, _ := util.DecodeCursor(endTime)
	var historys []models.BybitFutureOrderHistory

	collection := s.db.MongoConn().Collection(collectionName)
	filter := bson.D{{"created_at", bson.D{{"$gt", start}, {"$lt", end}}}}

	findOptions := options.Find().SetLimit(int64(pageSize)).SetSkip(int64(pageIndex))
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var user models.BybitFutureOrderHistory
		if err := cursor.Decode(user); err != nil {
			logrus.Fatal(err)
			return nil, err
		}
		historys = append(historys, user)
	}
	cursor.Close(ctx)
	return historys, nil
}
func (s *HistoryRepositoryImpl) FindByBetweenCreatedTimeAndSymbol(ctx context.Context, collectionName string, userId string, symbol string, startTime string, endTime string, pageIndex int, pageSize int) ([]models.BybitFutureOrderHistory, error) {
	start, _ := util.DecodeCursor(startTime)
	end, _ := util.DecodeCursor(endTime)
	var historys []models.BybitFutureOrderHistory

	collection := s.db.MongoConn().Collection(collectionName)
	filter := bson.D{{"created_at", bson.D{{"$gt", start}, {"$lt", end}}},
		{"user_id", bson.D{{"$eq", userId}}},
		{"symbol", bson.D{{"$eq", symbol}}}}

	findOptions := options.Find().SetLimit(int64(pageSize)).SetSkip(int64(pageIndex))
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var user models.BybitFutureOrderHistory
		if err := cursor.Decode(user); err != nil {
			logrus.Fatal(err)
			return nil, err
		}
		historys = append(historys, user)
	}
	cursor.Close(ctx)
	return historys, nil
}
