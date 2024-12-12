package params_bybit_grpc

import (
	"encoding/json"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/market"
	models_grpc "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/models"
	"github.com/bxcodec/go-clean-arch/util"
	bybit "github.com/wuhewuhe/bybit.go.api"
	"google.golang.org/protobuf/types/known/anypb"
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
type InstrumentInfoInverseDto struct {
	Symbol     string `json:"symbol"`
	BaseCoin   string `json:"baseCoin"`
	QuoteCoin  string `json:"quoteCoin"`
	Innovation string `json:"innovation"`
	Status     string `json:"status"`
}
type ResultListInverseDto struct {
	Category string                     `json:"category"`
	List     []InstrumentInfoInverseDto `json:"list"`
}
type GetInstrumentInfoInverseDto struct {
	RetCode    int                  `json:"retCode"`
	RetMsg     string               `json:"retMsg"`
	Result     ResultListInverseDto `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

type GetKlineResponseDto struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Symbol   string     `json:"symbol"`
		Category string     `json:"category"`
		List     [][]string `json:"list"`
	} `json:"result"`
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
func ToByBitMarketGetInstrumentsInfoSpot(data GetInstrumentInfoSpotDto, time int64) []models_grpc.ByBitMarketGetInstrumentsInfoSpot {
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
			CreatedAt:      util.TimestampToTime(time),
			UpdatedAt:      util.TimestampToTime(time),
		})
	}
	return ret
}
func ToByBitMarketGetInstrumentsInfoLinear(data GetInstrumentInfoLinearDto, time int64) []models_grpc.ByBitMarketGetInstrumentsInfoLinear {
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
			CreatedAt:          util.TimestampToTime(time),
			UpdatedAt:          util.TimestampToTime(time),
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

func ToGetKlineResponse(data GetKlineResponseDto) market.GetKlineResponse {
	response := market.GetKlineResponse{}
	response.RetMsg = data.RetMsg
	response.RetCode = int32(data.RetCode)
	var toAny *anypb.Any
	toAny, err := util.ConvertInterfaceToAny(data.Result.List)
	if err != nil {
		toAny = nil
	}
	response.Result = &market.GetKlineResponse_Result{
		Symbol:   data.Result.Symbol,
		Category: data.Result.Category,
		List:     toAny,
	}
	response.Time = uint64(data.Time)
	return response
}
func ToGetKlineDto(data *bybit.ServerResponse) GetKlineResponseDto {
	marshal, err := json.Marshal(data)
	var pl GetKlineResponseDto
	if err != nil {
		return pl
	}
	err = json.Unmarshal(marshal, &pl)
	if err != nil {
		return pl
	}
	return pl
}
func ToGetInstrumentsInfoInverseResponse(data []models_grpc.ByBitMarketGetInstrumentsInfoInverse) market.GetInstrumentsInfoInverseResponse {
	var mainOutput market.GetInstrumentsInfoInverseResponse
	var resultOutput market.GetInstrumentsInfoInverseResponse_Result
	mainOutput.RetCode = 200
	mainOutput.RetMsg = "OK"
	mainOutput.Time = uint64(time.Now().UnixMilli())

	mainOutput.Result = &resultOutput

	var Phases []*market.Phases
	for _, item := range data {

		for _, phase := range item.PreListingInfo.Phases {
			Phases = append(Phases, &market.Phases{
				Phase:     phase.Phase,
				StartTime: phase.StartTime,
				EndTime:   phase.EndTime,
			})
		}

		resultOutput.List = append(resultOutput.List, &market.GetInstrumentsInfoInverseResponse_List{
			Symbol:          item.Symbol,
			ContractType:    item.ContractType,
			Status:          item.Status,
			BaseCoin:        item.BaseCoin,
			QuoteCoin:       item.QuoteCoin,
			LaunchTime:      item.LaunchTime,
			DeliveryTime:    item.DeliveryTime,
			DeliveryFeeRate: item.DeliveryFeeRate,
			PriceScale:      item.PriceScale,
			LeverageFilter: &market.LeverageFilter{
				MinLeverage:  item.LeverageFilter.MinLeverage,
				MaxLeverage:  item.LeverageFilter.MaxLeverage,
				LeverageStep: item.LeverageFilter.LeverageStep,
			},
			PriceFilter: &market.PriceFilter{
				MinPrice: item.PriceFilter.MinPrice,
				MaxPrice: item.PriceFilter.MaxPrice,
				TickSize: item.PriceFilter.TickSize,
			},
			LotSizeFilter: &market.LotSizeFilter{
				MaxOrderQty:         item.LotSizeFilter.MaxOrderQty,
				MinOrderQty:         item.LotSizeFilter.MinOrderQty,
				QtyStep:             item.LotSizeFilter.QtyStep,
				PostOnlyMaxOrderQty: item.LotSizeFilter.PostOnlyMaxOrderQty,
				MinNotionalValue:    item.LotSizeFilter.MinNotionalValue,
				MaxMktOrderQty:      item.LotSizeFilter.MaxMktOrderQty,
			},
			UnifiedMarginTrade: item.UnifiedMarginTrade,
			FundingInterval:    uint32(item.FundingInterval),
			SettleCoin:         item.SettleCoin,
			CopyTrading:        item.CopyTrading,
			UpperFundingRate:   item.UpperFundingRate,
			LowerFundingRate:   item.LowerFundingRate,
			IsPreListing:       item.IsPreListing,
			PreListingInfo: &market.PreListingInfo{
				CurAuctionPhase: item.PreListingInfo.CurAuctionPhase,
				Phases:          Phases,
				AuctionFeeInfo: &market.AuctionFeeInfo{
					AuctionFeeRate: item.PreListingInfo.AuctionFeeInfo.AuctionFeeRate,
					TakerFeeRate:   item.PreListingInfo.AuctionFeeInfo.TakerFeeRate,
					MakerFeeRate:   item.PreListingInfo.AuctionFeeInfo.MakerFeeRate,
				},
			},
		})
	}
	return mainOutput
}
func ToGetInstrumentsInfoSpotResponse(data []models_grpc.ByBitMarketGetInstrumentsInfoSpot) market.GetInstrumentsInfoSpotResponse {
	var mainOutput market.GetInstrumentsInfoSpotResponse
	var resultOutput market.GetInstrumentsInfoSpotResponse_Result
	mainOutput.RetCode = 200
	mainOutput.RetMsg = "OK"
	mainOutput.Time = uint64(time.Now().UnixMilli())

	mainOutput.Result = &resultOutput

	for _, item := range data {

		resultOutput.List = append(resultOutput.List, &market.GetInstrumentsInfoSpotResponse_List{
			Symbol:    item.Symbol,
			Status:    item.Status,
			BaseCoin:  item.BaseCoin,
			QuoteCoin: item.QuoteCoin,
			PriceFilter: &market.PriceFilter{
				TickSize: item.PriceFilter.TickSize,
			},
			LotSizeFilter: &market.LotSizeFilter{
				MaxOrderQty:    item.LotSizeFilter.MaxOrderQty,
				MinOrderQty:    item.LotSizeFilter.MinOrderQty,
				BasePrecision:  item.LotSizeFilter.BasePrecision,
				QuotePrecision: item.LotSizeFilter.QuotePrecision,
				MinOrderAmt:    item.LotSizeFilter.MaxOrderQty,
				MaxOrderAmt:    item.LotSizeFilter.MaxOrderQty,
			},
		})
	}
	return mainOutput
}
func ToGetInstrumentsInfoLinearResponse(data []models_grpc.ByBitMarketGetInstrumentsInfoLinear) market.GetInstrumentsInfoLinearResponse {
	var mainOutput market.GetInstrumentsInfoLinearResponse
	var resultOutput market.GetInstrumentsInfoLinearResponse_Result
	mainOutput.RetCode = 200
	mainOutput.RetMsg = "OK"
	mainOutput.Time = uint64(time.Now().UnixMilli())

	mainOutput.Result = &resultOutput

	var Phases []*market.Phases
	for _, item := range data {

		for _, phase := range item.PreListingInfo.Phases {
			Phases = append(Phases, &market.Phases{
				Phase:     phase.Phase,
				StartTime: phase.StartTime,
				EndTime:   phase.EndTime,
			})
		}

		resultOutput.List = append(resultOutput.List, &market.GetInstrumentsInfoLinearResponse_List{
			Symbol:          item.Symbol,
			ContractType:    item.ContractType,
			Status:          item.Status,
			BaseCoin:        item.BaseCoin,
			QuoteCoin:       item.QuoteCoin,
			LaunchTime:      item.LaunchTime,
			DeliveryTime:    item.DeliveryTime,
			DeliveryFeeRate: item.DeliveryFeeRate,
			PriceScale:      item.PriceScale,
			LeverageFilter: &market.LeverageFilter{
				MinLeverage:  item.LeverageFilter.MinLeverage,
				MaxLeverage:  item.LeverageFilter.MaxLeverage,
				LeverageStep: item.LeverageFilter.LeverageStep,
			},
			PriceFilter: &market.PriceFilter{
				MinPrice: item.PriceFilter.MinPrice,
				MaxPrice: item.PriceFilter.MaxPrice,
				TickSize: item.PriceFilter.TickSize,
			},
			LotSizeFilter: &market.LotSizeFilter{
				MaxOrderQty:         item.LotSizeFilter.MaxOrderQty,
				MaxMktOrderQty:      item.LotSizeFilter.MaxMktOrderQty,
				MinOrderQty:         item.LotSizeFilter.MinOrderQty,
				QtyStep:             item.LotSizeFilter.QtyStep,
				PostOnlyMaxOrderQty: item.LotSizeFilter.PostOnlyMaxOrderQty,
				MinNotionalValue:    item.LotSizeFilter.MinNotionalValue,
				//BasePrecision     :  item.LotSizeFilter.BasePrecision,
				//QuotePrecision    :  item.LotSizeFilter.QuotePrecision,
				//MinOrderAmt      :   item.LotSizeFilter.MaxOrderQty,
				//MaxOrderAmt       :  item.LotSizeFilter.MaxOrderQty,
			},
			UnifiedMarginTrade: item.UnifiedMarginTrade,
			FundingInterval:    uint32(item.FundingInterval),
			SettleCoin:         item.SettleCoin,
			CopyTrading:        item.CopyTrading,
			UpperFundingRate:   item.UpperFundingRate,
			LowerFundingRate:   item.LowerFundingRate,
			IsPreListing:       item.IsPreListing,
			PreListingInfo: &market.PreListingInfo{
				CurAuctionPhase: item.PreListingInfo.CurAuctionPhase,
				Phases:          Phases,
				AuctionFeeInfo: &market.AuctionFeeInfo{
					AuctionFeeRate: item.PreListingInfo.AuctionFeeInfo.AuctionFeeRate,
					TakerFeeRate:   item.PreListingInfo.AuctionFeeInfo.TakerFeeRate,
					MakerFeeRate:   item.PreListingInfo.AuctionFeeInfo.MakerFeeRate,
				},
			},
		})
	}
	return mainOutput
}
