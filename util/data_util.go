package util

import "fmt"

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
