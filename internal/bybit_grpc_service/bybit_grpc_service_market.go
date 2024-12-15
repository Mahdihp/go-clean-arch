package bybit_grpc_service

import (
	"context"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/market"
	"github.com/bxcodec/go-clean-arch/config"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
	params_bybit_http "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/params"
	"github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/repository"
	"github.com/bxcodec/go-clean-arch/internal/validator"
	"github.com/bxcodec/go-clean-arch/util"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

type ByBitHttpServerMarket struct {
	market.UnimplementedMarketServiceServer
	Config         config.Config
	byBitClient    *bybit.Client
	validator      validator.ByBitMarketValidator
	bybitMarketSvc repository.ByBitMarketRepository
}

func NewByBitHttpServerMarket(cfg config.Config, bybitMarketSvc repository.ByBitMarketRepository) ByBitHttpServerMarket {
	return ByBitHttpServerMarket{
		Config:         cfg,
		validator:      validator.NewByBitMarketValidator(),
		byBitClient:    bybit.NewBybitHttpClient(cfg.ByBitWs.ApiKey, cfg.ByBitWs.ApiSecret, bybit.WithBaseURL(bybit.MAINNET)),
		bybitMarketSvc: bybitMarketSvc,
	}
}

func (s *ByBitHttpServerMarket) GetTickersSpot(ctx context.Context, in *market.GetTickersRequest) (*market.GetTickersSpotResponse, error) {
	//collection := selectCollection(in.Category)
	spot, err := s.bybitMarketSvc.FindTickerSpot(ctx)
	if err != nil {
		return &market.GetTickersSpotResponse{}, err
	}
	response := params_bybit_http.ToGetTickersSpot(spot)
	return &response, nil
}
func (s *ByBitHttpServerMarket) GetTickersLinear(ctx context.Context, in *market.GetTickersRequest) (*market.GetTickersLinearResponse, error) {
	//collection := selectCollection(in.Category)
	inverse, err := s.bybitMarketSvc.FindTickerLinear(ctx)
	if err != nil {
		return &market.GetTickersLinearResponse{}, err
	}
	response := params_bybit_http.ToGetTickersLinear(inverse)
	return &response, nil
}
func (s *ByBitHttpServerMarket) GetTickersInverse(ctx context.Context, in *market.GetTickersRequest) (*market.GetTickersInverseResponse, error) {
	//collection := selectCollection(in.Category)
	inverse, err := s.bybitMarketSvc.FindTickerInverse(ctx)
	if err != nil {
		return &market.GetTickersInverseResponse{}, err
	}
	response := params_bybit_http.ToGetTickersInverse(inverse)
	return &response, nil
}

func (s *ByBitHttpServerMarket) GetInstrumentsInfoLinear(ctx context.Context, in *market.GetInstrumentsInfoRequest) (*market.GetInstrumentsInfoLinearResponse, error) {
	collection := selectCollection(in.Category)
	linears, err := s.bybitMarketSvc.FindAllLinear(ctx, collection, in.Symbol)
	if err != nil {
		return &market.GetInstrumentsInfoLinearResponse{}, err
	}
	response := params_bybit_http.ToGetInstrumentsInfoLinearResponse(linears)
	return &response, nil
}
func (s *ByBitHttpServerMarket) GetInstrumentsInfoInverse(ctx context.Context, in *market.GetInstrumentsInfoRequest) (*market.GetInstrumentsInfoInverseResponse, error) {
	collection := selectCollection(in.Category)
	spots, err := s.bybitMarketSvc.FindAllInverse(ctx, collection, in.Symbol)
	if err != nil {
		return &market.GetInstrumentsInfoInverseResponse{}, err
	}

	response := params_bybit_http.ToGetInstrumentsInfoInverseResponse(spots)
	return &response, nil
}
func (s *ByBitHttpServerMarket) GetInstrumentsInfoSpot(ctx context.Context, in *market.GetInstrumentsInfoRequest) (*market.GetInstrumentsInfoSpotResponse, error) {
	collection := selectCollection(in.Category)
	spots, err := s.bybitMarketSvc.FindAllSpot(ctx, collection, in.Symbol)
	if err != nil {
		return &market.GetInstrumentsInfoSpotResponse{}, err
	}

	response := params_bybit_http.ToGetInstrumentsInfoSpotResponse(spots)
	return &response, nil
}
func (s *ByBitHttpServerMarket) GetKline(ctx context.Context, in *market.GetKlineRequest) (*market.GetKlineResponse, error) {
	errorList, err := s.validator.ValidateGetKline(in)
	var strErrorList = ""
	if err != nil {
		strErrorList += util.MapToString(errorList)
		return &market.GetKlineResponse{}, nil
	}

	params := map[string]interface{}{"category": in.Category, "symbol": &in.Symbol,
		"interval": in.Interval, "start": in.Start, "end": in.End, "limit": in.Limit}
	res, err := s.byBitClient.NewClassicalBybitServiceWithParams(params).GetMarketKline(ctx)
	if err != nil {
		return &market.GetKlineResponse{}, nil
	}

	marketDto := params_bybit_http.ToGetKlineDto(res)
	getKlineDto := params_bybit_http.ToGetKlineResponse(marketDto)
	return &getKlineDto, nil
}
func (s *ByBitHttpServerMarket) GetRiskLimitLinear(ctx context.Context, in *market.GetRiskLimitRequest) (*market.GetRiskLimitResponse, error) {
	return nil, nil
}
func (s *ByBitHttpServerMarket) GetRiskLimitInverse(ctx context.Context, in *market.GetRiskLimitRequest) (*market.GetRiskLimitResponse, error) {
	return nil, nil
}

func selectCollection(collName string) string {
	switch collName {
	case "spot":
		return models_grpc.Collection_ByBit_MGIIS
	case "linear":
		return models_grpc.Collection_ByBit_MGIIL
	case "inverse":
		return models_grpc.Collection_ByBit_MGIII
	default:
		return ""
	}
}
