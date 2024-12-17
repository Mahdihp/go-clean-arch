package jobs

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	params_http "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/params"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/repository"
	"github.com/bxcodec/go-clean-arch/params"
	"github.com/bxcodec/go-clean-arch/util"
	bybit "github.com/wuhewuhe/bybit.go.api"
	"time"
)

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

func UpdateRiskLimitLinear(cfg config.Config, marketSvc repository.ByBitMarketRepository) {
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
			//err := marketSvc.UpdateRiskLimit(context.Background(), models_grpc.Collection_ByBit_MGRL, risklimitDtoCollection)
			err := marketSvc.UpdateRiskLimitRedis(context.Background(), params.Market_Linear, risklimitDtoCollection)
			if err != nil {
				fmt.Println("Update Risk Limit Linear  error...", err)
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
	fmt.Printf("Update Risk Limit Linear is Complate.\n Retry Count %d\n", retryCount)
}

func UpdateRiskLimitInverse(cfg config.Config, marketSvc repository.ByBitMarketRepository) {
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
			//err := marketSvc.UpdateRiskLimit(context.Background(), models_grpc.Collection_ByBit_MGRL, risklimitDtoCollection)
			err := marketSvc.UpdateRiskLimitRedis(context.Background(), params.Market_Inverse, risklimitDtoCollection)
			if err != nil {
				fmt.Println("Update Risk Limit Inverse  error...", err)
			}
			all := util.StructToJson(risklimitDtoCollection)
			marketSvc.UpdateRiskLimitRedisAll(context.Background(), params.Market_Inverse, all)

		} else {
			retryCount++
			if retryCount == 10 {
				isFinal = false
			}
			time.Sleep(10 * time.Second)
			fmt.Println("Bybit Not Connect.")
		}
	}
	fmt.Printf("Update Risk Limit Inverse is Complate.\n Retry Count %d\n", retryCount)
}
