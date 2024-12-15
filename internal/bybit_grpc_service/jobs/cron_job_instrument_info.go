package jobs

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
	params_http "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/params"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/repository"
	"github.com/bxcodec/go-clean-arch/params"
	"github.com/go-co-op/gocron/v2"
	_ "github.com/go-co-op/gocron/v2"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func sendRequestToBybit_GetInstrumentInfo(cfg config.Config, category string) *bybit.ServerResponse {

	byBitClient := bybit.NewBybitHttpClient(cfg.ByBitWs.ApiKey, cfg.ByBitWs.ApiSecret, bybit.WithBaseURL(bybit.MAINNET))
	params := map[string]interface{}{"category": category, "limit": cfg.MaxCountMarket}
	spotList, err2 := byBitClient.NewClassicalBybitServiceWithParams(params).GetInstrumentInfo(context.Background())
	if err2 != nil {
		fmt.Println("sendRequestToBybit_GetInstrumentInfo is error:", err2)
		return nil
	}
	return spotList
}
func UpdateInstrumentInfoInverse(cfg config.Config, marketSvc repository.ByBitMarketRepository) gocron.Task {
	return gocron.NewTask(
		func() {

			toBybit := sendRequestToBybit_GetInstrumentInfo(cfg, params.Market_Inverse)
			if toBybit != nil {
				instrumentInfoDto := params_http.ToGetInstrumentInfoSpotDto(toBybit)
				coll_InstrumentsInfoSpot := params_http.ToByBitMarketGetInstrumentsInfoSpot(instrumentInfoDto, toBybit.Time)
				err := marketSvc.UpdateSpot(context.Background(), models_grpc.Collection_ByBit_MGIII,
					coll_InstrumentsInfoSpot)
				if err != nil {
					fmt.Println("Update InstrumentInfo Inverse error...", err)
				}
				fmt.Println("Update InstrumentInfo Inverse is Complate")
			} else {
				fmt.Println("Bybit Not Connect.")
			}
		},
	)
}
func UpdateInstrumentInfoSpot(cfg config.Config, marketSvc repository.ByBitMarketRepository) gocron.Task {
	return gocron.NewTask(
		func() {
			toBybit := sendRequestToBybit_GetInstrumentInfo(cfg, params.Market_Spot)
			if toBybit != nil {
				instrumentInfoDto := params_http.ToGetInstrumentInfoSpotDto(toBybit)
				coll_InstrumentsInfoSpot := params_http.ToByBitMarketGetInstrumentsInfoSpot(instrumentInfoDto, toBybit.Time)

				err := marketSvc.UpdateSpot(context.Background(), models_grpc.Collection_ByBit_MGIIS,
					coll_InstrumentsInfoSpot)
				if err != nil {
					fmt.Println("Update InstrumentInfo Spot error...", err)
				}
				fmt.Println("Update InstrumentInfo Spot is Complate")
			} else {
				fmt.Println("Bybit Not Connect.")
			}
		},
	)
}
func UpdateInstrumentInfoLinear(cfg config.Config, marketSvc repository.ByBitMarketRepository) gocron.Task {
	return gocron.NewTask(
		func() {
			toBybit := sendRequestToBybit_GetInstrumentInfo(cfg, params.Market_Linear)
			if toBybit != nil {
				instrumentInfoDto := params_http.ToGetInstrumentInfoLinearDto(toBybit)
				coll_InstrumentsInfoLinear := params_http.ToByBitMarketGetInstrumentsInfoLinear(instrumentInfoDto, toBybit.Time)

				err := marketSvc.UpdateLinear(context.Background(), models_grpc.Collection_ByBit_MGIIL,
					coll_InstrumentsInfoLinear)
				if err != nil {
					fmt.Println("Update InstrumentInfo Linear error...", err)
				}
				fmt.Println("Update InstrumentInfo Linear is Complate")
			} else {
				fmt.Println("Bybit Not Connect.")
			}
		},
	)
}
