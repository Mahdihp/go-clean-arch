package bybit_historical_grpc_service

import (
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/order"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/validator"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

type ByBitHisGrpcServer struct {
	order.UnimplementedOrderServiceServer
	Config      config.Config
	byBitClient *bybit.Client
	validator   validator.ByBitValidator
}

func NewByBitHisGrpcServer(cfg config.Config) ByBitHisGrpcServer {
	return ByBitHisGrpcServer{
		Config:      cfg,
		validator:   validator.NewByBitValidator(),
		byBitClient: bybit.NewBybitHttpClient(cfg.ByBit.ApiKey, cfg.ByBit.ApiSecret, bybit.WithBaseURL(bybit.MAINNET)),
	}
}
