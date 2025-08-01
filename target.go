package mlib

import (
	"encoding/json"
	"os"
	"time"
)

// MonitorTarget
type MT struct {
	TypeID string `json:"TypeID"`
	Name   string `json:"Name"`
	IP     string `json:"IP"`
	Port   string `json:"Port"`
}

func (m *MT) GetMRList(timeout time.Duration) []MR {
	resp, err := ZabbixGet(m.IP+":"+m.Port, GetZabbixKey(m.TypeID), timeout*time.Second)
	if err != nil {
		return []MR{NewMR(m.Name, m.IP, m.TypeID, "", err.Error())}
	}
	return ParseMRListFromMessage(m.Name, m.IP, m.TypeID, resp)
}

func LoadMTListFromJSON(jsonBytes []byte) ([]MT, error) {
	var result []MT
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func LoadMTListFromFile(filename string) ([]MT, error) {
	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return LoadMTListFromJSON(contentBytes)
}
