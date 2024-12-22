package repository

import (
	"context"
	"encoding/json"
	"fmt"
	models_grpc "go-clean-arch/internal/bybit_grpc_service/models"
	"go-clean-arch/params"
	"go-clean-arch/util"
	"strings"
	"time"
)

func (s MarketRepository) UpdateTickerInverseRedis(ctx context.Context, dto []models_grpc.BybitMarketGetTickerInverse) {
	key := params.Market_Ticker + ":" + strings.ToLower(params.Market_Inverse)
	json := util.StructToJson(dto)
	_, err := s.Redisdb.Client().Set(ctx, key, json, 0).Result()
	if err != nil {
		return
	}
}
func (s MarketRepository) UpdateTickerLinearRedis(ctx context.Context, dto []models_grpc.BybitMarketGetTickerLinear) {
	key := params.Market_Ticker + ":" + strings.ToLower(params.Market_Linear)
	json := util.StructToJson(dto)
	_, err := s.Redisdb.Client().Set(ctx, key, json, 0).Result()
	if err != nil {
		return
	}
}
func (s MarketRepository) UpdateTickerSpotRedis(ctx context.Context, dto []models_grpc.BybitMarketGetTickerSpot) {
	key := params.Market_Ticker + ":" + strings.ToLower(params.Market_Spot)
	json := util.StructToJson(dto)
	_, err := s.Redisdb.Client().Set(ctx, key, json, 0).Result()
	if err != nil {
		return
	}
}

func (s MarketRepository) FindTickerInverse(ctx context.Context) ([]models_grpc.BybitMarketGetTickerInverse, error) {
	var spots []models_grpc.BybitMarketGetTickerInverse
	key := params.Market_Ticker + ":" + strings.ToLower(params.Market_Inverse)
	fmt.Println(time.Now())
	result, err := s.Redisdb.Client().Get(ctx, key).Result()
	if err != nil {
		return spots, err
	}
	err = json.Unmarshal([]byte(result), &spots)
	if err != nil {
		return spots, nil
	}
	fmt.Println(time.Now())
	return spots, nil
}
func (s MarketRepository) FindTickerLinear(ctx context.Context) ([]models_grpc.BybitMarketGetTickerLinear, error) {
	var spots []models_grpc.BybitMarketGetTickerLinear
	key := params.Market_Ticker + ":" + strings.ToLower(params.Market_Linear)
	fmt.Println(time.Now())
	result, err := s.Redisdb.Client().Get(ctx, key).Result()
	if err != nil {
		return spots, err
	}
	err = json.Unmarshal([]byte(result), &spots)
	if err != nil {
		return spots, nil
	}
	fmt.Println(time.Now())
	return spots, nil
}
func (s MarketRepository) FindTickerSpot(ctx context.Context) ([]models_grpc.BybitMarketGetTickerSpot, error) {
	var spots []models_grpc.BybitMarketGetTickerSpot
	key := params.Market_Ticker + ":" + strings.ToLower(params.Market_Spot)
	fmt.Println(time.Now())
	result, err := s.Redisdb.Client().Get(ctx, key).Result()
	if err != nil {
		return spots, nil
	}
	err = json.Unmarshal([]byte(result), &spots)
	if err != nil {
		return spots, nil
	}
	fmt.Println(time.Now())
	return spots, err
}
