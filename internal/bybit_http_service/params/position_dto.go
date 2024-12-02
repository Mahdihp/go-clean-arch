package params_bybit_http

import (
	"encoding/json"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/position"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

type PositionList struct {
	RetCode    int             `json:"retCode"`
	Time       int64           `json:"time"`
	RetMsg     string          `json:"retMsg"`
	Result     PositionListDto `json:"result"`
	RetExtInfo RetExtInfo
}

type RetExtInfo struct {
	RetExtInfo string `json:"retExtInfo"`
}
type PositionListDto struct {
	Category       string        `json:"category"`
	List           []PositionDto `json:"list"`
	NextPageCursor string        `json:"nextPageCursor"`
}
type PositionDto struct {
	positionIdx    int32  `json:"positionIdx"`
	riskIdc        int32  `json:"riskIdc"`
	tradeMode      int32  `json:"tradeMode"`
	autoAddMargin  int32  `json:"autoAddMargin"`
	riskLimitValue string `json:"riskLimitValue"`
	symbol         string `json:"symbol"`
	side           string `json:"side"`
	size           string `json:"size"`
	avgPrice       string `json:"avgPrice"`
	positionValue  string `json:"positionValue"`
	positionStatus string `json:"positionStatus"`
	leverage       string `json:"leverage"`
	markPrice      string `json:"markPrice"`
}

func StringToPositionList(string *bybit.ServerResponse) PositionList {
	marshal, err := json.Marshal(string)
	var pl PositionList
	if err != nil {
		return pl
	}
	err = json.Unmarshal(marshal, &pl)
	if err != nil {
		return pl
	}
	return pl
}
func PositionListToDataList(data PositionList) position.PositionInfoResponse {
	response := position.PositionInfoResponse{}
	var lists []*position.PositionList
	for i, v := range data.Result.List {
		lists[i] = &position.PositionList{
			PositionIdx:    v.positionIdx,
			RiskId:         v.riskIdc,
			TradeMode:      v.tradeMode,
			AutoAddMargin:  v.autoAddMargin,
			RiskLimitValue: v.riskLimitValue,
			Symbol:         v.symbol,
			Side:           v.side,
			Size:           v.size,
			AvgPrice:       v.avgPrice,
			PositionValue:  v.positionValue,
			PositionStatus: v.positionStatus,
			Leverage:       v.leverage,
			MarkPrice:      v.markPrice,
		}
	}
	response.List = lists
	response.RetMsg = data.RetMsg
	response.Category = data.Result.Category
	response.NextPageCursor = data.Result.NextPageCursor
	return response
}
