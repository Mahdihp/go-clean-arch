package jobs

import (
	"context"
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
	"go-clean-arch/config"
	models_grpc "go-clean-arch/internal/bybit_grpc_service/models"
	params_http "go-clean-arch/internal/bybit_grpc_service/params"
	"go-clean-arch/internal/bybit_grpc_service/repository"
	"go-clean-arch/params"
	"go-clean-arch/util"
	"time"
)

// send request to bybit api server from category params
func sendRequestToBybit_GetRiskLimit(cfg config.Config, category string, cursor string) *bybit.ServerResponse {
	byBitClient := bybit.NewBybitHttpClient(cfg.ByBitWs.ApiKey, cfg.ByBitWs.ApiSecret, bybit.WithBaseURL(bybit.MAINNET))
	params := map[string]interface{}{"category": category, "cursor": cursor}
	spotList, err2 := byBitClient.NewClassicalBybitServiceWithParams(params).GetMarketRiskLimits(context.Background())
	if err2 != nil {
		fmt.Println("sendRequestToBybit_GetMarketTickers is error:", err2)
		return nil
	}

	return spotList
}

// save & update risk limit data with linear category
func UpdateRiskLimitLinear(cfg config.Config, marketSvc repository.MarketRepository) {
	var allData []models_grpc.BybitMarketGetRiskLimit
	var isFinal bool = true
	var retryCount int = 1
	var cursor string
	for isFinal {
		toBybit := sendRequestToBybit_GetRiskLimit(cfg, params.Market_Linear, cursor)
		if toBybit != nil {
			riskLimitDto := params_http.ToBybitMarketGetRiskLimitDto(toBybit)
			if len(riskLimitDto.Result.NextPageCursor) > 0 {
				cursor = riskLimitDto.Result.NextPageCursor
			} else {
				isFinal = false
			}
			fmt.Println("Update Risk Limit Linear  Cursor...", cursor)
			risklimitDtoCollection := params_http.ToBybitMarketGetRiskLimitCollection(riskLimitDto, toBybit.Time, params.Market_Linear)
			err := marketSvc.UpdateRiskLimitRedis(context.Background(), params.Market_Linear, risklimitDtoCollection)
			if err != nil {
				fmt.Println("Update Risk Limit Linear  error...", err)
			}
			if len(risklimitDtoCollection) > 0 {
				allData = append(allData, risklimitDtoCollection...)
			}

		} else {
			retryCount++
			if retryCount == 10 {
				isFinal = false
			}
			time.Sleep(10 * time.Second)
			fmt.Println("Bybit Not Connect.")
		}
	}
	all := util.StructToJson(allData)
	marketSvc.UpdateRiskLimitRedisAll(context.Background(), params.Market_Linear, all)
	fmt.Printf("Update Risk Limit Linear is Complate.\n Retry Count %d\n", retryCount)
}

// save & update risk limit data with inverse category
func UpdateRiskLimitInverse(cfg config.Config, marketSvc repository.MarketRepository) {
	var allData []models_grpc.BybitMarketGetRiskLimit
	var isFinal bool = true
	var retryCount int = 1
	var cursor string
	for isFinal {
		toBybit := sendRequestToBybit_GetRiskLimit(cfg, params.Market_Inverse, "")
		if toBybit != nil {
			riskLimitDto := params_http.ToBybitMarketGetRiskLimitDto(toBybit)
			if len(riskLimitDto.Result.NextPageCursor) > 0 {
				cursor = riskLimitDto.Result.NextPageCursor
			} else {
				isFinal = false
			}
			fmt.Println("Update Risk Limit Inverse  Cursor...", cursor)
			risklimitDtoCollection := params_http.ToBybitMarketGetRiskLimitCollection(riskLimitDto, toBybit.Time, params.Market_Inverse)
			err := marketSvc.UpdateRiskLimitRedis(context.Background(), params.Market_Inverse, risklimitDtoCollection)
			if err != nil {
				fmt.Println("Update Risk Limit Inverse  error...", err)
			}
			if len(risklimitDtoCollection) > 0 {
				allData = append(allData, risklimitDtoCollection...)
			}
		} else {
			retryCount++
			if retryCount == 10 {
				isFinal = false
			}
			time.Sleep(10 * time.Second)
			fmt.Println("Bybit Not Connect.")
		}
	}
	all := util.StructToJson(allData)
	marketSvc.UpdateRiskLimitRedisAll(context.Background(), params.Market_Inverse, all)
	fmt.Printf("Update Risk Limit Inverse is Complate.\n Retry Count %d\n", retryCount)
}
