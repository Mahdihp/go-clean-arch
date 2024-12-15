package jobs

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	params_http "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/params"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/repository"
	"github.com/bxcodec/go-clean-arch/params"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func sendRequestToBybit_GetRiskLimit(cfg config.Config, category string) *bybit.ServerResponse {

	byBitClient := bybit.NewBybitHttpClient(cfg.ByBitWs.ApiKey, cfg.ByBitWs.ApiSecret, bybit.WithBaseURL(bybit.MAINNET))
	params := map[string]interface{}{"category": category}
	spotList, err2 := byBitClient.NewClassicalBybitServiceWithParams(params).GetMarketRiskLimits(context.Background())
	if err2 != nil {
		fmt.Println("sendRequestToBybit_GetMarketTickers is error:", err2)
		return nil
	}

	return spotList
}

func UpdateRiskLimitLinear(cfg config.Config, marketSvc repository.ByBitMarketRepository) {
	toBybit := sendRequestToBybit_GetMarketTickers(cfg, params.Market_Linear)
	if toBybit != nil {
		riskLimitDto := params_http.ToBybitMarketGetRiskLimitDto(toBybit)
		//risklimitDtoCollection := params_http.ToBybitMarketGetRiskLimitCollection(riskLimitDto, toBybit.Time, params.Market_Linear)
		fmt.Println("%+v", riskLimitDto)
		//marketSvc.UpdateRiskLimit(context.Background(), models_grpc.Collection_ByBit_MGRL, risklimitDtoCollection)
		fmt.Println("Update Risk Limit Linear is Complate")
	} else {
		fmt.Println("Bybit Not Connect.")
	}
}
