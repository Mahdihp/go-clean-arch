package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05" // reduce precision from RFC3339Nano as date format
)

func MapToString(data map[string]string) string {
	var str string
	for key, value := range data {
		str += fmt.Sprintf("%s: %s ,", key, value)
	}
	return str
}

func InterfaceToString(data interface{}) string {
	return fmt.Sprintf("%v", data)
}

// DecodeCursor will decode cursor from bybit_ws for mysql
func DecodeCursor(encodedTime string) (time.Time, error) {
	//byt, err := base64.StdEncoding.DecodeString(encodedTime)
	//if err != nil {
	//	return time.Time{}, err
	//}
	//
	//timeString := string(byt)
	t, err := time.Parse(timeFormat, encodedTime)
	return t, err
}

// EncodeCursor will encode cursor from mysql to bybit_ws
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)
	return base64.StdEncoding.EncodeToString([]byte(timeString))
}

func ConvertInterfaceToAny(v interface{}) (*anypb.Any, error) {
	anyValue := &anypb.Any{}
	bytes, _ := json.Marshal(v)
	bytesValue := &wrappers.BytesValue{
		Value: bytes,
	}
	err := anypb.MarshalFrom(anyValue, bytesValue, proto.MarshalOptions{})
	return anyValue, err
}

func ConvertAnyToInterface(anyValue *anypb.Any) (interface{}, error) {
	var value interface{}
	bytesValue := &wrappers.BytesValue{}
	err := anypb.UnmarshalTo(anyValue, bytesValue, proto.UnmarshalOptions{})
	if err != nil {
		return value, err
	}
	uErr := json.Unmarshal(bytesValue.Value, &value)
	if uErr != nil {
		return value, uErr
	}
	return value, nil
}
