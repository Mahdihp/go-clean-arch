package util

import (
	"encoding/base64"
	"fmt"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
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
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(timeFormat, timeString)

	return t, err
}

// EncodeCursor will encode cursor from mysql to bybit_ws
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
