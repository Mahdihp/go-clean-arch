package bybit_http_service

import (
	"context"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/order"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/validator"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

type ByBitHttpServerOrder struct {
	order.UnimplementedOrderServiceServer
	Config      config.Config
	byBitClient *bybit.Client
	validator   validator.ByBitValidator
}

func NewByBitHttpServerOrder(cfg config.Config) ByBitHttpServerOrder {
	return ByBitHttpServerOrder{
		Config:      cfg,
		validator:   validator.NewByBitValidator(),
		byBitClient: bybit.NewBybitHttpClient(cfg.ByBit.ApiKey, cfg.ByBit.ApiSecret, bybit.WithBaseURL(bybit.MAINNET)),
	}
}

//func (s *ByBitHttpServerOrder) SetRoutes(e *echo.Echo) {

//userGroup := e.Group(string(params.Order))
//userGroup.GET(string(params.Order+"create"), create)
//userGroup.GET(string(params.Order+"amend"), amend)
//userGroup.GET(string(params.Order+"cancel"), cancel)
//}

func (s *ByBitHttpServerOrder) Create(ctx context.Context, in *order.PlaceOrderRequest) (*order.PlaceOrderResponse, error) {

	return &order.PlaceOrderResponse{
		OrderId: "3213213213213",
	}, nil
	// send request to bybit web service and response to grpc request
}
func (s *ByBitHttpServerOrder) Amend(ctx context.Context, in *order.AmendOrderRequest) (*order.PlaceOrderResponse, error) {
	return &order.PlaceOrderResponse{
		OrderId: "3213213213213",
	}, nil
	// send request to bybit web service and response to grpc request
}
func (s *ByBitHttpServerOrder) Cancel(ctx context.Context, in *order.CancelOrderRequest) (*order.PlaceOrderResponse, error) {
	return &order.PlaceOrderResponse{
		OrderId: "3213213213213",
	}, nil
	// send request to bybit web service and response to grpc request
}
