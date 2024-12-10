package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	Coll_ByBitUser               string = "ByBitUser"
	Coll_BybitFutureOrderHistory string = "BybitFutureOrderHistory"
	Coll_BybitFutureTradeHistory string = "BybitFutureTradeHistory"
	Coll_BybitFuturePnlHistory   string = "BybitFuturePnlHistory"
	Coll_BybitSpotOrderHistory   string = "BybitSpotOrderHistory"
	Coll_BybitSpotTradelHistory  string = "BybitSpotTradelHistory"
)

type ByBitUser struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ApiKey      string             `bson:"api_key"`
	ApiSecret   string             `bson:"api_secret"`
	Username    string             `bson:"username"`
	Email       string             `bson:"email"`
	PhoneNumber string             `bson:"phone_number"`
	IsTrading   bool               `bson:"is_trading"`
	LastUpdate  time.Time          `bson:"last_update"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type BybitFutureOrderHistory struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	UserID             primitive.ObjectID `bson:"user_id"`
	OrderId            string             `bson:"order_id"`
	OrderLinkId        string             `bson:"order_link_id"`
	BlockTradeId       string             `bson:"block_trade_id"`
	Symbol             string             `bson:"symbol"`
	Price              string             `bson:"price"`
	Qty                string             `bson:"qty"`
	Side               string             `bson:"side"`
	IsLeverage         string             `bson:"is_leverage"`
	PositionIdx        int16              `bson:"position_idx"`
	OrderStatus        string             `bson:"order_status"`
	CancelType         string             `bson:"cancel_type"`
	RejectReason       string             `bson:"reject_reason"`
	AvgPrice           string             `bson:"avg_price"`
	LeavesQty          string             `bson:"leaves_qty"`
	LeavesValue        string             `bson:"leaves_value"`
	CumExecQty         string             `bson:"cum_exec_qty"`
	CumExecValue       string             `bson:"cum_exec_value"`
	CumExecFee         string             `bson:"cum_exec_fee"`
	TimeInForce        string             `bson:"time_in_force"`
	OrderType          string             `bson:"order_type"`
	StopOrderType      string             `bson:"stop_order_type"`
	OrderIv            string             `bson:"order_iv"`
	TriggerPrice       string             `bson:"trigger_price"`
	TakeProfit         string             `bson:"take_profit"`
	StopLoss           string             `bson:"stop_loss"`
	TpTriggerBy        string             `bson:"tp_trigger_by"`
	SlTriggerBy        string             `bson:"sl_trigger_by"`
	TriggerDirection   string             `bson:"trigger_direction"`
	TriggerBy          string             `bson:"trigger_by"`
	LastPriceOnCreated string             `bson:"last_price_on_created"`
	ReduceOnly         bool               `bson:"reduce_only"`
	CloseOnTrigger     bool               `bson:"close_on_trigger"`
	SmpType            string             `bson:"smp_type"`
	SmpGroup           string             `bson:"smp_group"`
	SmpOrderId         string             `bson:"smp_order_id"`
	TpslMode           string             `bson:"tpsl_mode"`
	TpLimitPrice       string             `bson:"tp_limit_price"`
	SlLimitPrice       string             `bson:"sl_limit_price"`
	PlaceType          string             `bson:"place_type"`
	CreatedAt          time.Time          `bson:"created_at"`
	UpdatedAt          time.Time          `bson:"updated_at"`
}
type BybitFutureTradeHistory struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id"`

	Symbol          string    `bson:"symbol"`
	OrderID         string    `bson:"orderId"`
	OrderLinkID     string    `bson:"order_link_id"`
	Side            string    `bson:"side"`
	OrderPrice      string    `bson:"order_price"`
	OrderQty        string    `bson:"order_qty"`
	LeavesQty       string    `bson:"leaves_qty"`
	OrderType       string    `bson:"order_type"`
	StopOrderType   string    `bson:"stop_order_type"`
	ExecFee         string    `bson:"exec_fee"`
	ExecID          string    `bson:"exec_id"`
	ExecPrice       string    `bson:"exec_price"`
	ExecQty         string    `bson:"exec_qty"`
	ExecType        string    `bson:"exec_type"`
	ExecValue       string    `bson:"exec_value"`
	ExecTime        time.Time `bson:"exec_time"`
	IsMaker         bool      `bson:"isMaker"`
	FeeRate         string    `bson:"fee_rate"`
	TradeIv         string    `bson:"trade_iv"`
	MarkIv          string    `bson:"mark_iv"`
	MarkPrice       string    `bson:"mark_price"`
	IndexPrice      string    `bson:"index_price"`
	UnderlyingPrice string    `bson:"underlying_price"`
	BlockTradeID    string    `bson:"block_trade_id"`
	ClosedSize      string    `bson:"closed_size"`
	Seq             int64     `bson:"seq"`
	CreateType      string    `bson:"create_type"`
	CreatedAt       time.Time `bson:"created_at"`
	UpdatedAt       time.Time `bson:"updated_at"`
}
type BybitFuturePnlHistory struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id"`

	Symbol        string    `bson:"symbol"`
	OrderID       string    `bson:"order_id"`
	Side          string    `bson:"side"`
	Qty           string    `bson:"qty"`
	OrderPrice    string    `bson:"order_price"`
	OrderType     string    `bson:"order_type"`
	ExecType      string    `bson:"exec_type"`
	ClosedSize    string    `bson:"closed_size"`
	CumEntryValue string    `bson:"cum_entry_value"`
	AvgEntryPrice string    `bson:"avg_entry_price"`
	CumExitValue  string    `bson:"cum_exit_value"`
	AvgExitPrice  string    `bson:"avg_exit_price"`
	ClosedPnl     string    `bson:"closed_pnl"`
	FillCount     string    `bson:"fill_count"`
	Leverage      string    `bson:"leverage"`
	CreatedTime   time.Time `bson:"created_time"`
	UpdatedTime   time.Time `bson:"updated_time"`
	CreatedAt     time.Time `bson:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at"`
}
type BybitSpotOrderHistory struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id"`

	Symbol             string    `bson:"symbol"`
	OrderType          string    `bson:"order_type"`
	OrderLinkID        string    `bson:"order_link_id"`
	SlLimitPrice       string    `bson:"sl_limit_price"`
	OrderID            string    `bson:"order_id"`
	AvgPrice           string    `bson:"avg_price"`
	CancelType         string    `bson:"cancel_type"`
	StopOrderType      string    `bson:"stop_order_type"`
	LastPriceOnCreated string    `bson:"last_price_on_created"`
	OrderStatus        string    `bson:"order_status"`
	TakeProfit         string    `bson:"take_profit"`
	CumExecValue       string    `bson:"cum_exec_value"`
	SmpType            string    `bson:"smp_type"`
	TriggerDirection   int       `bson:"trigger_direction"`
	BlockTradeID       string    `bson:"block_trade_id"`
	IsLeverage         string    `bson:"is_leverage"`
	RejectReason       string    `bson:"reject_reason"`
	Price              string    `bson:"price"`
	OrderIv            string    `bson:"order_iv"`
	CreatedTime        time.Time `bson:"created_time"`
	PositionIdx        int16     `bson:"position_idx"`
	TpTriggerBy        string    `bson:"tp_trigger_by"`
	TimeInForce        string    `bson:"time_in_force"`
	LeavesValue        string    `bson:"leaves_value"`
	UpdatedTime        time.Time `bson:"updated_time"`
	Side               string    `bson:"side"`
	SmpGroup           int       `bson:"smp_group"`
	TpLimitPrice       string    `bson:"tp_limit_price"`
	TriggerPrice       string    `bson:"trigger_price"`
	CumExecFee         string    `bson:"cum_exec_fee"`
	LeavesQty          string    `bson:"leaves_qty"`
	SlTriggerBy        string    `bson:"sl_trigger_by"`
	CloseOnTrigger     bool      `bson:"close_on_trigger"`
	CumExecQty         string    `bson:"cum_exec_qty"`
	ReduceOnly         bool      `bson:"reduce_only"`
	Qty                string    `bson:"qty"`
	StopLoss           string    `bson:"stop_loss"`
	SmpOrderID         string    `bson:"smp_order_id"`
	TriggerBy          string    `bson:"trigger_by"`
	CreatedAt          time.Time `bson:"created_at"`
	UpdatedAt          time.Time `bson:"updated_at"`
}
type BybitSpotTradelHistory struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id"`

	Symbol          string    `bson:"symbol"`
	OrderType       string    `bson:"order_type"`
	UnderlyingPrice string    `bson:"underlying_price"`
	OrderLinkID     string    `bson:"order_link_id"`
	Side            string    `bson:"side"`
	IndexPrice      string    `bson:"index_price"`
	OrderID         string    `bson:"order_id"`
	StopOrderType   string    `bson:"stop_order_type"`
	LeavesQty       string    `bson:"leaves_qty"`
	ExecTime        time.Time `bson:"exec_time"`
	IsMaker         bool      `bson:"is_maker"`
	ExecFee         string    `bson:"exec_fee"`
	FeeRate         string    `bson:"fee_rate"`
	ExecID          string    `bson:"exec_id"`
	TradeIv         string    `bson:"trade_iv"`
	BlockTradeID    string    `bson:"block_trade_id"`
	MarkPrice       string    `bson:"mark_price"`
	ExecPrice       string    `bson:"exec_price"`
	MarkIv          string    `bson:"mark_iv"`
	OrderQty        string    `bson:"order_qty"`
	OrderPrice      string    `bson:"order_price"`
	ExecValue       string    `bson:"exec_value"`
	ExecType        string    `bson:"exec_type"`
	ExecQty         string    `bson:"exec_qty"`
	CreatedAt       time.Time `bson:"created_at"`
	UpdatedAt       time.Time `bson:"updated_at"`
}

func SelectCollection(collName string) string {
	switch collName {
	case "ByBitUser":
		return Coll_ByBitUser
	case "BybitFutureOrder":
		return Coll_BybitFutureOrderHistory
	case "BybitFutureTrade":
		return Coll_BybitFutureTradeHistory
	case "BybitFuturePnl":
		return Coll_BybitFuturePnlHistory
	case "BybitSpotOrder":
		return Coll_BybitSpotOrderHistory
	case "BybitSpotTradel":
		return Coll_BybitSpotTradelHistory
	default:
		return ""
	}
}
