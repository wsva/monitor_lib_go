package mlib

import (
	"regexp"
)

type FilterRegexp struct {
	Key    string `json:"Key"`
	Regexp string `json:"Regexp"`
}

func GetFilterResult(filterList []FilterRegexp, mr MR) bool {
	for _, v := range filterList {
		reg := regexp.MustCompile(v.Regexp)
		switch v.Key {
		case "Name":
			if reg.MatchString(mr.Name) {
				return true
			}
		case "Address":
			if reg.MatchString(mr.Address) {
				return true
			}
		case "MonitorType":
			if reg.MatchString(mr.MonitorType) {
				return true
			}
		case "Warning":
			if reg.MatchString(mr.GetWarning()) {
				return true
			}
		case "NameAndType":
			if reg.MatchString(mr.Name + mr.MonitorType) {
				return true
			}
		case "AddressAndType":
			if reg.MatchString(mr.Address + mr.MonitorType) {
				return true
			}
		}
	}
	return false
}
