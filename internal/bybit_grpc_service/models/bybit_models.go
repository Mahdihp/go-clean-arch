package models_grpc

import (
	"time"
)

const (
	Coll_ByBitMarketGetInstrumentsInfoLinear  string = "ByBitMarketGetInstrumentsInfoLinear"
	Coll_ByBitMarketGetInstrumentsInfoInverse string = "ByBitMarketGetInstrumentsInfoInverse"
	Coll_ByBitMarketGetInstrumentsInfoSpot    string = "ByBitMarketGetInstrumentsInfoSpot"
)

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
type RiskParameters struct {
	LimitParameter  string `bson:"limitParameter"`
	MarketParameter string `bson:"marketParameter"`
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
	PreListingInfo     any            `json:"preListingInfo"`
	NextPageCursor     string         `json:"nextPageCursor"`
	CreatedAt          time.Time      `bson:"created_at"`
	UpdatedAt          time.Time      `bson:"updated_at"`
}
