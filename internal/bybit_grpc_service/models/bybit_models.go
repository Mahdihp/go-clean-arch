package models_grpc

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	Coll_ByBitMarketGetInstrumentsInfoLinear string = "ByBitMarketGetInstrumentsInfoLinear"
	Coll_ByBitMarketGetInstrumentsInfoOption string = "ByBitMarketGetInstrumentsInfoOption"
	Coll_ByBitMarketGetInstrumentsInfoSpot   string = "ByBitMarketGetInstrumentsInfoSpot"
)

type ByBitMarketGetInstrumentsInfoLinear struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Symbol          string `bson:"symbol"`
	ContractType    string `bson:"contractType"`
	Status          string `bson:"status"`
	BaseCoin        string `bson:"baseCoin"`
	QuoteCoin       string `bson:"quoteCoin"`
	LaunchTime      string `bson:"launchTime"`
	DeliveryTime    string `bson:"deliveryTime"`
	DeliveryFeeRate string `bson:"deliveryFeeRate"`
	PriceScale      string `bson:"priceScale"`
	LeverageFilter  struct {
		MinLeverage  string `bson:"minLeverage"`
		MaxLeverage  string `bson:"maxLeverage"`
		LeverageStep string `bson:"leverageStep"`
	} `bson:"leverageFilter"`
	PriceFilter struct {
		MinPrice string `bson:"minPrice"`
		MaxPrice string `bson:"maxPrice"`
		TickSize string `bson:"tickSize"`
	} `bson:"priceFilter"`
	LotSizeFilter struct {
		MaxOrderQty         string `bson:"maxOrderQty"`
		MaxMktOrderQty      string `bson:"maxMktOrderQty"`
		MinOrderQty         string `bson:"minOrderQty"`
		QtyStep             string `bson:"qtyStep"`
		PostOnlyMaxOrderQty string `bson:"postOnlyMaxOrderQty"`
		MinNotionalValue    string `bson:"minNotionalValue"`
	} `bson:"lotSizeFilter"`
	UnifiedMarginTrade bool   `bson:"unifiedMarginTrade"`
	FundingInterval    int    `bson:"fundingInterval"`
	SettleCoin         string `bson:"settleCoin"`
	CopyTrading        string `bson:"copyTrading"`
	UpperFundingRate   string `bson:"upperFundingRate"`
	LowerFundingRate   string `bson:"lowerFundingRate"`
	IsPreListing       bool   `bson:"isPreListing"`
	PreListingInfo     struct {
		CurAuctionPhase string `bson:"curAuctionPhase"`
		Phases          []struct {
			Phase     string `bson:"phase"`
			StartTime string `bson:"startTime"`
			EndTime   string `bson:"endTime"`
		} `bson:"phases"`
		AuctionFeeInfo struct {
			AuctionFeeRate string `bson:"auctionFeeRate"`
			TakerFeeRate   string `bson:"takerFeeRate"`
			MakerFeeRate   string `bson:"makerFeeRate"`
		} `bson:"auctionFeeInfo"`
	} `bson:"preListingInfo"`

	NextPageCursor string `bson:"nextPageCursor"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
type ByBitMarketGetInstrumentsInfoOption struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	NextPageCursor  string `bson:"nextPageCursor"`
	Symbol          string `bson:"symbol"`
	Status          string `bson:"status"`
	BaseCoin        string `bson:"baseCoin"`
	QuoteCoin       string `bson:"quoteCoin"`
	SettleCoin      string `bson:"settleCoin"`
	OptionsType     string `bson:"optionsType"`
	LaunchTime      string `bson:"launchTime"`
	DeliveryTime    string `bson:"deliveryTime"`
	DeliveryFeeRate string `bson:"deliveryFeeRate"`
	PriceFilter     struct {
		MinPrice string `bson:"minPrice"`
		MaxPrice string `bson:"maxPrice"`
		TickSize string `bson:"tickSize"`
	} `bson:"priceFilter"`
	LotSizeFilter struct {
		MaxOrderQty string `bson:"maxOrderQty"`
		MinOrderQty string `bson:"minOrderQty"`
		QtyStep     string `bson:"qtyStep"`
	} `bson:"lotSizeFilter"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
type ByBitMarketGetInstrumentsInfoSpot struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Symbol        string `bson:"symbol"`
	BaseCoin      string `bson:"baseCoin"`
	QuoteCoin     string `bson:"quoteCoin"`
	Innovation    string `bson:"innovation"`
	Status        string `bson:"status"`
	MarginTrading string `bson:"marginTrading"`
	StTag         string `bson:"stTag"`
	LotSizeFilter struct {
		BasePrecision  string `bson:"basePrecision"`
		QuotePrecision string `bson:"quotePrecision"`
		MinOrderQty    string `bson:"minOrderQty"`
		MaxOrderQty    string `bson:"maxOrderQty"`
		MinOrderAmt    string `bson:"minOrderAmt"`
		MaxOrderAmt    string `bson:"maxOrderAmt"`
	} `bson:"lotSizeFilter"`
	PriceFilter struct {
		TickSize string `bson:"tickSize"`
	} `bson:"priceFilter"`
	RiskParameters struct {
		LimitParameter  string `bson:"limitParameter"`
		MarketParameter string `bson:"marketParameter"`
	} `bson:"riskParameters"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
