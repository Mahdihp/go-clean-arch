package bybit_grpc_service

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/db/mongodb"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
	"github.com/bxcodec/go-clean-arch/params"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
)

//	type ByBitMarketRepository interface {
//		FindBySymbol(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error)
//		FindLastItemBySymbol(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error)
//		FindOneAndUpdateLinear(ctx context.Context, collectionName string, items []models_grpc.ByBitMarketGetInstrumentsInfoLinear) error
//		FindAll(ctx context.Context, collectionName string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error)
//	}
type ByBitMarketService struct {
	mongodb *mongodb.MongoDb
}

func NewByBitMarketService(mongodb *mongodb.MongoDb) ByBitMarketService {
	return ByBitMarketService{
		mongodb: mongodb,
	}
}

func (s *ByBitMarketService) FindAll(ctx context.Context, collectionName string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error) {
	collection := s.mongodb.MongoConn().Collection(collectionName)
	var linears []models_grpc.ByBitMarketGetInstrumentsInfoLinear

	opts := options.Find().SetSort(bson.D{{params.Field_Search_CreatedAt, -1}})
	filter := bson.D{{}}

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return []models_grpc.ByBitMarketGetInstrumentsInfoLinear{}, err
	}
	if err = cursor.All(context.TODO(), &linears); err != nil {
		panic(err)
	}
	return linears, nil
}
func (s *ByBitMarketService) FindOneAndUpdateInverse(ctx context.Context, collectionName string, items []models_grpc.ByBitMarketGetInstrumentsInfoSpot) error {

}
func (s *ByBitMarketService) FindOneAndUpdateSpot(ctx context.Context, collectionName string, items []models_grpc.ByBitMarketGetInstrumentsInfoSpot) error {
	collection := s.mongodb.MongoConn().Collection(collectionName)
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
				_, err = collection.InsertOne(context.TODO(), item)
				if err != nil {
					log.Fatal("InsertOne error:", err)
				}
			}
		}
	} else {
		for _, item := range items {
			_, err = collection.InsertOne(context.TODO(), item)
			if err != nil {
				log.Fatal("InsertOne error:", err)
			}
		}
	}
	return err
}
func (s *ByBitMarketService) FindOneAndUpdateLinear(ctx context.Context, collectionName string, items []models_grpc.ByBitMarketGetInstrumentsInfoLinear) error {
	collection := s.mongodb.MongoConn().Collection(collectionName)
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
				_, err = collection.InsertOne(context.TODO(), item)
				if err != nil {
					log.Fatal("InsertOne error:", err)
				}
			}
		}
	} else {
		for _, item := range items {
			_, err = collection.InsertOne(context.TODO(), item)
			if err != nil {
				log.Fatal("InsertOne error:", err)
			}
		}
	}
	return err
}
func (s *ByBitMarketService) FindLastItemBySymbol(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error) {
	var linears []models_grpc.ByBitMarketGetInstrumentsInfoLinear
	collection := s.mongodb.MongoConn().Collection(collectionName)

	opts := options.Find().SetSort(bson.D{{params.Field_Search_CreatedAt, -1}})

	filter := bson.D{{params.Field_Search_Symbol, bson.D{{params.Equal_Opt, symbol}}}}
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return []models_grpc.ByBitMarketGetInstrumentsInfoLinear{}, err
	}
	if err = cursor.All(context.TODO(), &linears); err != nil {
		panic(err)
	}
	return linears, nil

}
func (s *ByBitMarketService) FindBySymbol(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error) {
	var linears []models_grpc.ByBitMarketGetInstrumentsInfoLinear
	collection := s.mongodb.MongoConn().Collection(collectionName)

	filter := bson.D{{params.Field_Search_Symbol, bson.D{{params.Equal_Opt, symbol}}}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return []models_grpc.ByBitMarketGetInstrumentsInfoLinear{}, err
	}
	if err = cursor.All(context.TODO(), &linears); err != nil {
		panic(err)
	}
	return linears, nil
}
