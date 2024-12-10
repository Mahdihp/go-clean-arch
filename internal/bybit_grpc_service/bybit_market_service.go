package bybit_grpc_service

import (
	"context"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/db/mongodb"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"log"
)

type ByBitMarketRepository interface {
	FindBySymbol(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error)
}
type ByBitMarketService struct {
	cfg     config.MongoDb
	mongodb *mongodb.MongoDb
}

func NewByBitMarketService(cfg config.MongoDb) ByBitMarketService {
	return ByBitMarketService{
		cfg:     cfg,
		mongodb: mongodb.NewMongoDb(cfg),
	}
}
func (s *ByBitMarketService) FindBySymbol(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error) {
	var linears []models_grpc.ByBitMarketGetInstrumentsInfoLinear
	collection := s.mongodb.MongoConn().Collection(collectionName)

	filter := bson.D{{"symbol", bson.D{{"$eq", symbol}}}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return []models_grpc.ByBitMarketGetInstrumentsInfoLinear{}, err
	}
	for cursor.Next(ctx) {
		var newItem models_grpc.ByBitMarketGetInstrumentsInfoLinear
		err := cursor.Decode(&newItem)
		if err != nil {
			log.Fatal(err)
			return []models_grpc.ByBitMarketGetInstrumentsInfoLinear{}, err
		}

		linears = append(linears, newItem)
	}
	cursor.Close(ctx)
	return linears, nil
}
