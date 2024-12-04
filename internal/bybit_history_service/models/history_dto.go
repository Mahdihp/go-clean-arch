package models

type HistoryDto struct {
	ApiKey      string `json:"api_key"`
	Category    string `json:"category"`
	Symbol      string `json:"symbol"`
	BaseCoin    string `json:"base_coin"`
	SettleCoin  string `json:"settle_coin"`
	OrderId     string `json:"order_id"`
	OrderLinkId string `json:"order_link_id"`
	OrderFilter string `json:"order_filter"`
	OrderStatus string `json:"order_status"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	Limit       int8   `json:"limit"`
	Cursor      string `json:"cursor"`
}
