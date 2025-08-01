package mlib

import (
	"encoding/json"
	"fmt"
	"strings"
)

/*
JDBC weblogic JDBCConnectionPoolRuntime
MaxCapacity: 最大容量
MinCapacity: 初始容量
CurrentCapacity: 当前容量
ActiveCount: 当前活动连接总数
ActiveHighCount 最大活动连接数
WaitHighCount: 最大等待连接数
*/
type JDBC struct {
	Name            string `json:"Name"`
	MaxCapacity     int    `json:"MaxCapacity"`
	MinCapacity     int    `json:"MinCapacity"`
	CurrentCapacity int    `json:"CurrentCapacity"`
	ActiveCount     int    `json:"ActiveCount"`
	ActiveHighCount int    `json:"ActiveHighCount"`
	WaitHighCount   int    `json:"WaitHighCount"`
}

/*
Weblogic a weblogic server
JVMHeapSize: MB
JVMHeapUsed: percent
*/
type Weblogic struct {
	ServerName    string `json:"ServerName"`
	RunningState  string `json:"RunningState"`
	HealthState   string `json:"HealthState"`
	JVMHeapSize   string `json:"JVMHeapSize"`
	JVMHeapUsed   int    `json:"JVMHeapUsed"`
	JMSConnection int    `json:"JMSConnection"`
	QueueLength   int    `json:"QueueLength"`
	Hogging       int    `json:"Hogging"`
	Stuck         int    `json:"Stuck"`
}

/*
MDWeblogicDomain : monitor detail of weblogic domain
JDBCList: JDBCConnectionPoolRuntimeList
*/
type MDWeblogicDomain struct {
	WeblogicList []Weblogic `json:"WeblogicList"`
	JDBCList     []JDBC     `json:"JDBCList"`
	ErrorString  string     `json:"ErrorString"`
}

// DetailString comment
func (m MDWeblogicDomain) DetailString() string {
	var result = ""
	result += "WeblogicList:"
	if m.WeblogicList != nil {
		result += "\n"
		for _, v := range m.WeblogicList {
			result += fmt.Sprintf("%v:RunningState %v,HealthState %v,JVM:%v,%v%%,QueueLength %v,Hogging %v,Stuck %v;\n",
				v.ServerName, v.RunningState, v.HealthState, v.JVMHeapSize,
				v.JVMHeapUsed, v.QueueLength, v.Hogging, v.Stuck)
		}
	} else {
		result += "empty;\n"
	}
	result += "JDBCList:"
	if m.JDBCList != nil {
		result += "\n"
		for _, v := range m.JDBCList {
			result += fmt.Sprintf("%v:MaxCapacity %v,MinCapacity %v,CurrentCapacity %v,ActiveCount %v,ActiveHighCount %v,WaitHighCount %v;\n",
				v.Name, v.MaxCapacity, v.MinCapacity, v.CurrentCapacity, v.ActiveCount,
				v.ActiveHighCount, v.WaitHighCount)
		}
	} else {
		result += "empty;\n"
	}
	return result
}

// JSONString comment
func (m MDWeblogicDomain) JSONString() (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// WarningString comment
func (m MDWeblogicDomain) WarningString() string {
	result := ""
	for _, v := range m.WeblogicList {
		var resultTemp string
		if !strings.Contains(v.RunningState, "RUNNING") {
			resultTemp += fmt.Sprintf("State:%v;", v.RunningState)
		}
		if !strings.Contains(v.HealthState, "HEALTH_OK") {
			resultTemp += fmt.Sprintf("Health:%v;", v.HealthState)
		}
		if v.JVMHeapUsed > WarnSTDJVMHeapUsed {
			resultTemp += fmt.Sprintf("JVM:%v,%v%%;", v.JVMHeapSize, v.JVMHeapUsed)
		}
		if v.QueueLength > WarnSTDQueueLength {
			resultTemp += fmt.Sprintf("Queue:%v;", v.QueueLength)
		}
		if v.Hogging > 0 {
			resultTemp += fmt.Sprintf("Hogging:%v;", v.Hogging)
		}
		if v.Stuck > 0 {
			resultTemp += fmt.Sprintf("Stuck:%v;", v.Stuck)
		}
		if resultTemp != "" {
			result += fmt.Sprintf("%v: %v\n", v.ServerName, resultTemp)
		}
	}
	if m.ErrorString != "" {
		result += m.ErrorString + ";"
	}
	if result == "" {
		result = "ok"
	}
	return result
}

// GetMDWeblogicDomainFromJSON comment
func GetMDWeblogicDomainFromJSON(jsonString string) (*MDWeblogicDomain, error) {
	var result MDWeblogicDomain
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMDWeblogicDomainListFromJSON comment
func GetMDWeblogicDomainListFromJSON(jsonString string) ([]MDWeblogicDomain, error) {
	var result []MDWeblogicDomain
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
