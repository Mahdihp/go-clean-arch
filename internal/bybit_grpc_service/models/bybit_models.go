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
	Category  string      `bson:"category"`
	Symbol    string      `bson:"symbol"`
	RiskLimit []RiskLimit `bson:"riskLimit"`
	CreatedAt time.Time   `bson:"created_at"`
	UpdatedAt time.Time   `bson:"updated_at"`
}
type RiskLimit struct {
	ID                int    `bson:"id"`
	IsLowestRisk      int    `bson:"isLowestRisk"`
	RiskLimitValue    string `bson:"riskLimitValue"`
	MaintenanceMargin string `bson:"maintenanceMargin"`
	InitialMargin     string `bson:"initialMargin"`
	MaxLeverage       string `bson:"maxLeverage"`
	MmDeduction       string `bson:"mmDeduction"`
}

type LeverageFilter struct {
	MinLeverage  string `bson:"minLeverage"`
	MaxLeverage  string `bson:"maxLeverage"`
	LeverageStep string `bson:"leverageStep"`
}
type PriceFilter struct {
	MinPrice string `bson:"minPrice"`
	MaxPrice string `bson:"maxPrice"`
	TickSize string `bson:"tickSize"`
}
type LotSizeFilter struct {
	MaxOrderQty         string `bson:"maxOrderQty"`
	MaxMktOrderQty      string `bson:"maxMktOrderQty"`
	MinOrderQty         string `bson:"minOrderQty"`
	QtyStep             string `bson:"qtyStep"`
	PostOnlyMaxOrderQty string `bson:"postOnlyMaxOrderQty"`
	MinNotionalValue    string `bson:"minNotionalValue"`
	BasePrecision       string `bson:"basePrecision"`
	QuotePrecision      string `bson:"quotePrecision"`
	MinOrderAmt         string `bson:"minOrderAmt"`
	MaxOrderAmt         string `bson:"maxOrderAmt"`
}
type Phases struct {
	Phase     string `bson:"phase"`
	StartTime string `bson:"startTime"`
	EndTime   string `bson:"endTime"`
}
type AuctionFeeInfo struct {
	AuctionFeeRate string `bson:"auctionFeeRate"`
	TakerFeeRate   string `bson:"takerFeeRate"`
	MakerFeeRate   string `bson:"makerFeeRate"`
}
type PreListingInfo struct {
	CurAuctionPhase string         `bson:"curAuctionPhase"`
	Phases          []Phases       `bson:"phases"`
	AuctionFeeInfo  AuctionFeeInfo `bson:"auctionFeeInfo"`
}
type RiskParameters struct {
	LimitParameter  string `bson:"limitParameter"`
	MarketParameter string `bson:"marketParameter"`
}

type ByBitMarketGetInstrumentsInfoLinear struct {
	//ID              bson.ObjectID `bson:"_id"`
	Symbol             string         `bson:"symbol"`
	ContractType       string         `bson:"contractType"`
	Status             string         `bson:"status"`
	BaseCoin           string         `bson:"baseCoin"`
	QuoteCoin          string         `bson:"quoteCoin"`
	LaunchTime         string         `bson:"launchTime"`
	DeliveryTime       string         `bson:"deliveryTime"`
	DeliveryFeeRate    string         `bson:"deliveryFeeRate"`
	PriceScale         string         `bson:"priceScale"`
	LeverageFilter     LeverageFilter `bson:"leverageFilter"`
	PriceFilter        PriceFilter    `bson:"priceFilter"`
	LotSizeFilter      LotSizeFilter  `bson:"lotSizeFilter"`
	UnifiedMarginTrade bool           `bson:"unifiedMarginTrade"`
	FundingInterval    int            `bson:"fundingInterval"`
	SettleCoin         string         `bson:"settleCoin"`
	CopyTrading        string         `bson:"copyTrading"`
	UpperFundingRate   string         `bson:"upperFundingRate"`
	LowerFundingRate   string         `bson:"lowerFundingRate"`
	IsPreListing       bool           `bson:"isPreListing"`
	PreListingInfo     PreListingInfo `bson:"preListingInfo"`

	NextPageCursor string `bson:"nextPageCursor"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
type ByBitMarketGetInstrumentsInfoSpot struct {
	Symbol         string         `bson:"symbol"`
	BaseCoin       string         `bson:"baseCoin"`
	QuoteCoin      string         `bson:"quoteCoin"`
	Innovation     string         `bson:"innovation"`
	Status         string         `bson:"status"`
	MarginTrading  string         `bson:"marginTrading"`
	StTag          string         `bson:"stTag"`
	LotSizeFilter  LotSizeFilter  `bson:"lotSizeFilter"`
	PriceFilter    PriceFilter    `bson:"priceFilter"`
	RiskParameters RiskParameters `bson:"riskParameters"`
	CreatedAt      time.Time      `bson:"created_at"`
	UpdatedAt      time.Time      `bson:"updated_at"`
}
type ByBitMarketGetInstrumentsInfoInverse struct {
	Symbol             string         `bson:"symbol"`
	ContractType       string         `bson:"contractType"`
	Status             string         `bson:"status"`
	BaseCoin           string         `bson:"baseCoin"`
	QuoteCoin          string         `bson:"quoteCoin"`
	LaunchTime         string         `bson:"launchTime"`
	DeliveryTime       string         `bson:"deliveryTime"`
	DeliveryFeeRate    string         `bson:"deliveryFeeRate"`
	PriceScale         string         `bson:"priceScale"`
	LeverageFilter     LeverageFilter `bson:"leverageFilter"`
	PriceFilter        PriceFilter    `bson:"priceFilter"`
	LotSizeFilter      LotSizeFilter  `bson:"lotSizeFilter"`
	UnifiedMarginTrade bool           `bson:"unifiedMarginTrade"`
	FundingInterval    int            `bson:"fundingInterval"`
	SettleCoin         string         `bson:"settleCoin"`
	CopyTrading        string         `bson:"copyTrading"`
	UpperFundingRate   string         `bson:"upperFundingRate"`
	LowerFundingRate   string         `bson:"lowerFundingRate"`
	IsPreListing       bool           `bson:"isPreListing"`
	PreListingInfo     PreListingInfo `bson:"preListingInfo"`
	NextPageCursor     string         `bson:"nextPageCursor"`
	CreatedAt          time.Time      `bson:"created_at"`
	UpdatedAt          time.Time      `bson:"updated_at"`
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
