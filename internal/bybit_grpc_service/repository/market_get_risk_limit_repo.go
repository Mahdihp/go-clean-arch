package repository

import (
	"context"
	"fmt"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
	"github.com/bxcodec/go-clean-arch/params"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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
func (s ByBitMarketRepository) FindAllRiskLimitPagination(ctx context.Context, collectionName string, category string, symbol string, pageIndex int, pageSize int) ([]models_grpc.BybitMarketGetRiskLimit, error) {
	collection := s.Mongodb.MongoConn().Collection(collectionName)
	var spots []models_grpc.BybitMarketGetRiskLimit
	var cursor *mongo.Cursor
	var err error
	filter := bson.D{{}}

	if len(symbol) <= 0 {
		filter = bson.D{{params.Field_Search_Category, bson.D{{params.Equal_Opt, category}}}}
	} else {
		filter = bson.D{
			{params.And_Opt,
				bson.A{
					bson.D{{params.Field_Search_Category, bson.D{{params.Equal_Opt, category}}}},
					bson.D{{params.Field_Search_Symbol, bson.D{{params.Equal_Opt, symbol}}}},
				},
			},
		}
	}
	if pageSize > 0 && pageIndex >= 0 {
		findOptions := options.Find().SetLimit(int64(pageSize)).SetSkip(int64(pageIndex))
		cursor, err = collection.Find(ctx, filter, findOptions)
	} else {
		cursor, err = collection.Find(ctx, filter)
	}

	if err != nil {
		return []models_grpc.BybitMarketGetRiskLimit{}, err
	}
	if err = cursor.All(ctx, &spots); err != nil {
		return []models_grpc.BybitMarketGetRiskLimit{}, err
	}
	return spots, nil
}
