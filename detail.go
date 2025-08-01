package mlib

//MD MonitorDetail interface
type MD interface {
	JSONString() (string, error)
	DetailString() string
	WarningString() string
}

func GetMD(mType, detailJSON string) (MD, error) {
	switch mType {
	case "host":
		result, err := GetMDHostFromJSON(detailJSON)
		if err != nil {
			return nil, err
		}
		return *result, err
	case "oracle":
		result, err := GetMDOracleFromJSON(detailJSON)
		if err != nil {
			return nil, err
		}
		return *result, err
	case "oracle_ogg":
		result, err := GetMDOGGFromJSON(detailJSON)
		if err != nil {
			return nil, err
		}
		return *result, err
	case "weblogic_domain":
		result, err := GetMDWeblogicDomainFromJSON(detailJSON)
		if err != nil {
			return nil, err
		}
		return *result, err
	default:
		result, err := GetMDCommonFromJSON(detailJSON)
		if err != nil {
			return nil, err
		}
		return *result, err
	}
}
