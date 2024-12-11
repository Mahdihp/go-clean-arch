package bybit_grpc_service

import (
	"context"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/market"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/db/mongodb"
	params_bybit_http "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/params"
	"github.com/bxcodec/go-clean-arch/internal/validator"
	"github.com/bxcodec/go-clean-arch/util"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

type ByBitHttpServerMarket struct {
	market.UnimplementedMarketServiceServer
	Config         config.Config
	byBitClient    *bybit.Client
	validator      validator.ByBitMarketValidator
	bybitMarketSvc ByBitMarketService
}

func NewByBitHttpServerMarket(cfg config.Config, db *mongodb.MongoDb) ByBitHttpServerMarket {
	return ByBitHttpServerMarket{
		Config:         cfg,
		validator:      validator.NewByBitMarketValidator(),
		byBitClient:    bybit.NewBybitHttpClient(cfg.ByBitWs.ApiKey, cfg.ByBitWs.ApiSecret, bybit.WithBaseURL(bybit.MAINNET)),
		bybitMarketSvc: NewByBitMarketService(db),
	}
}

func (s *ByBitHttpServerMarket) GetInstrumentsInfoLinear(ctx context.Context, in *market.GetInstrumentsInfoRequest) (*market.GetInstrumentsInfoLinearResponse, error) {

	//category, _ := s.bybitMarketSvc.FindAll(ctx, models_grpc.Coll_ByBitMarketGetInstrumentsInfoLinear)

	//s.bybitMarketSvc.FindOneAndUpdateLinear(ctx, models_grpc.Coll_ByBitMarketGetInstrumentsInfoLinear, category)
	//for i, linear := range category {
	//	fmt.Println(i, linear.Symbol, linear.CreatedAt)
	//}
	return &market.GetInstrumentsInfoLinearResponse{}, nil
}

func (s *ByBitHttpServerMarket) GetInstrumentsInfoOption(ctx context.Context, in *market.GetInstrumentsInfoRequest) (*market.GetInstrumentsInfoOptionResponse, error) {
	return &market.GetInstrumentsInfoOptionResponse{}, nil
}

func (s *ByBitHttpServerMarket) GetInstrumentsInfoSpot(ctx context.Context, in *market.GetInstrumentsInfoRequest) (*market.GetInstrumentsInfoSpotResponse, error) {
	return &market.GetInstrumentsInfoSpotResponse{}, nil
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
