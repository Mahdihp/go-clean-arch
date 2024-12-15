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

func SaveTickerSpot(cfg config.Config, marketSvc repository.ByBitMarketRepository) {
	toBybit := sendRequestToBybit_GetMarketTickers(cfg, params.Market_Spot)
	if toBybit != nil {
		instrumentInfoDto := params_http.ToBybitMarketGetTickerSpotDto(toBybit)
		marketSvc.SaveTickerSpot(context.Background(), instrumentInfoDto.List)
		fmt.Println("Save Ticker Spot is Complate")
	} else {
		fmt.Println("Bybit Not Connect.")
	}
}

func SaveTickerLinear(cfg config.Config, marketSvc repository.ByBitMarketRepository) {
	toBybit := sendRequestToBybit_GetMarketTickers(cfg, params.Market_Linear)
	if toBybit != nil {
		instrumentInfoDto := params_http.ToBybitMarketGetTickerLinearDto(toBybit)
		marketSvc.SaveTickerLinear(context.Background(), instrumentInfoDto.List)
		fmt.Println("Save Ticker Linear is Complate")
	} else {
		fmt.Println("Bybit Not Connect.")
	}
}

func SaveTickerInverse(cfg config.Config, marketSvc repository.ByBitMarketRepository) {
	toBybit := sendRequestToBybit_GetMarketTickers(cfg, params.Market_Inverse)
	if toBybit != nil {
		instrumentInfoDto := params_http.ToBybitMarketGetTickerInverseDto(toBybit)
		marketSvc.SaveTickerInverse(context.Background(), instrumentInfoDto.List)
		fmt.Println("Save Ticker Inverse is Complate")
	} else {
		fmt.Println("Bybit Not Connect.")
	}
}
