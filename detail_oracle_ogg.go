package mlib

import (
	"encoding/json"
	"strings"
)

/*
MDOGG comment
*/
type MDOGG struct {
	Directory    string   `json:"Directory"`
	NormalList   []string `json:"NormalList"`
	AbnormalList []string `json:"AbnormalList"`
}

// JSONString comment
func (m MDOGG) JSONString() (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// DetailString comment
func (m MDOGG) DetailString() string {
	result := ""
	result += m.Directory + "\n"
	result += "NormalList:"
	if m.NormalList != nil {
		for _, v := range m.NormalList {
			result += v + ","
		}
	} else {
		result += "empty;"
	}
	result += "\n"
	result += "AbnormalList:"
	if m.NormalList != nil {
		for _, v := range m.AbnormalList {
			result += v + ","
		}
	} else {
		result += "empty;"
	}
	return result
}

// WarningString comment
func (m MDOGG) WarningString() string {
	result := ""
	for _, v := range m.AbnormalList {
		if strings.Contains(v, "RUNNING") &&
			strings.Contains(v, "unknown") &&
			strings.Contains(v, "00:00") {
			continue
		}
		result += v + "\n"
	}
	if result == "" {
		result = "ok"
	} else {
		result = m.Directory + "\n" + result
	}
	return result
}

// GetMDOGGFromJSON comment
func GetMDOGGFromJSON(jsonString string) (*MDOGG, error) {
	var result MDOGG
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMDOGGListFromJSON comment
func GetMDOGGListFromJSON(jsonString string) ([]MDOGG, error) {
	var result []MDOGG
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
