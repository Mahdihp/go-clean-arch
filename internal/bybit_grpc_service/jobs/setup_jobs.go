package jobs

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"go-clean-arch/config"
	"go-clean-arch/internal/bybit_grpc_service/repository"
	"strconv"
)

// SetupCronJob setup cron jobs
func SetupCronJob(cfg config.Config, svc repository.MarketRepository) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("setupCronJob. Error:\n", r)
		}
	}()
	//c := cron.New(cron.WithSeconds())
	c := cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)))
	//------------------------------Get Ticker---------------------------------------------
	TickerInterval := "@every 00h00m" + strconv.Itoa(cfg.CronJob.TickerInterval) + "s"

	c.AddFunc(TickerInterval, func() {
		SaveTickerSpot(cfg, svc)
	})
	c.AddFunc(TickerInterval, func() {
		SaveTickerLinear(cfg, svc)
	})
	c.AddFunc(TickerInterval, func() {
		SaveTickerInverse(cfg, svc)
	})
	//------------------------------Instrument Info-------------------------------------------------------
	//"@every 00h00m00s"
	InstrumentInfoInterval := "@every " + strconv.Itoa(cfg.CronJob.InstrumentInfoInterval) + "h00m00s"
	c.AddFunc(InstrumentInfoInterval, func() {
		UpdateInstrumentInfoLinear(cfg, svc)
	})

	InstrumentInfoInterval = "@every " + strconv.Itoa(cfg.CronJob.InstrumentInfoInterval) + "h00m00s"
	c.AddFunc(InstrumentInfoInterval, func() {
		UpdateInstrumentInfoSpot(cfg, svc)
	})

	InstrumentInfoInterval = "@every " + strconv.Itoa(cfg.CronJob.InstrumentInfoInterval) + "h00m00s"
	c.AddFunc(InstrumentInfoInterval, func() {
		UpdateInstrumentInfoInverse(cfg, svc)
	})

	//------------------------------Risk Limit-------------------------------------------------------
	RiskLimitInterval := "@every " + strconv.Itoa(cfg.CronJob.RiskLimitInterval) + "h00m00s"
	c.AddFunc(RiskLimitInterval, func() {
		UpdateRiskLimitLinear(cfg, svc)
	})

	RiskLimitInterval = "@every " + strconv.Itoa(cfg.CronJob.RiskLimitInterval) + "h00m00s"
	c.AddFunc(RiskLimitInterval, func() {
		UpdateRiskLimitInverse(cfg, svc)
	})

	// Start the cron scheduler
	c.Start()
	fmt.Println("Cron Scheduler Initialized")
	// Keep the main program running
	select {}
}
