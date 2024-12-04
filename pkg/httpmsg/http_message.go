package httpmsg

import "time"

type Message string

const (
	UserNotFound    Message = "User Not Found"
	HistoryNotFound Message = "History Not Found"
)

type BaseMessage struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	Result     string      `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

func NewMassage(msg Message) BaseMessage {
	r := BaseMessage{}
	r.RetMsg = string(msg)
	r.Time = time.Now().Unix()
	return r
}

func NewStrMassage(msg string) BaseMessage {
	r := BaseMessage{}
	r.RetMsg = msg
	r.Time = time.Now().Unix()
	return r
}
