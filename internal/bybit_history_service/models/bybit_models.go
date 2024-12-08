package models

import (
	"time"
)

type ByBitUser struct {
	ID                      int64                     `gorm:"not null;" json:"not null"`
	ApiKey                  string                    `gorm:"type:varchar;not null;unique" json:"api_key"`
	ApiSecret               string                    `gorm:"type:varchar" json:"api_secret"`
	Username                string                    `gorm:"type:varchar" json:"username;unique"`
	Email                   string                    `gorm:"type:varchar" json:"email"`
	PhoneNumber             string                    `gorm:"type:varchar" json:"phone_number;unique"`
	IsTrading               bool                      `gorm:"type:bool" json:"is_trading"`
	LastUpdate              time.Time                 `gorm:"type:timestamp" json:"last_update"`
	CreatedAt               time.Time                 `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt               time.Time                 `gorm:"type:timestamp" json:"updated_at"`
	BybitFutureOrderHistory []BybitFutureOrderHistory `gorm:"foreignKey:UserID"`
	BybitFutureTradeHistory []BybitFutureTradeHistory `gorm:"foreignKey:UserID"`
	BybitFuturePnlHistory   []BybitFuturePnlHistory   `gorm:"foreignKey:UserID"`
	BybitSpotOrderHistory   []BybitSpotOrderHistory   `gorm:"foreignKey:UserID"`
	BybitSpotTradelHistory  []BybitSpotTradelHistory  `gorm:"foreignKey:UserID"`
}

type BybitFutureOrderHistory struct {
	ID                 int64     `gorm:"not null;" json:"id"`
	UserID             int64     `gorm:"not null;" json:"user_id"`
	OrderId            string    `gorm:"type:varchar" json:"order_id" binding:"required"`
	OrderLinkId        string    `gorm:"type:varchar" json:"order_link_id"`
	BlockTradeId       string    `gorm:"type:varchar" json:"block_trade_id"`
	Symbol             string    `gorm:"type:varchar" json:"symbol"`
	Price              string    `gorm:"type:varchar" json:"price"`
	Qty                string    `gorm:"type:varchar" json:"qty"`
	Side               string    `gorm:"type:varchar" json:"side"`
	IsLeverage         string    `gorm:"type:varchar" json:"is_leverage"`
	PositionIdx        int16     `gorm:"type:int4" json:"position_idx"`
	OrderStatus        string    `gorm:"type:varchar" json:"order_status"`
	CancelType         string    `gorm:"type:varchar" json:"cancel_type"`
	RejectReason       string    `gorm:"type:varchar" json:"reject_reason"`
	AvgPrice           string    `gorm:"type:varchar" json:"avg_price"`
	LeavesQty          string    `gorm:"type:varchar" json:"leaves_qty"`
	LeavesValue        string    `gorm:"type:varchar" json:"leaves_value"`
	CumExecQty         string    `gorm:"type:varchar" json:"cum_exec_qty"`
	CumExecValue       string    `gorm:"type:varchar" json:"cum_exec_value"`
	CumExecFee         string    `gorm:"type:varchar" json:"cum_exec_fee"`
	TimeInForce        string    `gorm:"type:varchar" json:"time_in_force"`
	OrderType          string    `gorm:"type:varchar" json:"order_type"`
	StopOrderType      string    `gorm:"type:varchar" json:"stop_order_type"`
	OrderIv            string    `gorm:"type:varchar" json:"order_iv"`
	TriggerPrice       string    `gorm:"type:varchar" json:"trigger_price"`
	TakeProfit         string    `gorm:"type:varchar" json:"take_profit"`
	StopLoss           string    `gorm:"type:varchar" json:"stop_loss"`
	TpTriggerBy        string    `gorm:"type:varchar" json:"tp_trigger_by"`
	SlTriggerBy        string    `gorm:"type:varchar" json:"sl_trigger_by"`
	TriggerDirection   string    `gorm:"type:varchar" json:"trigger_direction"`
	TriggerBy          string    `gorm:"type:varchar" json:"trigger_by"`
	LastPriceOnCreated string    `gorm:"type:varchar" json:"last_price_on_created"`
	ReduceOnly         bool      `gorm:"type:bool" json:"reduce_only"`
	CloseOnTrigger     bool      `gorm:"type:bool" json:"close_on_trigger"`
	SmpType            string    `gorm:"type:varchar" json:"smp_type"`
	SmpGroup           string    `gorm:"type:varchar" json:"smp_group"`
	SmpOrderId         string    `gorm:"type:varchar" json:"smp_order_id"`
	TpslMode           string    `gorm:"type:varchar" json:"tpsl_mode"`
	TpLimitPrice       string    `gorm:"type:varchar" json:"tp_limit_price"`
	SlLimitPrice       string    `gorm:"type:varchar" json:"sl_limit_price"`
	PlaceType          string    `gorm:"type:varchar" json:"place_type"`
	CreatedAt          time.Time `gorm:"type:timestamp" json:"last_update"`
	UpdatedAt          time.Time `gorm:"type:timestamp" json:"updated_at"`
}
type BybitFutureTradeHistory struct {
	ID     int64 `gorm:"not null;" json:"id"`
	UserID int64 `gorm:"not null;" json:"user_id"`

	Symbol          string    `gorm:"type:timestamp" json:"symbol"`
	OrderID         string    `gorm:"type:varchar" json:"orderId"`
	OrderLinkID     string    `gorm:"type:varchar" json:"order_link_id"`
	Side            string    `gorm:"type:varchar" json:"side"`
	OrderPrice      string    `gorm:"type:varchar" json:"order_price"`
	OrderQty        string    `gorm:"type:varchar" json:"order_qty"`
	LeavesQty       string    `gorm:"type:varchar" json:"leaves_qty"`
	OrderType       string    `gorm:"type:varchar" json:"order_type"`
	StopOrderType   string    `gorm:"type:varchar" json:"stop_order_type"`
	ExecFee         string    `gorm:"type:varchar" json:"exec_fee"`
	ExecID          string    `gorm:"type:varchar" json:"exec_id"`
	ExecPrice       string    `gorm:"type:varchar" json:"exec_price"`
	ExecQty         string    `gorm:"type:varchar" json:"exec_qty"`
	ExecType        string    `gorm:"type:varchar" json:"exec_type"`
	ExecValue       string    `gorm:"type:varchar" json:"exec_value"`
	ExecTime        time.Time `gorm:"type:timestamp" json:"exec_time"`
	IsMaker         bool      `gorm:"type:bool" json:"isMaker"`
	FeeRate         string    `gorm:"type:varchar" json:"fee_rate"`
	TradeIv         string    `gorm:"type:varchar" json:"trade_iv"`
	MarkIv          string    `gorm:"type:varchar" json:"mark_iv"`
	MarkPrice       string    `gorm:"type:varchar" json:"mark_price"`
	IndexPrice      string    `gorm:"type:varchar" json:"index_price"`
	UnderlyingPrice string    `gorm:"type:varchar" json:"underlying_price"`
	BlockTradeID    string    `gorm:"type:varchar" json:"block_trade_id"`
	ClosedSize      string    `gorm:"type:varchar" json:"closed_size"`
	Seq             int64     `gorm:"type:int8" json:"seq"`
	CreateType      string    `gorm:"type:varchar" json:"create_type"`
}
type BybitFuturePnlHistory struct {
	ID     int64 `gorm:"not null;" json:"id"`
	UserID int64 `gorm:"not null;" json:"user_id"`

	Symbol        string    `gorm:"type:varchar" json:"symbol"`
	OrderID       string    `gorm:"type:varchar" json:"order_id"`
	Side          string    `gorm:"type:varchar" json:"side"`
	Qty           string    `gorm:"type:varchar" json:"qty"`
	OrderPrice    string    `gorm:"type:varchar" json:"order_price"`
	OrderType     string    `gorm:"type:varchar" json:"order_type"`
	ExecType      string    `gorm:"type:varchar" json:"exec_type"`
	ClosedSize    string    `gorm:"type:varchar" json:"closed_size"`
	CumEntryValue string    `gorm:"type:varchar" json:"cum_entry_value"`
	AvgEntryPrice string    `gorm:"type:varchar" json:"avg_entry_price"`
	CumExitValue  string    `gorm:"type:varchar" json:"cum_exit_value"`
	AvgExitPrice  string    `gorm:"type:varchar" json:"avg_exit_price"`
	ClosedPnl     string    `gorm:"type:varchar" json:"closed_pnl"`
	FillCount     string    `gorm:"type:varchar" json:"fill_count"`
	Leverage      string    `gorm:"type:varchar" json:"leverage"`
	CreatedTime   time.Time `gorm:"type:timestamp" json:"created_time"`
	UpdatedTime   time.Time `gorm:"type:timestamp" json:"updated_time"`
}
type BybitSpotOrderHistory struct {
	ID     int64 `gorm:"not null;" json:"id"`
	UserID int64 `gorm:"not null;" json:"user_id"`

	Symbol             string    `gorm:"type:varchar" json:"symbol"`
	OrderType          string    `gorm:"type:varchar" json:"order_type"`
	OrderLinkID        string    `gorm:"type:varchar" json:"order_link_id"`
	SlLimitPrice       string    `gorm:"type:varchar" json:"sl_limit_price"`
	OrderID            string    `gorm:"type:varchar" json:"order_id"`
	AvgPrice           string    `gorm:"type:varchar" json:"avg_price"`
	CancelType         string    `gorm:"type:varchar" json:"cancel_type"`
	StopOrderType      string    `gorm:"type:varchar" json:"stop_order_type"`
	LastPriceOnCreated string    `gorm:"type:varchar" json:"last_price_on_created"`
	OrderStatus        string    `gorm:"type:varchar" json:"order_status"`
	TakeProfit         string    `gorm:"type:varchar" json:"take_profit"`
	CumExecValue       string    `gorm:"type:varchar" json:"cum_exec_value"`
	SmpType            string    `gorm:"type:varchar" json:"smp_type"`
	TriggerDirection   int       `gorm:"type:int4" json:"trigger_direction"`
	BlockTradeID       string    `gorm:"type:varchar" json:"block_trade_id"`
	IsLeverage         string    `gorm:"type:varchar" json:"is_leverage"`
	RejectReason       string    `gorm:"type:varchar" json:"reject_reason"`
	Price              string    `gorm:"type:varchar" json:"price"`
	OrderIv            string    `gorm:"type:varchar" json:"order_iv"`
	CreatedTime        time.Time `gorm:"type:timestamp" json:"created_time"`
	PositionIdx        int16     `gorm:"type:int4" json:"position_idx"`
	TpTriggerBy        string    `gorm:"type:varchar" json:"tp_trigger_by"`
	TimeInForce        string    `gorm:"type:varchar" json:"time_in_force"`
	LeavesValue        string    `gorm:"type:varchar" json:"leaves_value"`
	UpdatedTime        time.Time `gorm:"type:timestamp" json:"updated_time"`
	Side               string    `gorm:"type:varchar" json:"side"`
	SmpGroup           int       `gorm:"type:int4" json:"smp_group"`
	TpLimitPrice       string    `gorm:"type:varchar" json:"tp_limit_price"`
	TriggerPrice       string    `gorm:"type:varchar" json:"trigger_price"`
	CumExecFee         string    `gorm:"type:varchar" json:"cum_exec_fee"`
	LeavesQty          string    `gorm:"type:varchar" json:"leaves_qty"`
	SlTriggerBy        string    `gorm:"type:varchar" json:"sl_trigger_by"`
	CloseOnTrigger     bool      `gorm:"type:bool" json:"close_on_trigger"`
	CumExecQty         string    `gorm:"type:varchar" json:"cum_exec_qty"`
	ReduceOnly         bool      `gorm:"type:bool" json:"reduce_only"`
	Qty                string    `gorm:"type:varchar" json:"qty"`
	StopLoss           string    `gorm:"type:varchar" json:"stop_loss"`
	SmpOrderID         string    `gorm:"type:varchar" json:"smp_order_id"`
	TriggerBy          string    `gorm:"type:varchar" json:"trigger_by"`
}
type BybitSpotTradelHistory struct {
	ID     int64 `gorm:"not null;" json:"id"`
	UserID int64 `gorm:"not null;" json:"user_id"`

	Symbol          string    `gorm:"type:varchar" json:"symbol"`
	OrderType       string    `gorm:"type:varchar" json:"order_type"`
	UnderlyingPrice string    `gorm:"type:varchar" json:"underlying_price"`
	OrderLinkID     string    `gorm:"type:varchar" json:"order_link_id"`
	Side            string    `gorm:"type:varchar" json:"side"`
	IndexPrice      string    `gorm:"type:varchar" json:"index_price"`
	OrderID         string    `gorm:"type:varchar" json:"order_id"`
	StopOrderType   string    `gorm:"type:varchar" json:"stop_order_type"`
	LeavesQty       string    `gorm:"type:varchar" json:"leaves_qty"`
	ExecTime        time.Time `gorm:"type:timestamp" json:"exec_time"`
	IsMaker         bool      `gorm:"type:bool" json:"is_maker"`
	ExecFee         string    `gorm:"type:varchar" json:"exec_fee"`
	FeeRate         string    `gorm:"type:varchar" json:"fee_rate"`
	ExecID          string    `gorm:"type:varchar" json:"exec_id"`
	TradeIv         string    `gorm:"type:varchar" json:"trade_iv"`
	BlockTradeID    string    `gorm:"type:varchar" json:"block_trade_id"`
	MarkPrice       string    `gorm:"type:varchar" json:"mark_price"`
	ExecPrice       string    `gorm:"type:varchar" json:"exec_price"`
	MarkIv          string    `gorm:"type:varchar" json:"mark_iv"`
	OrderQty        string    `gorm:"type:varchar" json:"order_qty"`
	OrderPrice      string    `gorm:"type:varchar" json:"order_price"`
	ExecValue       string    `gorm:"type:varchar" json:"exec_value"`
	ExecType        string    `gorm:"type:varchar" json:"exec_type"`
	ExecQty         string    `gorm:"type:varchar" json:"exec_qty"`
}
