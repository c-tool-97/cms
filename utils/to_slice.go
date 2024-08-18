package utils

import "strings"

func StringToInterfaceSlice(in string) []interface{} {
	resp := []interface{}{}
	split := strings.Split(in, ",")
	for _, val := range split {
		if len(val) <= 0 {
			continue
		}
		resp = append(resp, val)
	}
	return resp
}
