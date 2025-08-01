package mlib

import (
	"encoding/json"
	"errors"
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

type MTypeConfig struct {
	MTypeList []MType           `json:"MTypeList"`
	MTypeMap  map[string]*MType `json:"MTypeMap"`
	OrderList []string          `json:"OrderList"`
}

func (m *MTypeConfig) Get(id string) (*MType, error) {
	if mtype, ok := m.MTypeMap[id]; ok {
		return mtype, nil
	}
	return nil, errors.New("type not found by id")
}

func (m *MTypeConfig) LoadFromJSON(jsonBytes []byte) error {
	var result []MType
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return err
	}
	m.MTypeList = result
	m.OrderList = make([]string, len(result))
	for k := range result {
		m.MTypeMap[result[k].ID] = &result[k]
		m.OrderList[k] = result[k].ID
	}
	return nil
}

func (m *MTypeConfig) LoadFromFile(filename string) error {
	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = m.LoadFromJSON(contentBytes)
	return err
}

func GetZabbixKey(mtypeID string) string {
	return fmt.Sprintf("gm_%v", mtypeID)
}
