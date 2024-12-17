package jobs

import (
	"fmt"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/repository"
	"github.com/robfig/cron/v3"
)

func SetupCronJob(cfg config.Config, svc repository.ByBitMarketRepository) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("setupCronJob. Error:\n", r)
		}
	}()
	c := cron.New(cron.WithSeconds())

	//------------------------------Get Ticker---------------------------------------------
	//IntervalSecondMarketTicker := "@every 00h00m" + strconv.Itoa(cfg.RedisMarket.IntervalSecondMarketTicker) + "s"
	//
	//c.AddFunc(IntervalSecondMarketTicker, func() {
	//	SaveTickerSpot(cfg, svc)
	//})
	//c.AddFunc(IntervalSecondMarketTicker, func() {
	//	SaveTickerLinear(cfg, svc)
	//})
	//c.AddFunc(IntervalSecondMarketTicker, func() {
	//	SaveTickerInverse(cfg, svc)
	//})
	//------------------------------Instrument Info-------------------------------------------------------
	//"@every 00h00m00s"
	//DurationBySecond := "@every 00h" + strconv.Itoa(cfg.CronJob.DurationBySecond/60) + "m00s"
	//c.AddFunc(DurationBySecond, func() {
	UpdateInstrumentInfoLinear(cfg, svc)
	//})
	//
	//DurationBySecond = "@every 00h" + strconv.Itoa(cfg.CronJob.DurationBySecond/50) + "m00s"
	//c.AddFunc(DurationBySecond, func() {
	UpdateInstrumentInfoSpot(cfg, svc)
	//})
	//
	//DurationBySecond = "@every 00h" + strconv.Itoa(cfg.CronJob.DurationBySecond/40) + "m00s"
	//c.AddFunc(DurationBySecond, func() {
	UpdateInstrumentInfoInverse(cfg, svc)
	//})
	//------------------------------Risk Limit-------------------------------------------------------
	//DurationBySecond = "@every 00h" + strconv.Itoa(cfg.CronJob.DurationBySecond/40) + "m00s"
	//c.AddFunc(DurationBySecond, func() {
	//	UpdateRiskLimitInverse(cfg, svc)
	//})
	//
	//DurationBySecond = "@every 00h" + strconv.Itoa(cfg.CronJob.DurationBySecond/35) + "m00s"
	//c.AddFunc(DurationBySecond, func() {
	//	UpdateRiskLimitInverse(cfg, svc)
	//})

	// Start the cron scheduler
	c.Start()
	fmt.Println("Cron scheduler initialized")
	// Keep the main program running
	select {}
}
