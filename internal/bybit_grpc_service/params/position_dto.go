package params_bybit_grpc

import (
	"encoding/json"
	bybit "github.com/wuhewuhe/bybit.go.api"
	"go-clean-arch/adapter/grpc-proto/order"
	"go-clean-arch/adapter/grpc-proto/position"
	"go-clean-arch/util"
	"strconv"
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

type OrderDto struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		OrderID     string `json:"orderId"`
		OrderLinkID string `json:"orderLinkId"`
	} `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

type CancelAllDto struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []struct {
			OrderID     string `json:"orderId"`
			OrderLinkID string `json:"orderLinkId"`
		} `json:"list"`
		Success string `json:"success"`
	} `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

func OrderToCancelAllOrderDto(data *bybit.ServerResponse) CancelAllDto {
	marshal, err := json.Marshal(data)
	var pl CancelAllDto
	if err != nil {
		return pl
	}
	err = json.Unmarshal(marshal, &pl)
	if err != nil {
		return pl
	}
	return pl
}
func OrderToOrderDto(data *bybit.ServerResponse) OrderDto {
	marshal, err := json.Marshal(data)
	var pl OrderDto
	if err != nil {
		return pl
	}
	err = json.Unmarshal(marshal, &pl)
	if err != nil {
		return pl
	}
	return pl
}
func OrderDtoToCancelAllResponse(data CancelAllDto) order.CancelAllResponse {
	response := order.CancelAllResponse{}
	response.RetMsg = data.RetMsg
	response.RetCode = int32(data.RetCode)
	response.RetExtInfo = &order.CancelAllResponse_RetExtInfo{}
	orderidList := make([]*order.OrderId, len(data.Result.List))
	response.Result = &order.CancelAllResponse_Result{
		List: orderidList,
	}
	response.Time = data.Time
	return response
}
func OrderDtoToPlaceOrderResponse(data OrderDto) order.PlaceOrderResponse {
	response := order.PlaceOrderResponse{}
	response.OrderId = data.Result.OrderID
	response.OrderLinkId = data.Result.OrderLinkID
	response.RetCode = strconv.Itoa(data.RetCode)
	response.RetMsg = data.RetMsg
	response.RetExtInfo = util.InterfaceToString(data.RetExtInfo)
	response.Time = data.Time
	return response
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
