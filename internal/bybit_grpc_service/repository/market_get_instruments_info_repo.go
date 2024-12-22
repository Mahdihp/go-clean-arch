package repository

import (
	"context"
	"encoding/json"
	"fmt"
	redis2 "github.com/redis/go-redis/v9"
	"go-clean-arch/config"
	"go-clean-arch/db/mongodb"
	"go-clean-arch/db/redis"
	models_grpc "go-clean-arch/internal/bybit_grpc_service/models"
	params_http "go-clean-arch/internal/bybit_grpc_service/params"
	"go-clean-arch/params"
	"go-clean-arch/util"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"log"
	"strings"
)

type MarketRepository struct {
	Mongodb *mongodb.MongoDb
	Redisdb redis.RedisDb
	cfg     config.Config
}

func New(cfg config.Config) MarketRepository {
	return MarketRepository{
		Mongodb: mongodb.NewMongoDb(cfg.MongoDbMarket),
		Redisdb: redis.NewRedis(cfg.RedisMarket),
	}
}
func (receiver MarketRepository) GetCountCollecton(ctx context.Context, collName string) int64 {
	collection := receiver.Mongodb.MongoConn().Collection(collName)
	filter := bson.D{{}}
	documents, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Fatal("CountDocuments error:", err)
		return 0
	}
	return documents
}

func (s *MarketRepository) FindAllInverse_MongoDB(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoInverse, error) {
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
func (s *MarketRepository) FindAllSpot_MongoDB(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoSpot, error) {
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
func (s *MarketRepository) FindAllLinear_MongoDB(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error) {
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
func (s *MarketRepository) UpdateInverse_MongoDB(ctx context.Context, collectionName string, items []models_grpc.ByBitMarketGetInstrumentsInfoSpot) error {
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
func (s *MarketRepository) UpdateLinear_MongoDB(ctx context.Context, collectionName string, items []models_grpc.ByBitMarketGetInstrumentsInfoLinear) error {
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
func (s *MarketRepository) FindLastItemBySymbol_MongoDB(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error) {
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
func (s *MarketRepository) FindBySymbol_MongoDB(ctx context.Context, collectionName string, symbol string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error) {
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

func (s *MarketRepository) UpdateSpot_MongoDB(ctx context.Context, collectionName string, items []models_grpc.ByBitMarketGetInstrumentsInfoSpot) error {
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
func (s *MarketRepository) FindRedisInverse(ctx context.Context, category string, symbol []string) ([]models_grpc.ByBitMarketGetInstrumentsInfoInverse, error) {
	var items []models_grpc.ByBitMarketGetInstrumentsInfoInverse
	var key string
	pipeliner := s.Redisdb.Client().Pipeline()
	if len(symbol) > 0 {
		for _, item := range symbol {
			key = params.Market_InstrumentInfo + ":" + category + ":" + item
			pipeliner.Get(ctx, key)
		}
		var err error
		items, err := retrivePiplineInstrumentInfoInverseArray(ctx, pipeliner)
		if err != nil {
			return items, err
		}
		return items, err
	} else {
		key = params.Market_InstrumentInfo + ":" + params.Market_All + ":" + category
		keys := s.Redisdb.Client().Keys(ctx, key)
		if keys != nil {
			for _, key := range keys.Val() {
				pipeliner.Get(ctx, key)
			}
			var err error
			items, err := retrivePiplineInstrumentInfoInverseAll(ctx, pipeliner)
			if err != nil {
				return items, err
			}
			return items, err
		}
	}
	return items, nil
}
func (s *MarketRepository) FindRedisSpot(ctx context.Context, category string, symbol []string) ([]models_grpc.ByBitMarketGetInstrumentsInfoSpot, error) {
	var items []models_grpc.ByBitMarketGetInstrumentsInfoSpot
	var key string
	pipeliner := s.Redisdb.Client().Pipeline()
	if len(symbol) > 0 {
		for _, item := range symbol {
			key = params.Market_InstrumentInfo + ":" + category + ":" + item
			pipeliner.Get(ctx, key)
		}
		var err error
		items, err := retrivePiplineInstrumentInfoSpotArray(ctx, pipeliner)
		if err != nil {
			return items, err
		}
		return items, err
	} else {
		key = params.Market_InstrumentInfo + ":" + params.Market_All + ":" + category
		keys := s.Redisdb.Client().Keys(ctx, key)
		if keys != nil {
			for _, key := range keys.Val() {
				pipeliner.Get(ctx, key)
			}
			var err error
			items, err := retrivePiplineInstrumentInfoSpotAll(ctx, pipeliner)
			if err != nil {
				return items, err
			}
			return items, err
		}
	}
	return items, nil
}

func (s *MarketRepository) FindRedisLinear(ctx context.Context, category string, symbol []string) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error) {
	var items []models_grpc.ByBitMarketGetInstrumentsInfoLinear
	var key string
	pipeliner := s.Redisdb.Client().Pipeline()
	if len(symbol) > 0 {
		for _, item := range symbol {
			key = params.Market_InstrumentInfo + ":" + category + ":" + item
			pipeliner.Get(ctx, key)
		}
		var err error
		items, err := retrivePiplineInstrumentInfoLinearArray(ctx, pipeliner)
		if err != nil {
			return items, err
		}
		return items, err
	} else {
		key = params.Market_InstrumentInfo + ":" + params.Market_All + ":" + category
		keys := s.Redisdb.Client().Keys(ctx, key)
		if keys != nil {
			for _, key := range keys.Val() {
				pipeliner.Get(ctx, key)
			}
			var err error
			spots, err := retrivePiplineInstrumentInfoLinearAll(ctx, pipeliner)
			if err != nil {
				return spots, err
			}
			return spots, err
		}
	}
	return items, nil
}

func (s *MarketRepository) UpdateRedisInverse(ctx context.Context, keyGroup string, items []models_grpc.ByBitMarketGetInstrumentsInfoInverse) error {
	key := keyGroup
	pipeliner := s.Redisdb.Client().Pipeline()
	for _, item := range items {
		Key := params.Market_InstrumentInfo + ":" + key + ":" + item.Symbol
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
func (s *MarketRepository) UpdateRedisSpot(ctx context.Context, keyGroup string, items []params_http.InstrumentInfoSpotDto) error {
	key := keyGroup
	pipeliner := s.Redisdb.Client().Pipeline()
	for _, item := range items {
		Key := params.Market_InstrumentInfo + ":" + key + ":" + item.Symbol
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
func (s *MarketRepository) UpdateRedisAll(ctx context.Context, keyGroup string, json interface{}) {
	key := params.Market_InstrumentInfo + ":" + params.Market_All + ":" + keyGroup
	//json := util.StructToJson(items)
	_, err := s.Redisdb.Client().Set(ctx, key, json, 0).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func (s *MarketRepository) UpdateRedisLinear(ctx context.Context, keyGroup string, items []params_http.InstrumentInfoLinearDto) error {
	key := keyGroup
	pipeliner := s.Redisdb.Client().Pipeline()
	for _, item := range items {
		Key := params.Market_InstrumentInfo + ":" + key + ":" + item.Symbol
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

func retrivePiplineInstrumentInfoSpotAll(ctx context.Context, pipeliner redis2.Pipeliner) ([]models_grpc.ByBitMarketGetInstrumentsInfoSpot, error) {
	var items []models_grpc.ByBitMarketGetInstrumentsInfoSpot
	//var main models_grpc.ByBitMarketGetInstrumentsInfoSpot
	//main.Result.Category = params.Market_Linear
	exec, err := pipeliner.Exec(ctx)
	if err != nil {
		return items, err
	}
	for _, cmder := range exec {
		split := strings.Split(cmder.String(), " ")
		err := json.Unmarshal([]byte(split[2]), &items)
		if err != nil {
			return items, err
		}
	}
	return items, nil
}
func retrivePiplineInstrumentInfoLinearAll(ctx context.Context, pipeliner redis2.Pipeliner) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error) {
	var items []models_grpc.ByBitMarketGetInstrumentsInfoLinear
	var main models_grpc.GetInstrumentInfoLinearDto
	main.Result.Category = params.Market_Linear
	exec, err := pipeliner.Exec(ctx)
	if err != nil {
		return items, err
	}
	for _, cmder := range exec {
		split := strings.Split(cmder.String(), " ")
		err := json.Unmarshal([]byte(split[2]), &items)
		if err != nil {
			return items, err
		}
	}
	main.Result.List = append(main.Result.List, items...)
	return items, nil
}
func retrivePiplineInstrumentInfoInverseAll(ctx context.Context, pipeliner redis2.Pipeliner) ([]models_grpc.ByBitMarketGetInstrumentsInfoInverse, error) {
	var items []models_grpc.ByBitMarketGetInstrumentsInfoInverse
	exec, err := pipeliner.Exec(ctx)
	if err != nil {
		return items, err
	}
	for _, cmder := range exec {
		split := strings.Split(cmder.String(), " ")
		err := json.Unmarshal([]byte(split[2]), &items)
		if err != nil {
			return items, err
		}
	}
	return items, nil
}
func retrivePiplineInstrumentInfoInverseArray(ctx context.Context, pipeliner redis2.Pipeliner) ([]models_grpc.ByBitMarketGetInstrumentsInfoInverse, error) {
	var items []models_grpc.ByBitMarketGetInstrumentsInfoInverse
	exec, err := pipeliner.Exec(ctx)
	if err != nil {
		return items, err
	}
	for _, cmder := range exec {
		split := strings.Split(cmder.String(), " ")
		var ee models_grpc.ByBitMarketGetInstrumentsInfoInverse
		err := json.Unmarshal([]byte(split[2]), &ee)
		if err != nil {
			return items, err
		}
		items = append(items, ee)
	}
	return items, nil
}
func retrivePiplineInstrumentInfoSpotArray(ctx context.Context, pipeliner redis2.Pipeliner) ([]models_grpc.ByBitMarketGetInstrumentsInfoSpot, error) {
	var items []models_grpc.ByBitMarketGetInstrumentsInfoSpot
	exec, err := pipeliner.Exec(ctx)
	if err != nil {
		return items, err
	}
	for _, cmder := range exec {
		split := strings.Split(cmder.String(), " ")
		var ee models_grpc.ByBitMarketGetInstrumentsInfoSpot
		err := json.Unmarshal([]byte(split[2]), &ee)
		if err != nil {
			return items, err
		}
		items = append(items, ee)
	}
	return items, nil
}
func retrivePiplineInstrumentInfoLinearArray(ctx context.Context, pipeliner redis2.Pipeliner) ([]models_grpc.ByBitMarketGetInstrumentsInfoLinear, error) {
	var items []models_grpc.ByBitMarketGetInstrumentsInfoLinear
	exec, err := pipeliner.Exec(ctx)
	if err != nil {
		return items, err
	}
	for _, cmder := range exec {
		split := strings.Split(cmder.String(), " ")
		var ee models_grpc.ByBitMarketGetInstrumentsInfoLinear
		err := json.Unmarshal([]byte(split[2]), &ee)
		if err != nil {
			return items, err
		}
		items = append(items, ee)
	}
	return items, nil
}
