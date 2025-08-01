package mlib

import (
	"encoding/json"
	"os"
	"time"
)

/*
MonitorResult
*/
type MR struct {
	Name        string `json:"Name"`
	Address     string `json:"Address"`
	Timestamp   int64  `json:"Timestamp"` // Unix()
	TimeString  string `json:"TimeString"`
	MonitorType string `json:"MonitorType"`
	ErrorString string `json:"ErrorString"`
	DetailJSON  string `json:"DetailJSON"`
}

func (m *MR) JSONString() (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func (m *MR) GetWarning() string {
	if m.ErrorString != "" {
		return m.ErrorString
	}
	md, err := GetMD(m.MonitorType, m.DetailJSON)
	if err != nil {
		return err.Error()
	}
	return md.WarningString()
}

func NewMR(name, addrress, monitorType, detailJSON, errorString string) MR {
	timeNow := time.Now()
	return MR{
		Name:        name,
		Address:     addrress,
		MonitorType: monitorType,
		Timestamp:   timeNow.Unix(),
		TimeString:  timeNow.Format("2006-01-02 15:04:05"),
		ErrorString: errorString,
		DetailJSON:  detailJSON,
	}
}

func GetMRFromJSON(jsonBytes []byte) (*MR, error) {
	var result MR
	err := json.Unmarshal(jsonBytes, &result)
	return &result, err
}

func GetMRListFromJSON(jsonBytes []byte) ([]MR, error) {
	var result []MR
	err := json.Unmarshal(jsonBytes, &result)
	return result, err
}

func WriteMRListToFile(result []MR, filename string) error {
	contentBytes, err := json.Marshal(result)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, contentBytes, 0666)
	if err != nil {
		return err
	}
	return nil
}

func GetMRListFromFile(filename string) ([]MR, error) {
	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var mrList []MR
	err = json.Unmarshal(contentBytes, &mrList)
	return mrList, err
}
