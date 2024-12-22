package jobs

import (
	"context"
	"fmt"
	_ "github.com/go-co-op/gocron/v2"
	bybit "github.com/wuhewuhe/bybit.go.api"
	"go-clean-arch/config"
	params_http "go-clean-arch/internal/bybit_grpc_service/params"
	"go-clean-arch/internal/bybit_grpc_service/repository"
	"go-clean-arch/params"
	"go-clean-arch/util"
)

// send request to bybit api server from category params
func sendRequestToBybit_GetInstrumentInfo(cfg config.Config, category string) *bybit.ServerResponse {

	byBitClient := bybit.NewBybitHttpClient(cfg.ByBitWs.ApiKey, cfg.ByBitWs.ApiSecret, bybit.WithBaseURL(bybit.MAINNET))
	params := map[string]interface{}{"category": category, "limit": cfg.MaxCountMarket}
	spotList, err2 := byBitClient.NewClassicalBybitServiceWithParams(params).GetInstrumentInfo(context.Background())
	if err2 != nil {
		fmt.Println("sendRequestTo Bybit GetInstrumentInfo is error:", err2)
		return nil
	}
	return spotList
}

// update & add instrument info data with inverse category
func UpdateInstrumentInfoInverse(cfg config.Config, marketSvc repository.MarketRepository) {
	toBybit := sendRequestToBybit_GetInstrumentInfo(cfg, params.Market_Inverse)
	if toBybit != nil {
		instrumentInfoDto := params_http.ToGetInstrumentInfoInverseDto(toBybit)
		err := marketSvc.UpdateRedisInverse(context.Background(), params.Market_Inverse, instrumentInfoDto.Result.List)
		if err != nil {
			fmt.Println("Update InstrumentInfo Inverse error...", err)
		}
		fmt.Println("Update InstrumentInfo Inverse is Complate")
		all := util.StructToJson(instrumentInfoDto.Result.List)
		marketSvc.UpdateRedisAll(context.Background(), params.Market_Inverse, all)
	} else {
		fmt.Println("Bybit Not Connect.")
	}
}

// update & add instrument info data with spot category
func UpdateInstrumentInfoSpot(cfg config.Config, marketSvc repository.MarketRepository) {
	toBybit := sendRequestToBybit_GetInstrumentInfo(cfg, params.Market_Spot)
	if toBybit != nil {
		instrumentInfoDto := params_http.ToGetInstrumentInfoSpotDto(toBybit)
		err := marketSvc.UpdateRedisSpot(context.Background(), params.Market_Spot, instrumentInfoDto.Result.List)
		if err != nil {
			fmt.Println("Update InstrumentInfo Spot error...", err)
		}
		fmt.Println("Update InstrumentInfo Spot is Complate")
		all := util.StructToJson(instrumentInfoDto.Result.List)
		marketSvc.UpdateRedisAll(context.Background(), params.Market_Spot, all)
	} else {
		fmt.Println("Bybit Not Connect.")
	}
}

// update & add instrument info data with linear category
func UpdateInstrumentInfoLinear(cfg config.Config, marketSvc repository.MarketRepository) {
	toBybit := sendRequestToBybit_GetInstrumentInfo(cfg, params.Market_Linear)
	if toBybit != nil {
		instrumentInfoDto := params_http.ToGetInstrumentInfoLinearDto(toBybit)
		err := marketSvc.UpdateRedisLinear(context.Background(), params.Market_Linear, instrumentInfoDto.Result.List)
		if err != nil {
			fmt.Println("Update InstrumentInfo Linear error...", err)
		}
		fmt.Println("Update InstrumentInfo Linear is Complate")
		all := util.StructToJson(instrumentInfoDto.Result.List)
		marketSvc.UpdateRedisAll(context.Background(), params.Market_Linear, all)
	} else {
		fmt.Println("Bybit Not Connect.")
	}
}
