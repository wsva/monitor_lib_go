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

type MTConfig struct {
	MTList []MT `json:"MTList"`
}

func (m *MTConfig) LoadFromJSON(jsonBytes []byte) error {
	var result []MT
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return err
	}
	m.MTList = result
	return nil
}

func (m *MTConfig) LoadFromFile(filename string) error {
	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = m.LoadFromJSON(contentBytes)
	return err
}
