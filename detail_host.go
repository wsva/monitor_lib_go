package mlib

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
CPUUsage comment
Size: number of logic cpu
Used: percent
*/
type CPUUsage struct {
	Number int `json:"Number"`
	Used   int `json:"Used"`
}

/*
MemoryUsage : details of memory usage
Platform: AIX, Linux, Windows
Size: GB
Used: percent

	    AIX: svmon -G, virtual/size, https://www.unixhealthcheck.com/blog?id=255
		Linux: usage(without buffer/cache) on Linux
		Windows: psutil
*/
type MemoryUsage struct {
	Size int `json:"Size"`
	Used int `json:"Used"`
}

/*
SwapUsage comment
Swap on Linux
Paging space on AIX
Virtual memory on Windows
Size: GB
Used: percent
*/
type SwapUsage struct {
	Size int `json:"Size"`
	Used int `json:"Used"`
}

// DiskUsage comment
type DiskUsage struct {
	Name  string `json:"Name"`
	Size  int    `json:"Size"`
	Used  int    `json:"Used"`
	Iused int    `json:"Iused"`
}

/*
Linux

IOWait
iostat -x | sed -n '/avg-cpu/{n;p}' | awk '{print $4}'

IOBusyMax
ioutil -dx | awk '{print $NF}' | grep -P "\d+" | sort -r | head -n1
*/

/*
MDHost MonitorDetail-Host
*/
type MDHost struct {
	Platform  string        `json:"Platform"`
	OS        string        `json:"OS"`
	Hostname  string        `json:"Hostname"`
	CPU       CPUUsage      `json:"CPU"`
	Memory    MemoryUsage   `json:"Memory"`
	Swap      SwapUsage     `json:"Swap"`
	DiskList  []DiskUsage   `json:"DiskList"`
	NTPOffset time.Duration `json:"NTPOffset"` //in seconds
}

// JSONString comment
func (m MDHost) JSONString() (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// DetailString comment
func (m MDHost) DetailString() string {
	var result = ""
	result += fmt.Sprintf("%v,%v,%v;\n", m.Platform, m.OS, m.Hostname)
	result += fmt.Sprintf("CPU:%v,%v%%;\n", m.CPU.Number, m.CPU.Used)
	result += fmt.Sprintf("Mem:%vGB,%v%%;\n", m.Memory.Size, m.Memory.Used)
	result += fmt.Sprintf("Swap:%vGB,%v%%;\n", m.Swap.Size, m.Swap.Used)
	resultDisk := ""
	for _, v := range m.DiskList {
		resultDisk += fmt.Sprintf("Disk:%v,%vGB,%v%%,%v%%;", v.Name, v.Size, v.Used, v.Iused)
	}
	result += "Disk:" + resultDisk + "\n"
	result += fmt.Sprintf("NTPOffset:%v;", (m.NTPOffset * time.Second).String())
	return result
}

// WarningString comment
func (m MDHost) WarningString() string {
	result := ""
	if m.CPU.Used > WarnStdCPU {
		result += fmt.Sprintf("CPU:%v,%v%%;", m.CPU.Number, m.CPU.Used)
	}
	/*
		if m.Memory.Used > WarnStdMemory {
			result += fmt.Sprintf("Mem:%vGB,%v%%;", m.Memory.Size, m.Memory.Used)
		}
	*/
	if m.Swap.Used > WarnStdSwap {
		result += fmt.Sprintf("Swap:%vGB,%v%%;", m.Swap.Size, m.Swap.Used)
	}
	resultDisk := ""
	for _, v := range m.DiskList {
		if v.Iused > WarnStdDisk || v.Used > WarnStdDisk {
			resultDisk += fmt.Sprintf("%v,%vGB,%v%%,%v%%;", v.Name, v.Size, v.Used, v.Iused)
		}
	}
	if resultDisk != "" {
		result += "Disk:" + resultDisk
	}
	if m.NTPOffset*time.Second > WarnStdNTPOffset || m.NTPOffset*time.Second < WarnStdNTPOffset*-1 {
		result += fmt.Sprintf("NTPOffset:%v;", (m.NTPOffset * time.Second).String())
	}
	if result == "" {
		result = "ok"
	}
	return result
}

// GetMDHostFromJSON comment
func GetMDHostFromJSON(jsonString string) (*MDHost, error) {
	var result MDHost
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMDHostFromJSONString comment
func GetMDHostFromJSONString(jsonString string) (*MDHost, error) {
	var result MDHost
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
