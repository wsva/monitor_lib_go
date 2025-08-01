package mlib

import (
	"encoding/json"
	"fmt"
)

/*
ArchiveLog oracle
Size: GB
Used: percent
*/
type ArchiveLog struct {
	Size int `json:"Size"`
	Used int `json:"Used"`
}

/*
ASM oracle
Size: GB
*/
type ASM struct {
	Name string `json:"Name"`
	Size int    `json:"Size"`
	Used int    `json:"Used"`
}

/*
TableSpace oracle
Size: GB
*/
type TableSpace struct {
	Name string `json:"Name"`
	Size int    `json:"Size"`
	Used int    `json:"Used"`
}

// TableLock comment
type TableLock struct {
	Name     string `json:"Name"`
	Username string `json:"Username"`
	Count    int    `json:"Count"`
}

// PasswordExpire oracle
type PasswordExpire struct {
	Username string `json:"Username"`
}

/*
MDOracle : monitor detail of oracle database
ConnectivityOK: execute a sql through sqlplus
ListenerOK: connenct through sqlplus a/b@ip/c and check error type
*/
type MDOracle struct {
	ConnectivityOK     bool         `json:"ConnectivityOK"`
	ArchiveLogExist    bool         `json:"ArchiveLogExist"`
	ArchiveLog         ArchiveLog   `json:"ArchiveLog"`
	ASMExist           bool         `json:"ASMExist"`
	ASMList            []ASM        `json:"ASMList"`
	TableSpaceList     []TableSpace `json:"TableSpaceList"`
	TableLockList      []TableLock  `json:"TableLockList"`
	PasswordExpireList []string     `json:"PasswordExpireList"`
	ErrorString        string       `json:"ErrorString"`
}

// JSONString comment
func (m MDOracle) JSONString() (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// DetailString comment
func (m MDOracle) DetailString() string {
	var result = ""
	if m.ConnectivityOK {
		result += "Connectivity:success;\n"
	} else {
		result += "Connectivity:fail;\n"
	}
	if m.ArchiveLogExist {
		result += fmt.Sprintf("ArchiveLog:%vGB,%v%%\n",
			m.ArchiveLog.Size, m.ArchiveLog.Used)
	} else {
		result += "ArchiveLog:not exist;\n"
	}
	if m.ASMExist {
		result += "ASM:"
		for _, v := range m.ASMList {
			result += fmt.Sprintf("%v,%vGB,%v%%;",
				v.Name, v.Size, v.Used)
		}
		result += "\n"
	} else {
		result += "ASM:not exist;\n"
	}
	result += "TableSpace:"
	if m.TableSpaceList != nil {
		for _, v := range m.TableSpaceList {
			result += fmt.Sprintf("%v,%vGB,%v%%;",
				v.Name, v.Size, v.Used)
		}
		result += "\n"
	} else {
		result += "not exist;\n"
	}
	if m.TableLockList != nil {
		tableLockMap := make(map[string]string)
		tableLockCountMap := make(map[string]int)
		for _, v := range m.TableLockList {
			tableLockMap[v.Username] += fmt.Sprintf("%v,%v;", v.Name, v.Count)
			tableLockCountMap[v.Username] += v.Count
		}
		for k := range tableLockMap {
			result += fmt.Sprintf("TableLock:%v,%v:%v;\n",
				k, tableLockCountMap[k], tableLockMap[k])
		}
	} else {
		result += "TableLock:not exist;\n"
	}
	result += "PasswordExpire:"
	if m.PasswordExpireList != nil {
		for _, v := range m.PasswordExpireList {
			result += fmt.Sprintf("%v,", v)
		}
		result += ";"
	} else {
		result += "not exist;"
	}
	return result
}

// WarningString comment
func (m MDOracle) WarningString() string {
	result := ""
	if !m.ConnectivityOK {
		result += "Connectivity:fail;\n"
	}
	if m.ArchiveLogExist && m.ArchiveLog.Used > WarnStdArchiveLog {
		result += fmt.Sprintf("ArchiveLog:%vGB,%v%%;\n",
			m.ArchiveLog.Size, m.ArchiveLog.Used)
	}
	if m.ASMExist {
		beginAdded := false
		endAdded := false
		for _, v := range m.ASMList {
			if v.Used > WarnStdASM {
				if !beginAdded {
					beginAdded = true
					result += "ASM:"
				}
				result += fmt.Sprintf("%v,%vGB,%v%%;",
					v.Name, v.Size, v.Used)
			}
		}
		if beginAdded && !endAdded {
			result += ";\n"
		}
	}
	if m.TableSpaceList != nil {
		beginAdded := false
		endAdded := false
		for _, v := range m.TableSpaceList {
			if v.Used > WarnStdTableSpace {
				if !beginAdded {
					beginAdded = true
					result += "TableSpace:"
				}
				result += fmt.Sprintf("%v,%vGB,%v%%;",
					v.Name, v.Size, v.Used)
			}
		}
		if beginAdded && !endAdded {
			result += ";\n"
		}
	}
	if m.TableLockList != nil {
		tableLockCountMap := make(map[string]int)
		for _, v := range m.TableLockList {
			tableLockCountMap[v.Username] += v.Count
		}
		result += "TableLock:"
		for k, v := range tableLockCountMap {
			result += fmt.Sprintf("%v,%v;", k, v)
		}
		result += ";\n"
	}
	if m.PasswordExpireList != nil {
		result += "PasswordExpire:"
		for _, v := range m.PasswordExpireList {
			result += fmt.Sprintf("%v,", v)
		}
		result += ";\n"
	}
	if m.ErrorString != "" {
		result += m.ErrorString + ";"
	}
	if result == "" {
		result = "ok"
	}
	return result
}

// GetMDOracleromJSON comment
func GetMDOracleFromJSON(jsonString string) (*MDOracle, error) {
	var result MDOracle
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMDOracleListFromJSON comment
func GetMDOracleListFromJSON(jsonString string) ([]MDOracle, error) {
	var result []MDOracle
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
