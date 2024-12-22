package jobs

import (
	"context"
	bybit "github.com/wuhewuhe/bybit.go.api"
	"go-clean-arch/config"
	"log"
	"time"
)

type TradeHistoryActivity struct {
	IsComplate  bool
	Config      config.Config
	byBitClient *bybit.Client
}

func NewTradeHistoryActivity(cfg config.Config) TradeHistoryActivity {
	return TradeHistoryActivity{
		IsComplate:  false,
		Config:      cfg,
		byBitClient: bybit.NewBybitHttpClient(cfg.ByBitWs.ApiKey, cfg.ByBitWs.ApiSecret, bybit.WithBaseURL(bybit.MAINNET)),
	}
}

func (s *TradeHistoryActivity) Work2(ctx context.Context) {
	log.Println(s.IsComplate)
	if s.IsComplate == false {
		log.Printf("Work2 is Complate.%s \n", s.IsComplate)
		s.IsComplate = true

		params := map[string]interface{}{"category": "in.Category", "symbol": ""}
		_, err := s.byBitClient.NewClassicalBybitServiceWithParams(params).GetOrderHistory(ctx)
		if err != nil {
			return
		}
		time.Sleep(5 * time.Second)
		s.IsComplate = false
	}
	//time.Sleep(2 * time.Second)
	//return fmt.Sprint("Work2 is Complate.", str), nil
}
