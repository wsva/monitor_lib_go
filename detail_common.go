package mlib

import (
	"encoding/json"
	"regexp"
)

/*
MDCommon error in MonitorDetail format
*/
type MDCommon struct {
	//用来存储一些正常的信息
	InfoList []string `json:"InfoList"`

	//由于历史原因，Detail为ok，表示正常，否则存放错误信息
	Detail string `json:"Detail"`
}

// JSONString comment
func (m MDCommon) JSONString() (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// DetailString comment
func (m MDCommon) DetailString() string {
	var result string
	reg := regexp.MustCompile(`[\s\r\n]+$`)
	for _, v := range m.InfoList {
		result += reg.ReplaceAllString(v, "") + "\n"
	}
	return result + m.Detail
}

// WarningString comment
func (m MDCommon) WarningString() string {
	if m.Detail == "" {
		return "ok"
	}
	return m.Detail
}

// GetMDCommonFromJSON comment
func GetMDCommonFromJSON(jsonString string) (*MDCommon, error) {
	var result MDCommon
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMDCommonListFromJSON comment
func GetMDCommonListFromJSON(jsonString string) ([]MDCommon, error) {
	var result []MDCommon
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
