package models_grpc

import (
	"time"
)

const (
	Collection_ByBit_MGIIL string = "ByBitMarketGetInstrumentsInfoLinear"
	Collection_ByBit_MGIII string = "ByBitMarketGetInstrumentsInfoInverse"
	Collection_ByBit_MGIIS string = "ByBitMarketGetInstrumentsInfoSpot"
	Collection_ByBit_MGRL  string = "BybitMarketGetRiskLimit"
)

type BybitMarketGetRiskLimit struct {
	Category  string      `json:"category"`
	Symbol    string      `json:"symbol"`
	RiskLimit []RiskLimit `json:"riskLimit"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
type RiskLimit struct {
	ID                int    `json:"id"`
	IsLowestRisk      int    `json:"isLowestRisk"`
	RiskLimitValue    string `json:"riskLimitValue"`
	MaintenanceMargin string `json:"maintenanceMargin"`
	InitialMargin     string `json:"initialMargin"`
	MaxLeverage       string `json:"maxLeverage"`
	MmDeduction       string `json:"mmDeduction"`
}

type LeverageFilter struct {
	MinLeverage  string `json:"minLeverage"`
	MaxLeverage  string `json:"maxLeverage"`
	LeverageStep string `json:"leverageStep"`
}
type PriceFilter struct {
	MinPrice string `json:"minPrice"`
	MaxPrice string `json:"maxPrice"`
	TickSize string `json:"tickSize"`
}
type LotSizeFilter struct {
	MaxOrderQty         string `json:"maxOrderQty"`
	MaxMktOrderQty      string `json:"maxMktOrderQty"`
	MinOrderQty         string `json:"minOrderQty"`
	QtyStep             string `json:"qtyStep"`
	PostOnlyMaxOrderQty string `json:"postOnlyMaxOrderQty"`
	MinNotionalValue    string `json:"minNotionalValue"`
	BasePrecision       string `json:"basePrecision"`
	QuotePrecision      string `json:"quotePrecision"`
	MinOrderAmt         string `json:"minOrderAmt"`
	MaxOrderAmt         string `json:"maxOrderAmt"`
}
type Phases struct {
	Phase     string `json:"phase"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
type AuctionFeeInfo struct {
	AuctionFeeRate string `json:"auctionFeeRate"`
	TakerFeeRate   string `json:"takerFeeRate"`
	MakerFeeRate   string `json:"makerFeeRate"`
}
type PreListingInfo struct {
	CurAuctionPhase string         `json:"curAuctionPhase"`
	Phases          []Phases       `json:"phases"`
	AuctionFeeInfo  AuctionFeeInfo `json:"auctionFeeInfo"`
}
type RiskParameters struct {
	LimitParameter  string `json:"limitParameter"`
	MarketParameter string `json:"marketParameter"`
}

type ByBitMarketGetInstrumentsInfoLinear struct {
	//ID              json.ObjectID `json:"_id"`
	Symbol             string         `json:"symbol"`
	ContractType       string         `json:"contractType"`
	Status             string         `json:"status"`
	BaseCoin           string         `json:"baseCoin"`
	QuoteCoin          string         `json:"quoteCoin"`
	LaunchTime         string         `json:"launchTime"`
	DeliveryTime       string         `json:"deliveryTime"`
	DeliveryFeeRate    string         `json:"deliveryFeeRate"`
	PriceScale         string         `json:"priceScale"`
	LeverageFilter     LeverageFilter `json:"leverageFilter"`
	PriceFilter        PriceFilter    `json:"priceFilter"`
	LotSizeFilter      LotSizeFilter  `json:"lotSizeFilter"`
	UnifiedMarginTrade bool           `json:"unifiedMarginTrade"`
	FundingInterval    int            `json:"fundingInterval"`
	SettleCoin         string         `json:"settleCoin"`
	CopyTrading        string         `json:"copyTrading"`
	UpperFundingRate   string         `json:"upperFundingRate"`
	LowerFundingRate   string         `json:"lowerFundingRate"`
	IsPreListing       bool           `json:"isPreListing"`
	PreListingInfo     PreListingInfo `json:"preListingInfo"`

	NextPageCursor string `json:"nextPageCursor"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type ByBitMarketGetInstrumentsInfoSpot struct {
	Symbol         string         `json:"symbol"`
	BaseCoin       string         `json:"baseCoin"`
	QuoteCoin      string         `json:"quoteCoin"`
	Innovation     string         `json:"innovation"`
	Status         string         `json:"status"`
	MarginTrading  string         `json:"marginTrading"`
	StTag          string         `json:"stTag"`
	LotSizeFilter  LotSizeFilter  `json:"lotSizeFilter"`
	PriceFilter    PriceFilter    `json:"priceFilter"`
	RiskParameters RiskParameters `json:"riskParameters"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `jsonl:"updated_at"`
}
type ByBitMarketGetInstrumentsInfoInverse struct {
	Symbol             string         `json:"symbol"`
	ContractType       string         `json:"contractType"`
	Status             string         `json:"status"`
	BaseCoin           string         `json:"baseCoin"`
	QuoteCoin          string         `json:"quoteCoin"`
	LaunchTime         string         `json:"launchTime"`
	DeliveryTime       string         `json:"deliveryTime"`
	DeliveryFeeRate    string         `json:"deliveryFeeRate"`
	PriceScale         string         `json:"priceScale"`
	LeverageFilter     LeverageFilter `json:"leverageFilter"`
	PriceFilter        PriceFilter    `json:"priceFilter"`
	LotSizeFilter      LotSizeFilter  `json:"lotSizeFilter"`
	UnifiedMarginTrade bool           `json:"unifiedMarginTrade"`
	FundingInterval    int            `json:"fundingInterval"`
	SettleCoin         string         `json:"settleCoin"`
	CopyTrading        string         `json:"copyTrading"`
	UpperFundingRate   string         `json:"upperFundingRate"`
	LowerFundingRate   string         `json:"lowerFundingRate"`
	IsPreListing       bool           `json:"isPreListing"`
	PreListingInfo     PreListingInfo `json:"preListingInfo"`
	NextPageCursor     string         `json:"nextPageCursor"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
}

type BybitMarketGetTickerSpotDto struct {
	Category       string                     `json:"category"`
	List           []BybitMarketGetTickerSpot `json:"list"`
	NextPageCursor string                     `json:"nextPageCursor"`
}
type BybitMarketGetTickerSpot struct {
	Symbol        string `json:"symbol"`
	Bid1Price     string `json:"bid1Price"`
	Bid1Size      string `json:"bid1Size"`
	Ask1Price     string `json:"ask1Price"`
	Ask1Size      string `json:"ask1Size"`
	LastPrice     string `json:"lastPrice"`
	PrevPrice24H  string `json:"prevPrice24h"`
	Price24HPcnt  string `json:"price24hPcnt"`
	HighPrice24H  string `json:"highPrice24h"`
	LowPrice24H   string `json:"lowPrice24h"`
	Turnover24H   string `json:"turnover24h"`
	Volume24H     string `json:"volume24h"`
	UsdIndexPrice string `json:"usdIndexPrice"`
}
type BybitMarketGetTickerLinear struct {
	Symbol                 string `json:"symbol"`
	LastPrice              string `json:"lastPrice"`
	IndexPrice             string `json:"indexPrice"`
	MarkPrice              string `json:"markPrice"`
	PrevPrice24H           string `json:"prevPrice24h"`
	Price24HPcnt           string `json:"price24hPcnt"`
	HighPrice24H           string `json:"highPrice24h"`
	LowPrice24H            string `json:"lowPrice24h"`
	PrevPrice1H            string `json:"prevPrice1h"`
	OpenInterest           string `json:"openInterest"`
	OpenInterestValue      string `json:"openInterestValue"`
	Turnover24H            string `json:"turnover24h"`
	Volume24H              string `json:"volume24h"`
	FundingRate            string `json:"fundingRate"`
	NextFundingTime        string `json:"nextFundingTime"`
	PredictedDeliveryPrice string `json:"predictedDeliveryPrice"`
	BasisRate              string `json:"basisRate"`
	DeliveryFeeRate        string `json:"deliveryFeeRate"`
	DeliveryTime           string `json:"deliveryTime"`
	Ask1Size               string `json:"ask1Size"`
	Bid1Price              string `json:"bid1Price"`
	Ask1Price              string `json:"ask1Price"`
	Bid1Size               string `json:"bid1Size"`
	Basis                  string `json:"basis"`
	PreOpenPrice           string `json:"preOpenPrice"`
	PreQty                 string `json:"preQty"`
	CurPreListingPhase     string `json:"curPreListingPhase"`
}

type BybitMarketGetTickerLinearDto struct {
	Category string                       `json:"category"`
	List     []BybitMarketGetTickerLinear `json:"list"`
}

type BybitMarketGetTickerInverse struct {
	Symbol                 string `json:"symbol"`
	LastPrice              string `json:"lastPrice"`
	IndexPrice             string `json:"indexPrice"`
	MarkPrice              string `json:"markPrice"`
	PrevPrice24H           string `json:"prevPrice24h"`
	Price24HPcnt           string `json:"price24hPcnt"`
	HighPrice24H           string `json:"highPrice24h"`
	LowPrice24H            string `json:"lowPrice24h"`
	PrevPrice1H            string `json:"prevPrice1h"`
	OpenInterest           string `json:"openInterest"`
	OpenInterestValue      string `json:"openInterestValue"`
	Turnover24H            string `json:"turnover24h"`
	Volume24H              string `json:"volume24h"`
	FundingRate            string `json:"fundingRate"`
	NextFundingTime        string `json:"nextFundingTime"`
	PredictedDeliveryPrice string `json:"predictedDeliveryPrice"`
	BasisRate              string `json:"basisRate"`
	DeliveryFeeRate        string `json:"deliveryFeeRate"`
	DeliveryTime           string `json:"deliveryTime"`
	Ask1Size               string `json:"ask1Size"`
	Bid1Price              string `json:"bid1Price"`
	Ask1Price              string `json:"ask1Price"`
	Bid1Size               string `json:"bid1Size"`
	Basis                  string `json:"basis"`
	PreOpenPrice           string `json:"preOpenPrice"`
	PreQty                 string `json:"preQty"`
	CurPreListingPhase     string `json:"curPreListingPhase"`
}
type BybitMarketGetTickerInverseDto struct {
	Category string                        `json:"category"`
	List     []BybitMarketGetTickerInverse `json:"list"`
}
