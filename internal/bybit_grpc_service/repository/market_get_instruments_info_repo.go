package repository

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/db/mongodb"
	"github.com/bxcodec/go-clean-arch/db/redis"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
	"github.com/bxcodec/go-clean-arch/params"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

type ByBitMarketRepository struct {
	Mongodb *mongodb.MongoDb
	Redisdb redis.RedisDb
}

func New(cfg config.Config) ByBitMarketRepository {
	return ByBitMarketRepository{
		Mongodb: mongodb.NewMongoDb(cfg.MongoDbMarket),
		Redisdb: redis.NewRedis(cfg.RedisMarket),
	}
}
func (receiver ByBitMarketRepository) GetCountCollecton(ctx context.Context, collName string) int64 {
	collection := receiver.Mongodb.MongoConn().Collection(collName)
	filter := bson.D{{}}
	documents, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Fatal("CountDocuments error:", err)
		return 0
	}
	return documents
}

func (s *ByBitMarketRepository) FindAllInverse(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoInverse, error) {
	collection := s.Mongodb.MongoConn().Collection(collectionName)
	var spots []models_grpc.ByBitMarketGetInstrumentsInfoInverse

	opts := options.Find().SetSort(bson.D{{params.Field_Search_CreatedAt, -1}})
	filter := bson.D{{}}
	if symbol != "" {
		filter = bson.D{{params.Field_Search_Symbol, bson.D{{params.Equal_Opt, symbol}}}}
	}

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return []models_grpc.ByBitMarketGetInstrumentsInfoInverse{}, err
	}
	if err = cursor.All(ctx, &spots); err != nil {
		return []models_grpc.ByBitMarketGetInstrumentsInfoInverse{}, err
	}
	return spots, nil
}
func (s *ByBitMarketRepository) FindAllSpot(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoSpot, error) {
	collection := s.Mongodb.MongoConn().Collection(collectionName)
	var spots []models_grpc.ByBitMarketGetInstrumentsInfoSpot

	opts := options.Find().SetSort(bson.D{{params.Field_Search_CreatedAt, -1}})
	filter := bson.D{{}}
	if symbol != "" {
		filter = bson.D{{params.Field_Search_Symbol, bson.D{{params.Equal_Opt, symbol}}}}
	}

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return []models_grpc.ByBitMarketGetInstrumentsInfoSpot{}, err
	}
	if err = cursor.All(ctx, &spots); err != nil {
		return []models_grpc.ByBitMarketGetInstrumentsInfoSpot{}, err
	}
	return spots, nil
}
func (s *ByBitMarketRepository) FindAllLinear(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error) {
	collection := s.Mongodb.MongoConn().Collection(collectionName)
	var linears []models_grpc.ByBitMarketGetInstrumentsInfoLinear

	opts := options.Find().SetSort(bson.D{{params.Field_Search_CreatedAt, -1}})
	filter := bson.D{{}}
	if symbol != "" {
		filter = bson.D{{params.Field_Search_Symbol, bson.D{{params.Equal_Opt, symbol}}}}
	}

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return []models_grpc.ByBitMarketGetInstrumentsInfoLinear{}, err
	}
	if err = cursor.All(ctx, &linears); err != nil {
		return []models_grpc.ByBitMarketGetInstrumentsInfoLinear{}, err
	}
	return linears, nil
}

func (s *ByBitMarketRepository) UpdateInverse(ctx context.Context, collectionName string, items []models_grpc.ByBitMarketGetInstrumentsInfoSpot) error {
	collection := s.Mongodb.MongoConn().Collection(collectionName)
	var err error
	fmt.Println("items len: ", len(items))
	documents := s.GetCountCollecton(ctx, collectionName)

	fmt.Println("documents len: ", documents)
	if documents > 0 {
		for _, item := range items {
			filter := bson.D{{params.Field_Search_Symbol, item.Symbol}}
			update := bson.D{{"$set", item}}

			write, err := collection.UpdateOne(ctx, filter, update)
			if err != nil {
				log.Fatal("UpdateOne error:", err)
			}
			if write.MatchedCount == 0 {
				_, err = collection.InsertOne(ctx, item)
				if err != nil {
					log.Fatal("InsertOne error:", err)
				}
			}
		}
	} else {
		for _, item := range items {
			_, err = collection.InsertOne(ctx, item)
			if err != nil {
				log.Fatal("InsertOne error:", err)
			}
		}
	}
	return err
}
func (s *ByBitMarketRepository) UpdateSpot(ctx context.Context, collectionName string, items []models_grpc.ByBitMarketGetInstrumentsInfoSpot) error {
	collection := s.Mongodb.MongoConn().Collection(collectionName)
	var err error
	fmt.Println("items len: ", len(items))
	documents := s.GetCountCollecton(ctx, collectionName)

	fmt.Println("documents len: ", documents)
	if documents > 0 {
		for _, item := range items {
			filter := bson.D{{params.Field_Search_Symbol, item.Symbol}}
			update := bson.D{{"$set", item}}

			write, err := collection.UpdateOne(ctx, filter, update)
			if err != nil {
				log.Fatal("UpdateOne error:", err)
			}
			if write.MatchedCount == 0 {
				_, err = collection.InsertOne(ctx, item)
				if err != nil {
					log.Fatal("InsertOne error:", err)
				}
			}
		}
	} else {
		for _, item := range items {
			_, err = collection.InsertOne(ctx, item)
			if err != nil {
				log.Fatal("InsertOne error:", err)
			}
		}
	}
	return err
}
func (s *ByBitMarketRepository) UpdateLinear(ctx context.Context, collectionName string, items []models_grpc.ByBitMarketGetInstrumentsInfoLinear) error {
	collection := s.Mongodb.MongoConn().Collection(collectionName)
	var err error
	fmt.Println("items len: ", len(items))
	filter := bson.D{{}}
	documents, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Fatal("CountDocuments error:", err)
	}
	fmt.Println("documents len: ", documents)
	if documents > 0 {
		for _, item := range items {
			filter := bson.D{{params.Field_Search_Symbol, item.Symbol}}
			update := bson.D{{"$set", item}}

			write, err := collection.UpdateOne(ctx, filter, update)
			if err != nil {
				log.Fatal("UpdateOne error:", err)
			}
			if write.MatchedCount == 0 {
				_, err = collection.InsertOne(ctx, item)
				if err != nil {
					log.Fatal("InsertOne error:", err)
				}
			}
		}
	} else {
		for _, item := range items {
			_, err = collection.InsertOne(ctx, item)
			if err != nil {
				log.Fatal("InsertOne error:", err)
			}
		}
	}
	return err
}
func (s *ByBitMarketRepository) FindLastItemBySymbol(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error) {
	var linears []models_grpc.ByBitMarketGetInstrumentsInfoLinear
	collection := s.Mongodb.MongoConn().Collection(collectionName)

	opts := options.Find().SetSort(bson.D{{params.Field_Search_CreatedAt, -1}})

	filter := bson.D{{params.Field_Search_Symbol, bson.D{{params.Equal_Opt, symbol}}}}
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return []models_grpc.ByBitMarketGetInstrumentsInfoLinear{}, err
	}
	if err = cursor.All(ctx, &linears); err != nil {
		return []models_grpc.ByBitMarketGetInstrumentsInfoLinear{}, err
	}
	return linears, nil

}
func (s *ByBitMarketRepository) FindBySymbol(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error) {
	var linears []models_grpc.ByBitMarketGetInstrumentsInfoLinear
	collection := s.Mongodb.MongoConn().Collection(collectionName)

	filter := bson.D{{params.Field_Search_Symbol, bson.D{{params.Equal_Opt, symbol}}}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return []models_grpc.ByBitMarketGetInstrumentsInfoLinear{}, err
	}
	if err = cursor.All(ctx, &linears); err != nil {
		return []models_grpc.ByBitMarketGetInstrumentsInfoLinear{}, err
	}
	return linears, nil
}
