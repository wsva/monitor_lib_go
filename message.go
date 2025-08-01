package mlib

import (
	"strings"
)

const (
	MessageTypeMDSingle = "MonitorDetailSingle"
	MessageTypeMDList   = "MonitorDetailList"
	MessageTypeMRSingle = "MonitorResultSingle"
	MessageTypeMRList   = "MonitorResultList"
)

func ParseMRListFromMessage(name, address, monitorType, message string) []MR {
	//MDSingle
	if strings.Contains(message, MessageTypeMDSingle) {
		message = strings.ReplaceAll(message, MessageTypeMDSingle, "")
		return []MR{NewMR(name, address, monitorType, message, "")}
	}

	//MDList
	if strings.Contains(message, MessageTypeMDList) {
		message = strings.ReplaceAll(message, MessageTypeMDList, "")
		mdList, err := SplitJSONArray(message)
		if err != nil {
			return []MR{NewMR(name, address, monitorType, "", err.Error()+";"+message)}
		}
		var mrList []MR
		for _, v := range mdList {
			mrList = append(mrList, NewMR(name, address, monitorType, v, ""))
		}
		return mrList
	}

	//MRSingle
	if strings.Contains(message, MessageTypeMRSingle) {
		message = strings.ReplaceAll(message, MessageTypeMRSingle, "")
		mr, err := GetMRFromJSON([]byte(message))
		if err != nil {
			return []MR{NewMR(name, address, monitorType, "", err.Error()+";"+message)}
		}
		return []MR{*mr}
	}

	//MRList
	if strings.Contains(message, MessageTypeMRList) {
		message = strings.ReplaceAll(message, MessageTypeMRList, "")
		mrList, err := GetMRListFromJSON([]byte(message))
		if err != nil {
			return []MR{NewMR(name, address, monitorType, "", err.Error()+";"+message)}
		}
		return mrList
	}

	return []MR{NewMR(name, address, monitorType, "", "unknown message type: "+message)}
}
