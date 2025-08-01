package mlib

import (
	"encoding/json"
	"fmt"
)

func SplitJSONArray(src string) ([]string, error) {
	var list []any
	err := json.Unmarshal([]byte(src), &list)
	if err != nil {
		return nil, fmt.Errorf("unmarshal to list error: %v", err)
	}

	var result []string
	for _, v := range list {
		b, _ := json.Marshal(v)
		result = append(result, string(b))
	}

	return result, nil
}
