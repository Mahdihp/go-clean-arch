package bybit_grpc_service

import (
	"context"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/position"
	"github.com/bxcodec/go-clean-arch/config"
	params_bybit_http "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/params"
	"github.com/bxcodec/go-clean-arch/internal/validator"
	"github.com/bxcodec/go-clean-arch/util"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

type ByBitHttpServerPosition struct {
	position.UnimplementedPositionServiceServer
	Config      config.Config
	byBitClient *bybit.Client
	validator   validator.ByBitValidator
}

func NewByBitHttpServerPosition(cfg config.Config) ByBitHttpServerPosition {
	return ByBitHttpServerPosition{
		Config:      cfg,
		validator:   validator.NewByBitValidator(),
		byBitClient: bybit.NewBybitHttpClient(cfg.ByBit.ApiKey, cfg.ByBit.ApiSecret, bybit.WithBaseURL(bybit.MAINNET)),
	}
}

func (s *ByBitHttpServerPosition) GetPositionInfo(ctx context.Context, in *position.PositionInfoRequest) (*position.PositionInfoResponse, error) {

	errorList, err := s.validator.ValidateGetPositionInfo(*in)
	var strErrorList = ""
	if err != nil {
		strErrorList += util.MapToString(errorList)
		return &position.PositionInfoResponse{RetMsg: strErrorList}, nil
	}
	//params := map[string]interface{}{"category": "linear"}
	params := map[string]interface{}{"category": in.Category.Category, "settleCoin": in.Category.SettleCoin}
	res, err := s.byBitClient.NewClassicalBybitServiceWithParams(params).GetPositionList(ctx)

	if err != nil {
		return &position.PositionInfoResponse{RetMsg: err.Error()}, nil
	}
	dataPosition := params_bybit_http.StringToPositionList(res)
	positionList := params_bybit_http.PositionListToDataList(dataPosition)
	return &positionList, nil
}
func (s *ByBitHttpServerPosition) SetLeverage(ctx context.Context, in *position.SetLeverageRequest) (*position.BaseResponse, error) {
	errorList, err := s.validator.ValidateSetLeverage(*in)
	var strErrorList = ""
	if err != nil {
		strErrorList += util.MapToString(errorList)
		return &position.BaseResponse{RetMsg: strErrorList}, nil
	}
	//params := map[string]interface{}{"category": "CONTRACT"}
	params := map[string]interface{}{"category": in.Category, "symbol": in.Symbol, "buyLeverage": in.BuyLeverage, "sellLeverage": in.SellLeverage}
	res, err := s.byBitClient.NewClassicalBybitServiceWithParams(params).SetPositionLeverage(ctx)
	if err != nil {
		return &position.BaseResponse{RetMsg: err.Error()}, nil
	}

	return &position.BaseResponse{
		RetCode:    int32(res.RetCode),
		RetMsg:     res.RetMsg,
		Result:     util.InterfaceToString(res.Result),
		RetExtInfo: util.InterfaceToString(res.RetExtInfo),
	}, nil
}
func (s *ByBitHttpServerPosition) SwitchIsolated(ctx context.Context, in *position.SwitchIsolatedRequest) (*position.BaseResponse, error) {
	errorList, err := s.validator.ValidateSwitchIsolated(*in)
	var strErrorList = ""
	if err != nil {
		strErrorList += util.MapToString(errorList)
		return &position.BaseResponse{RetMsg: strErrorList}, nil
	}
	//params := map[string]interface{}{"category": "CONTRACT"}
	params := map[string]interface{}{"category": in.Category, "symbol": in.Symbol, "tradeMode": in.TradeMode, "buyLeverage": in.BuyLeverage, "sellLeverage": in.SellLeverage}
	res, err := s.byBitClient.NewClassicalBybitServiceWithParams(params).SwitchPositionMargin(ctx)
	if err != nil {
		return &position.BaseResponse{RetMsg: err.Error()}, nil
	}

	return &position.BaseResponse{
		RetCode:    int32(res.RetCode),
		RetMsg:     res.RetMsg,
		Result:     util.InterfaceToString(res.Result),
		RetExtInfo: util.InterfaceToString(res.RetExtInfo),
	}, nil
}
func (s *ByBitHttpServerPosition) SwitchMode(ctx context.Context, in *position.SwitchModeRequest) (*position.BaseResponse, error) {
	errorList, err := s.validator.ValidateSwitchMode(*in)
	var strErrorList = ""
	if err != nil {
		strErrorList += util.MapToString(errorList)
		return &position.BaseResponse{RetMsg: strErrorList}, nil
	}

	//params := map[string]interface{}{"category": "CONTRACT"}
	params := map[string]interface{}{"category": in.Category, "symbol": in.Symbol, "coin": in.Coin, "mode": in.Mode}
	res, err := s.byBitClient.NewClassicalBybitServiceWithParams(params).SwitchPositionMode(ctx)
	if err != nil {
		return &position.BaseResponse{RetMsg: err.Error()}, nil
	}

	return &position.BaseResponse{
		RetCode:    int32(res.RetCode),
		RetMsg:     res.RetMsg,
		Result:     util.InterfaceToString(res.Result),
		RetExtInfo: util.InterfaceToString(res.RetExtInfo),
	}, nil
}
func (s *ByBitHttpServerPosition) TradingStop(ctx context.Context, in *position.TradingStopRequest) (*position.BaseResponse, error) {
	errorList, err := s.validator.ValidateTradingStop(*in)
	var strErrorList = ""
	if err != nil {
		strErrorList += util.MapToString(errorList)
		return &position.BaseResponse{RetMsg: strErrorList}, nil
	}

	//params := map[string]interface{}{"category": "CONTRACT"}
	params := map[string]interface{}{"category": in.Category, "symbol": in.Symbol, "takeProfit": in.TakeProfit, "stopLoss": in.StopLoss,
		"trailingStop": in.TrailingStop, "tpTriggerBy": in.TpTriggerBy, "slTriggerBy": in.SlTriggerBy, "activePrice": in.ActivePrice,
		"tpslMode": in.TpslMode, "tpSize": in.TpSize, "slSize": in.SlSize, "tpLimitPrice": in.TpLimitPrice, "slLimitPrice": in.SlLimitPrice,
		"tpOrderType": in.TpOrderType, "slOrderType": in.SlOrderType, "positionIdx": in.PositionIdx}
	res, err := s.byBitClient.NewClassicalBybitServiceWithParams(params).SetPositionTradingStop(ctx)
	if err != nil {
		return &position.BaseResponse{RetMsg: err.Error()}, nil
	}

	return &position.BaseResponse{
		RetCode:    int32(res.RetCode),
		RetMsg:     res.RetMsg,
		Result:     util.InterfaceToString(res.Result),
		RetExtInfo: util.InterfaceToString(res.RetExtInfo),
	}, nil
}

//func (s *ByBitHttpServerOrder) SetRoutes(e *echo.Echo) {
//userGroup := e.Group(string(params.Order))
//userGroup.GET(string(params.Order+"create"), create)
//userGroup.GET(string(params.Order+"amend"), amend)
//userGroup.GET(string(params.Order+"cancel"), cancel)
//}
