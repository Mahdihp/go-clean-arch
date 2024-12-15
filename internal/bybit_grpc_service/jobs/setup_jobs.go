package jobs

import (
	"context"
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/repository"
	"github.com/robfig/cron/v3"
	"strconv"
)

func SetupCronJob(cfg config.Config, svc repository.ByBitMarketRepository) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("setupCronJob. Error:\n", r)
		}
	}()
	c := cron.New(cron.WithSeconds())

	IntervalSecondMarketTicker := "@every 00h00m" + strconv.Itoa(cfg.RedisMarket.IntervalSecondMarketTicker) + "s"

	c.AddFunc(IntervalSecondMarketTicker, func() {
		SaveTickerSpot(cfg, svc)
	})
	c.AddFunc(IntervalSecondMarketTicker, func() {
		SaveTickerLinear(cfg, svc)
	})
	c.AddFunc(IntervalSecondMarketTicker, func() {
		SaveTickerInverse(cfg, svc)
	})
	//------------------------------InstrumentInfo-------------------------------------------------------
	//"@every 00h00m00s"
	DurationBySecond := "@every 00h" + strconv.Itoa(cfg.CronJob.DurationBySecond/60) + "m00s"
	if svc.GetCountCollecton(context.Background(), models_grpc.Collection_ByBit_MGIIL) <= 0 {
		UpdateInstrumentInfoLinear(cfg, svc)
	} else {
		c.AddFunc(DurationBySecond, func() {
			UpdateInstrumentInfoLinear(cfg, svc)
		})
	}
	DurationBySecond = "@every 00h" + strconv.Itoa(cfg.CronJob.DurationBySecond/50) + "m00s"
	if svc.GetCountCollecton(context.Background(), models_grpc.Collection_ByBit_MGIIS) <= 0 {
		UpdateInstrumentInfoSpot(cfg, svc)
	} else {
		c.AddFunc(DurationBySecond, func() {
			UpdateInstrumentInfoSpot(cfg, svc)
		})
	}
	DurationBySecond = "@every 00h" + strconv.Itoa(cfg.CronJob.DurationBySecond/40) + "m00s"
	if svc.GetCountCollecton(context.Background(), models_grpc.Collection_ByBit_MGIII) <= 0 {
		UpdateInstrumentInfoInverse(cfg, svc)
	} else {
		c.AddFunc(DurationBySecond, func() {
			UpdateInstrumentInfoInverse(cfg, svc)
		})
	}
	//------------------------------Risk Limit-------------------------------------------------------
	//DurationBySecond = "@every 00h00m2s"
	//if svc.GetCountCollecton(context.Background(), models_grpc.Collection_ByBit_MGIII) <= 0 {
	//	UpdateRiskLimitLinear(cfg, svc)
	//} else {
	c.AddFunc("@every 00h00m2s", func() {
		UpdateRiskLimitLinear(cfg, svc)
	})
	//}

	// Start the cron scheduler
	c.Start()
	fmt.Println("Cron scheduler initialized")
	// Keep the main program running
	select {}
}
