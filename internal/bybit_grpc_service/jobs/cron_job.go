package jobs

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/db/mongodb"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
	params_bybit_http "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/params"
	params2 "github.com/bxcodec/go-clean-arch/params"
	"github.com/go-co-op/gocron/v2"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

func sendRequestToBybit(cfg config.Config, mongodb *mongodb.MongoDb, category string) *bybit.ServerResponse {
	byBitClient := bybit.NewBybitHttpClient(cfg.ByBitWs.ApiKey, cfg.ByBitWs.ApiSecret, bybit.WithBaseURL(bybit.MAINNET))
	params := map[string]interface{}{"category": category, "limit": 1000}
	spotList, err2 := byBitClient.NewClassicalBybitServiceWithParams(params).GetInstrumentInfo(context.Background())
	if err2 != nil {
		fmt.Println("sendRequestToBybit is error:", err2)
		return nil
	}
	return spotList
}
func UpdateInstrumentInfoInverse(cfg config.Config, mongodb *mongodb.MongoDb) gocron.Task {
	//return gocron.NewTask(
	//	func() {

	toBybit := sendRequestToBybit(cfg, mongodb, params2.Market_Inverse)
	if toBybit == nil {
		instrumentInfoDto := params_bybit_http.ToGetInstrumentInfoSpotDto(toBybit)
		coll_InstrumentsInfoSpot := params_bybit_http.ToByBitMarketGetInstrumentsInfoSpot(instrumentInfoDto)
		service := bybit_grpc_service.NewByBitMarketService(mongodb)
		err := service.FindOneAndUpdateSpot(context.Background(), models_grpc.Coll_ByBitMarketGetInstrumentsInfoInverse,
			coll_InstrumentsInfoSpot)
		if err != nil {
			fmt.Println("UpdateInstrumentInfoInverse error...", err)
		}
		fmt.Println("UpdateInstrumentInfoInverse is Complate")
	} else {
		fmt.Println("Bybit Not Connect.")
	}
	//},
	//)
	return nil
}
func UpdateInstrumentInfoSpot(cfg config.Config, mongodb *mongodb.MongoDb) gocron.Task {
	//return gocron.NewTask(
	//	func() {
	toBybit := sendRequestToBybit(cfg, mongodb, params2.Market_Spot)
	if toBybit == nil {
		instrumentInfoDto := params_bybit_http.ToGetInstrumentInfoSpotDto(toBybit)
		coll_InstrumentsInfoSpot := params_bybit_http.ToByBitMarketGetInstrumentsInfoSpot(instrumentInfoDto)
		service := bybit_grpc_service.NewByBitMarketService(mongodb)
		err := service.FindOneAndUpdateSpot(context.Background(), models_grpc.Coll_ByBitMarketGetInstrumentsInfoSpot,
			coll_InstrumentsInfoSpot)
		if err != nil {
			fmt.Println("UpdateInstrumentInfoSpot error...", err)
		}
		fmt.Println("UpdateInstrumentInfoSpot is Complate")
	} else {
		fmt.Println("Bybit Not Connect.")
	}
	//},
	//)
	return nil
}
func UpdateInstrumentInfoLinear(cfg config.Config, mongodb *mongodb.MongoDb) gocron.Task {
	return gocron.NewTask(
		func() {
			toBybit := sendRequestToBybit(cfg, mongodb, params2.Market_Linear)
			if toBybit == nil {
				instrumentInfoDto := params_bybit_http.ToGetInstrumentInfoLinearDto(toBybit)
				coll_InstrumentsInfoLinear := params_bybit_http.ToByBitMarketGetInstrumentsInfoLinear(instrumentInfoDto)
				service := bybit_grpc_service.NewByBitMarketService(mongodb)
				err := service.FindOneAndUpdateLinear(context.Background(), models_grpc.Coll_ByBitMarketGetInstrumentsInfoLinear,
					coll_InstrumentsInfoLinear)
				if err != nil {
					fmt.Println("UpdateInstrumentInfoLinear error...", err)
				}
				fmt.Println("UpdateInstrumentInfoLinear is Complate")
			} else {
				fmt.Println("Bybit Not Connect.")
			}
		},
	)
}
