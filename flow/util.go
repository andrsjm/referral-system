package flow

import (
	"encoding/json"
	"fmt"
)

func EntityStringify(v interface{}) string {
	byteData, err := json.Marshal(v)
	if err != nil {
		fmt.Println("e:jsonMarshal", byteData)
	}

	return string(byteData)
}
