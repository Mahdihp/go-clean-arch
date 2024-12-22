package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	models_grpc "go-clean-arch/internal/bybit_grpc_service/models"
	"go-clean-arch/params"
	"go-clean-arch/util"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"strings"
)

func (s MarketRepository) FindAllRiskLimit(ctx context.Context, category string, symbol []string) (models_grpc.GetRiskLimitLinearDto, error) {
	var spots models_grpc.GetRiskLimitLinearDto

	var key string
	pipeliner := s.Redisdb.Client().Pipeline()
	if len(symbol) > 0 {
		for _, item := range symbol {
			key = params.Market_RiskLimit + ":" + category + ":" + item
			pipeliner.Get(ctx, key)
		}
		var err error
		spots, err := retrivePiplineRiskLimitArray(ctx, pipeliner)
		if err != nil {
			return spots, err
		}
		return spots, err
	} else {
		key := params.Market_RiskLimit + ":" + params.Market_All + ":" + category
		keys := s.Redisdb.Client().Keys(ctx, key)
		if keys != nil {
			for _, key := range keys.Val() {
				pipeliner.Get(ctx, key)
			}
			var err error
			spots, err := retrivePiplineRiskLimiAll(ctx, pipeliner)
			if err != nil {
				return spots, err
			}
			return spots, err
		}
	}
	return spots, nil
}

func retrivePiplineRiskLimiAll(ctx context.Context, pipeliner redis.Pipeliner) (models_grpc.GetRiskLimitLinearDto, error) {
	var items []models_grpc.BybitMarketGetRiskLimit
	var main models_grpc.GetRiskLimitLinearDto
	exec, err := pipeliner.Exec(ctx)
	if err != nil {
		return main, err
	}
	for _, cmder := range exec {
		split := strings.Split(cmder.String(), " ")
		err := json.Unmarshal([]byte(split[2]), &items)
		if err != nil {
			return main, err
		}
	}
	main.Result.List = append(main.Result.List, items...)
	return main, nil
}
func retrivePiplineRiskLimitArray(ctx context.Context, pipeliner redis.Pipeliner) (models_grpc.GetRiskLimitLinearDto, error) {
	var items []models_grpc.BybitMarketGetRiskLimit
	var main models_grpc.GetRiskLimitLinearDto

	exec, err := pipeliner.Exec(ctx)
	if err != nil {
		return main, err
	}
	for _, cmder := range exec {
		split := strings.Split(cmder.String(), " ")
		var ee models_grpc.BybitMarketGetRiskLimit
		err := json.Unmarshal([]byte(split[2]), &ee)
		if err != nil {
			return main, err
		}
		items = append(items, ee)
	}
	main.Result.List = items
	return main, nil
}

func (s MarketRepository) UpdateRiskLimitRedisAll(ctx context.Context, keyGroup string, json interface{}) {
	key := params.Market_RiskLimit + ":" + params.Market_All + ":" + keyGroup
	//json := util.StructToJson(items)
	_, err := s.Redisdb.Client().Set(ctx, key, json, 0).Result()
	if err != nil {
		fmt.Println(err)
	}
}
func (s MarketRepository) UpdateRiskLimitRedis(ctx context.Context, keyGroup string, items []models_grpc.BybitMarketGetRiskLimit) error {
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
func (s MarketRepository) UpdateRiskLimit(ctx context.Context, collectionName string, items []models_grpc.BybitMarketGetRiskLimit) error {
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

func (s MarketRepository) FindAllRiskLimitPagination_MongoDB(ctx context.Context, collectionName string, category string, symbol string, pageIndex int, pageSize int) ([]models_grpc.BybitMarketGetRiskLimit, error) {
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
