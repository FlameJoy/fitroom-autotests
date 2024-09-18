package utils

import (
	"encoding/json"
	"fmt"
)

func URLFormat(m, s string) string {
	str := fmt.Sprintf("%s - %s", m, s)
	if len(str) < 90 {
		for len(str) < 90 {
			str += " "
		}
	} else if len(s) > 90 {
		str = str[0:90-3] + "..."
	} else {
		return str
	}
	return str
}

func PrepareReqBody(v any) []byte {
	if v == nil {
		return nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
