package utils

import "encoding/json"

func ConvertToJsonString(in interface{}) string {
	marshal, _ := json.Marshal(in)
	return string(marshal)
}
