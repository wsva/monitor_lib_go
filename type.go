package mlib

import (
	"encoding/json"
	"fmt"
	"os"
)

type MType struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

func (t *MType) GetZabbixKey() string {
	return GetZabbixKey(t.ID)
}

func LoadMTypeListFromJSON(jsonBytes []byte) ([]MType, error) {
	var result []MType
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func LoadMTypeListFromFile(filename string) ([]MType, error) {
	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return LoadMTypeListFromJSON(contentBytes)
}

func GetZabbixKey(mtypeID string) string {
	return fmt.Sprintf("gm_%v", mtypeID)
}
