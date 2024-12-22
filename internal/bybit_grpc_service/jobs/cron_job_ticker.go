package jobs

import (
	"context"
	"fmt"
	bybit "github.com/wuhewuhe/bybit.go.api"
	"go-clean-arch/config"
	params_http "go-clean-arch/internal/bybit_grpc_service/params"
	"go-clean-arch/internal/bybit_grpc_service/repository"
	"go-clean-arch/params"
)

// send request to bybit api server from category params
func sendRequestToBybit_GetMarketTickers(cfg config.Config, category string) *bybit.ServerResponse {

	byBitClient := bybit.NewBybitHttpClient(cfg.ByBitWs.ApiKey, cfg.ByBitWs.ApiSecret, bybit.WithBaseURL(bybit.MAINNET))
	params := map[string]interface{}{"category": category}
	spotList, err2 := byBitClient.NewClassicalBybitServiceWithParams(params).GetMarketTickers(context.Background())
	if err2 != nil {
		fmt.Println("sendRequestToBybit_GetMarketTickers is error:", err2)
		return nil
	}

	return spotList
}

// save & update ticker with spot category data
func SaveTickerSpot(cfg config.Config, marketSvc repository.MarketRepository) {
	toBybit := sendRequestToBybit_GetMarketTickers(cfg, params.Market_Spot)
	if toBybit != nil {
		instrumentInfoDto := params_http.ToBybitMarketGetTickerSpotDto(toBybit)
		marketSvc.UpdateTickerSpotRedis(context.Background(), instrumentInfoDto.List)
		fmt.Println("Save Ticker Spot is Complate")
	} else {
		fmt.Println("Bybit Not Connect.")
	}
}

// save & update ticker with linear category data
func SaveTickerLinear(cfg config.Config, marketSvc repository.MarketRepository) {
	toBybit := sendRequestToBybit_GetMarketTickers(cfg, params.Market_Linear)
	if toBybit != nil {
		instrumentInfoDto := params_http.ToBybitMarketGetTickerLinearDto(toBybit)
		marketSvc.UpdateTickerLinearRedis(context.Background(), instrumentInfoDto.List)
		fmt.Println("Save Ticker Linear is Complate")
	} else {
		fmt.Println("Bybit Not Connect.")
	}
}

// save & update ticker with inverse category data
func SaveTickerInverse(cfg config.Config, marketSvc repository.MarketRepository) {
	toBybit := sendRequestToBybit_GetMarketTickers(cfg, params.Market_Inverse)
	if toBybit != nil {
		instrumentInfoDto := params_http.ToBybitMarketGetTickerInverseDto(toBybit)
		marketSvc.UpdateTickerInverseRedis(context.Background(), instrumentInfoDto.List)
		fmt.Println("Save Ticker Inverse is Complate")
	} else {
		fmt.Println("Bybit Not Connect.")
	}
}
