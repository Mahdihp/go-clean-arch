package bybit_grpc_service

import (
	"context"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/order"
	"github.com/bxcodec/go-clean-arch/config"
	params_bybit_http "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/params"
	"github.com/bxcodec/go-clean-arch/internal/validator"
	"github.com/bxcodec/go-clean-arch/util"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

type ByBitHttpServerOrder struct {
	order.UnimplementedOrderServiceServer
	Config      config.Config
	byBitClient *bybit.Client
	validator   validator.ByBitTradeValidator
}

func NewByBitHttpServerOrder(cfg config.Config) ByBitHttpServerOrder {
	return ByBitHttpServerOrder{
		Config:      cfg,
		validator:   validator.NewByBitTradeValidator(),
		byBitClient: bybit.NewBybitHttpClient(cfg.ByBitWs.ApiKey, cfg.ByBitWs.ApiSecret, bybit.WithBaseURL(bybit.MAINNET)),
	}
}

func (s *ByBitHttpServerOrder) Create(ctx context.Context, in *order.PlaceOrderRequest) (*order.PlaceOrderResponse, error) {
	errorList, err := s.validator.ValidateCreate(*in)
	var strErrorList = ""
	if err != nil {
		strErrorList += util.MapToString(errorList)
		return &order.PlaceOrderResponse{}, nil
	}
	params := map[string]interface{}{"category": in.Category, "symbol": in.Symbol,
		"isLeverage": in.IsLeverage, "side": in.Side, "orderType": in.OrderType, "qty": in.Qty,
		"marketUnit": in.MarketUnit, "price": in.Price, "triggerDirection": in.TriggerDirection,
		"orderFilter": in.OrderFilter, "triggerPrice": in.TriggerPrice, "triggerBy": in.TriggerBy,
		"orderIv": in.OrderIv, "timeInForce": in.TimeInForce, "positionIdx": in.PositionIdx,
		"orderLinkId": in.OrderLinkId, "takeProfit": in.TakeProfit, "stopLoss": in.StopLoss,
		"tpTriggerBy": in.TpTriggerBy, "slTriggerBy": in.SlTriggerBy, "reduceOnly": in.ReduceOnly,
		"closeOnTrigger": in.CloseOnTrigger, "smpType": in.SmpType, "mmp": in.Mmp, "tpslMode": in.TpslMode,
		"tpLimitPrice": in.TpLimitPrice, "slLimitPrice": in.SlLimitPrice, "tpOrderType": in.TpOrderType, "slOrderType": in.SlOrderType}
	res, err := s.byBitClient.NewClassicalBybitServiceWithParams(params).PlaceOrder(ctx)
	if err != nil {
		return &order.PlaceOrderResponse{}, nil
	}
	orderDto := params_bybit_http.OrderToOrderDto(res)
	orderPor := params_bybit_http.OrderDtoToPlaceOrderResponse(orderDto)
	return &orderPor, nil
}
func (s *ByBitHttpServerOrder) Amend(ctx context.Context, in *order.AmendOrderRequest) (*order.PlaceOrderResponse, error) {
	errorList, err := s.validator.ValidateAmend(*in)
	var strErrorList = ""
	if err != nil {
		strErrorList += util.MapToString(errorList)
		return &order.PlaceOrderResponse{}, nil
	}

	params := map[string]interface{}{"category": in.Category, "symbol": in.Symbol,
		"orderId": in.OrderId, "orderLinkId": in.OrderLinkId, "orderIv": in.OrderIv,
		"triggerPrice": in.TriggerPrice, "qty": in.Qty, "price": in.Price, "tpslMode": in.TpslMode,
		"takeProfit": in.TakeProfit, "stopLoss": in.StopLoss, "tpTriggerBy": in.TpTriggerBy,
		"slTriggerBy": in.SlTriggerBy, "triggerBy": in.TriggerBy, "tpLimitPrice": in.TpLimitPrice, "slLimitPrice": in.SlLimitPrice}
	res, err := s.byBitClient.NewClassicalBybitServiceWithParams(params).AmendOrder(ctx)
	if err != nil {
		return &order.PlaceOrderResponse{}, nil
	}
	orderDto := params_bybit_http.OrderToOrderDto(res)
	orderPor := params_bybit_http.OrderDtoToPlaceOrderResponse(orderDto)
	return &orderPor, nil
}
func (s *ByBitHttpServerOrder) Cancel(ctx context.Context, in *order.CancelOrderRequest) (*order.PlaceOrderResponse, error) {
	errorList, err := s.validator.ValidateCancel(*in)
	var strErrorList = ""
	if err != nil {
		strErrorList += util.MapToString(errorList)
		return &order.PlaceOrderResponse{}, nil
	}

	params := map[string]interface{}{"category": in.Category, "symbol": in.Symbol,
		"orderId": in.OrderId, "orderLinkId": in.OrderLinkId, "orderFilter": in.OrderFilter}
	res, err := s.byBitClient.NewClassicalBybitServiceWithParams(params).CancelOrder(ctx)
	if err != nil {
		return &order.PlaceOrderResponse{}, nil
	}
	orderDto := params_bybit_http.OrderToOrderDto(res)
	orderPor := params_bybit_http.OrderDtoToPlaceOrderResponse(orderDto)
	return &orderPor, nil
}

//func (s *ByBitHttpServerOrder) SetRoutes(e *echo.Echo) {

//userGroup := e.Group(string(params.Order))
//userGroup.GET(string(params.Order+"create"), create)
//userGroup.GET(string(params.Order+"amend"), amend)
//userGroup.GET(string(params.Order+"cancel"), cancel)
//}
