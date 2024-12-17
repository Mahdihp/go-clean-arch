package repository

import (
	"context"
	"encoding/json"
	"fmt"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
	"github.com/bxcodec/go-clean-arch/params"
	"github.com/bxcodec/go-clean-arch/util"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"strings"
)

func (s ByBitMarketRepository) FindAllRiskLimit(ctx context.Context, category string, symbol string) ([]models_grpc.BybitMarketGetRiskLimit, error) {
	var spots []models_grpc.BybitMarketGetRiskLimit
	var key string
	if len(category) > 0 && len(symbol) > 0 {
		key = params.Market_RiskLimit + ":" + category + ":" + symbol
	}
	if len(category) > 0 && len(symbol) <= 0 {
		key = params.Market_RiskLimit + ":" + category + "*"
	}

	keys := s.Redisdb.Client().Keys(ctx, key)
	if keys != nil {
		pipeliner := s.Redisdb.Client().Pipeline()
		for _, key := range keys.Val() {
			pipeliner.Get(ctx, key)
		}
		exec, err := pipeliner.Exec(ctx)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		for _, cmder := range exec {
			split := strings.Split(cmder.String(), " ")
			var ee models_grpc.BybitMarketGetRiskLimit
			err := json.Unmarshal([]byte(split[2]), &ee)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			spots = append(spots, ee)
		}
	}

	return spots, nil
}
func (s ByBitMarketRepository) UpdateRiskLimitRedisAll(ctx context.Context, keyGroup string, json interface{}) error {
	key := params.Market_RiskLimit + ":" + params.Market_All + ":" + keyGroup
	//json := util.StructToJson(items)
	_, err := s.Redisdb.Client().Set(ctx, key, json, 0).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (s ByBitMarketRepository) UpdateRiskLimitRedis(ctx context.Context, keyGroup string, items []models_grpc.BybitMarketGetRiskLimit) error {
	key := keyGroup
	pipeliner := s.Redisdb.Client().Pipeline()
	for _, item := range items {
		Key := params.Market_RiskLimit + ":" + key + ":" + item.Symbol
		json := util.StructToJson(item)
		pipeliner.Set(ctx, Key, json, 0)
	}
	_, err := pipeliner.Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
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
