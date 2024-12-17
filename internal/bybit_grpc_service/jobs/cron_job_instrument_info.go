package jobs

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	params_http "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/params"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/repository"
	"github.com/bxcodec/go-clean-arch/params"
	"github.com/bxcodec/go-clean-arch/util"
	_ "github.com/go-co-op/gocron/v2"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

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
func UpdateInstrumentInfoInverse(cfg config.Config, marketSvc repository.ByBitMarketRepository) {
	toBybit := sendRequestToBybit_GetInstrumentInfo(cfg, params.Market_Inverse)
	if toBybit != nil {
		instrumentInfoDto := params_http.ToGetInstrumentInfoInverseDto(toBybit)
		//coll_InstrumentsInfoSpot := params_http.ToByBitMarketGetInstrumentsInfoInverse(instrumentInfoDto, toBybit.Time)
		//err := marketSvc.UpdateInverse(context.Background(), models_grpc.Collection_ByBit_MGIII, coll_InstrumentsInfoSpot)
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
func UpdateInstrumentInfoSpot(cfg config.Config, marketSvc repository.ByBitMarketRepository) {
	toBybit := sendRequestToBybit_GetInstrumentInfo(cfg, params.Market_Spot)
	if toBybit != nil {
		instrumentInfoDto := params_http.ToGetInstrumentInfoSpotDto(toBybit)
		//coll_InstrumentsInfoSpot := params_http.ToByBitMarketGetInstrumentsInfoSpot(instrumentInfoDto, toBybit.Time)

		//err := marketSvc.UpdateSpot(context.Background(), models_grpc.Collection_ByBit_MGIIS, coll_InstrumentsInfoSpot)
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
func UpdateInstrumentInfoLinear(cfg config.Config, marketSvc repository.ByBitMarketRepository) {
	toBybit := sendRequestToBybit_GetInstrumentInfo(cfg, params.Market_Linear)
	if toBybit != nil {
		instrumentInfoDto := params_http.ToGetInstrumentInfoLinearDto(toBybit)
		//coll_InstrumentsInfoLinear := params_http.ToByBitMarketGetInstrumentsInfoLinear(instrumentInfoDto, toBybit.Time)
		//err := marketSvc.UpdateLinear(context.Background(), models_grpc.Collection_ByBit_MGIIL,coll_InstrumentsInfoLinear)
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
