package params_bybit_http

import (
	"encoding/json"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
	bybit "github.com/wuhewuhe/bybit.go.api"
	"time"
)

type InstrumentInfoLinearDto struct {
	Symbol             string                     `json:"symbol"`
	ContractType       string                     `json:"contractType"`
	Status             string                     `json:"status"`
	BaseCoin           string                     `json:"baseCoin"`
	QuoteCoin          string                     `json:"quoteCoin"`
	LaunchTime         string                     `json:"launchTime"`
	DeliveryTime       string                     `json:"deliveryTime"`
	DeliveryFeeRate    string                     `json:"deliveryFeeRate"`
	PriceScale         string                     `json:"priceScale"`
	LeverageFilter     models_grpc.LeverageFilter `json:"leverageFilter"`
	PriceFilter        models_grpc.PriceFilter    `json:"priceFilter"`
	LotSizeFilter      models_grpc.LotSizeFilter  `json:"lotSizeFilter"`
	UnifiedMarginTrade bool                       `json:"unifiedMarginTrade"`
	FundingInterval    int                        `json:"fundingInterval"`
	SettleCoin         string                     `json:"settleCoin"`
	CopyTrading        string                     `json:"copyTrading"`
	UpperFundingRate   string                     `json:"upperFundingRate"`
	LowerFundingRate   string                     `json:"lowerFundingRate"`
	IsPreListing       bool                       `json:"isPreListing"`
	PreListingInfo     models_grpc.PreListingInfo `json:"preListingInfo"`
}
type ResultListLinearDto struct {
	Category       string                    `json:"category"`
	List           []InstrumentInfoLinearDto `json:"list"`
	NextPageCursor string                    `json:"nextPageCursor"`
}
type GetInstrumentInfoLinearDto struct {
	RetCode    int                 `json:"retCode"`
	RetMsg     string              `json:"retMsg"`
	Result     ResultListLinearDto `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

type InstrumentInfoSpotDto struct {
	Symbol         string                     `json:"symbol"`
	BaseCoin       string                     `json:"baseCoin"`
	QuoteCoin      string                     `json:"quoteCoin"`
	Innovation     string                     `json:"innovation"`
	Status         string                     `json:"status"`
	MarginTrading  string                     `json:"marginTrading"`
	StTag          string                     `json:"stTag"`
	LotSizeFilter  models_grpc.LotSizeFilter  `json:"lotSizeFilter"`
	PriceFilter    models_grpc.PriceFilter    `json:"priceFilter"`
	RiskParameters models_grpc.RiskParameters `json:"riskParameters"`
}
type ResultListSpotDto struct {
	Category string                  `json:"category"`
	List     []InstrumentInfoSpotDto `json:"list"`
}
type GetInstrumentInfoSpotDto struct {
	RetCode    int               `json:"retCode"`
	RetMsg     string            `json:"retMsg"`
	Result     ResultListSpotDto `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

func ToGetInstrumentInfoSpotDto(data *bybit.ServerResponse) GetInstrumentInfoSpotDto {
	marshal, err := json.Marshal(data)
	var pl GetInstrumentInfoSpotDto
	if err != nil {
		return pl
	}
	err = json.Unmarshal(marshal, &pl)
	if err != nil {
		return pl
	}
	return pl
}
func ToByBitMarketGetInstrumentsInfoSpot(data GetInstrumentInfoSpotDto) []models_grpc.ByBitMarketGetInstrumentsInfoSpot {
	var ret []models_grpc.ByBitMarketGetInstrumentsInfoSpot
	for _, item := range data.Result.List {
		ret = append(ret, models_grpc.ByBitMarketGetInstrumentsInfoSpot{
			Symbol:         item.Symbol,
			Status:         item.Status,
			BaseCoin:       item.BaseCoin,
			QuoteCoin:      item.QuoteCoin,
			PriceFilter:    item.PriceFilter,
			LotSizeFilter:  item.LotSizeFilter,
			Innovation:     item.Innovation,
			MarginTrading:  item.MarginTrading,
			StTag:          item.StTag,
			RiskParameters: item.RiskParameters,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		})
	}
	return ret
}
func ToByBitMarketGetInstrumentsInfoLinear(data GetInstrumentInfoLinearDto) []models_grpc.ByBitMarketGetInstrumentsInfoLinear {
	var ret []models_grpc.ByBitMarketGetInstrumentsInfoLinear
	for _, item := range data.Result.List {
		ret = append(ret, models_grpc.ByBitMarketGetInstrumentsInfoLinear{
			Symbol:             item.Symbol,
			ContractType:       item.ContractType,
			Status:             item.Status,
			BaseCoin:           item.BaseCoin,
			QuoteCoin:          item.QuoteCoin,
			LaunchTime:         item.LaunchTime,
			DeliveryTime:       item.DeliveryTime,
			DeliveryFeeRate:    item.DeliveryFeeRate,
			PriceScale:         item.PriceScale,
			LeverageFilter:     item.LeverageFilter,
			PriceFilter:        item.PriceFilter,
			LotSizeFilter:      item.LotSizeFilter,
			UnifiedMarginTrade: item.UnifiedMarginTrade,
			FundingInterval:    item.FundingInterval,
			SettleCoin:         item.SettleCoin,
			CopyTrading:        item.CopyTrading,
			UpperFundingRate:   item.UpperFundingRate,
			LowerFundingRate:   item.LowerFundingRate,
			IsPreListing:       item.IsPreListing,
			PreListingInfo:     item.PreListingInfo,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		})
	}
	return ret
}
func ToGetInstrumentInfoLinearDto(data *bybit.ServerResponse) GetInstrumentInfoLinearDto {
	marshal, err := json.Marshal(data)
	var pl GetInstrumentInfoLinearDto
	if err != nil {
		return pl
	}
	err = json.Unmarshal(marshal, &pl)
	if err != nil {
		return pl
	}
	return pl
}
