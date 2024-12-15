package repository

import (
	"context"
	"encoding/json"
	"fmt"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
	"github.com/bxcodec/go-clean-arch/params"
	"github.com/bxcodec/go-clean-arch/util"
	"strings"
	"time"
)

func (s ByBitMarketRepository) SaveTickerInverse(ctx context.Context, dto []models_grpc.BybitMarketGetTickerInverse) {
	cat := strings.ToLower(params.Market_Inverse)
	json := util.StructToJson(dto)
	_, err := s.Redisdb.Client().Set(ctx, cat, json, 0).Result()
	if err != nil {
		return
	}
}
func (s ByBitMarketRepository) FindTickerInverse(ctx context.Context) ([]models_grpc.BybitMarketGetTickerInverse, error) {
	var spots []models_grpc.BybitMarketGetTickerInverse
	cat := strings.ToLower(params.Market_Inverse)
	fmt.Println(time.Now())
	result, err := s.Redisdb.Client().Get(ctx, cat).Result()
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

func (s ByBitMarketRepository) SaveTickerLinear(ctx context.Context, dto []models_grpc.BybitMarketGetTickerLinear) {
	cat := strings.ToLower(params.Market_Linear)
	json := util.StructToJson(dto)
	_, err := s.Redisdb.Client().Set(ctx, cat, json, 0).Result()
	if err != nil {
		return
	}
}
func (s ByBitMarketRepository) FindTickerLinear(ctx context.Context) ([]models_grpc.BybitMarketGetTickerLinear, error) {
	var spots []models_grpc.BybitMarketGetTickerLinear
	cat := strings.ToLower(params.Market_Linear)
	fmt.Println(time.Now())
	result, err := s.Redisdb.Client().Get(ctx, cat).Result()
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

func (s ByBitMarketRepository) SaveTickerSpot(ctx context.Context, dto []models_grpc.BybitMarketGetTickerSpot) {
	cat := strings.ToLower(params.Market_Spot)
	json := util.StructToJson(dto)
	_, err := s.Redisdb.Client().Set(ctx, cat, json, 0).Result()
	if err != nil {
		return
	}
}
func (s ByBitMarketRepository) FindTickerSpot(ctx context.Context) ([]models_grpc.BybitMarketGetTickerSpot, error) {
	var spots []models_grpc.BybitMarketGetTickerSpot
	cat := strings.ToLower(params.Market_Spot)
	fmt.Println(time.Now())
	result, err := s.Redisdb.Client().Get(ctx, cat).Result()
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
