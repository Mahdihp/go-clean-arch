package bybit_http

import (
	"context"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/order"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/position"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/labstack/echo/v4"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

type ByBitHttpServer struct {
	Config      config.Config
	byBitClient *bybit.Client
	order.UnimplementedOrderServiceServer
	position.UnimplementedPositionServiceServer
}

func NewByBitHttpService(cfg config.Config) ByBitHttpServer {
	return ByBitHttpServer{
		Config:      cfg,
		byBitClient: bybit.NewBybitHttpClient(cfg.ByBit.ApiKey, cfg.ByBit.ApiSecret, bybit.WithBaseURL(bybit.MAINNET)),
	}
}

func (s *ByBitHttpServer) SetRoutes(e *echo.Echo) {

	//userGroup := e.Group(string(params.Order))
	//userGroup.GET(string(params.Order+"create"), create)
	//userGroup.GET(string(params.Order+"amend"), amend)
	//userGroup.GET(string(params.Order+"cancel"), cancel)
}

func (s *ByBitHttpServer) Create(ctx context.Context, in *order.PlaceOrderRequest) (*order.PlaceOrderResponse, error) {
	return &order.PlaceOrderResponse{
		OrderId: "3213213213213",
	}, nil
	// send request to bybit web service and response to grpc request
}

func (s *ByBitHttpServer) Amend(ctx context.Context, in *order.AmendOrderRequest) (*order.PlaceOrderResponse, error) {
	return &order.PlaceOrderResponse{
		OrderId: "3213213213213",
	}, nil
	// send request to bybit web service and response to grpc request
}
func (s *ByBitHttpServer) Cancel(ctx context.Context, in *order.CancelOrderRequest) (*order.PlaceOrderResponse, error) {
	return &order.PlaceOrderResponse{
		OrderId: "3213213213213",
	}, nil
	// send request to bybit web service and response to grpc request
}
