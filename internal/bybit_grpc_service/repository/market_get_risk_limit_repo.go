package repository

import (
	"context"
	"fmt"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
	"github.com/bxcodec/go-clean-arch/params"
	"go.mongodb.org/mongo-driver/v2/bson"
	"log"
)

func (s ByBitMarketRepository) UpdateRiskLimit(ctx context.Context, collectionName string, items []models_grpc.BybitMarketGetRiskLimit) error {
	collection := s.Mongodb.MongoConn().Collection(collectionName)
	var err error
	fmt.Println("collectionName: ", len(items), collectionName)
	documents := s.GetCountCollecton(ctx, collectionName)
	if documents > 0 {
		for _, item := range items {
			//filter := bson.D{
			//	{params.And_Opt,
			//		bson.A{
			//			bson.D{{params.Field_Search_Symbol, bson.D{{params.Equal_Opt, item.Symbol}}}},
			//			bson.D{{params.Field_Search_Category, bson.D{{params.Equal_Opt, item.Category}}}},
			//		},
			//	},
			//}
			filter := bson.D{{params.Field_Search_Symbol, item.Symbol},
				{params.Field_Search_Category, item.Category}}

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
func (s ByBitMarketRepository) FindAllRiskLimit(ctx context.Context, collectionName string, category string, dto []models_grpc.BybitMarketGetRiskLimit) ([]models_grpc.BybitMarketGetRiskLimit, error) {
	collection := s.Mongodb.MongoConn().Collection(collectionName)
	var spots []models_grpc.BybitMarketGetRiskLimit

	filter := bson.D{{params.Field_Search_Category, bson.D{{params.Equal_Opt, category}}}}
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return []models_grpc.BybitMarketGetRiskLimit{}, err
	}
	if err = cursor.All(context.TODO(), &spots); err != nil {
		panic(err)
	}
	return spots, nil
}
