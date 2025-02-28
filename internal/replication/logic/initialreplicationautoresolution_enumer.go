// Code generated by "enumer -type=InitialReplicationAutoResolution -trimprefix=InitialReplicationAutoResolution"; DO NOT EDIT.

package logic

import (
	"fmt"
)

const (
	_InitialReplicationAutoResolutionName_0 = "MostRecentAll"
	_InitialReplicationAutoResolutionName_1 = "Fail"
)

var (
	_InitialReplicationAutoResolutionIndex_0 = [...]uint8{0, 10, 13}
	_InitialReplicationAutoResolutionIndex_1 = [...]uint8{0, 4}
)

func (i InitialReplicationAutoResolution) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _InitialReplicationAutoResolutionName_0[_InitialReplicationAutoResolutionIndex_0[i]:_InitialReplicationAutoResolutionIndex_0[i+1]]
	case i == 4:
		return _InitialReplicationAutoResolutionName_1
	default:
		return fmt.Sprintf("InitialReplicationAutoResolution(%d)", i)
	}
}

var _InitialReplicationAutoResolutionValues = []InitialReplicationAutoResolution{1, 2, 4}

var _InitialReplicationAutoResolutionNameToValueMap = map[string]InitialReplicationAutoResolution{
	_InitialReplicationAutoResolutionName_0[0:10]:  1,
	_InitialReplicationAutoResolutionName_0[10:13]: 2,
	_InitialReplicationAutoResolutionName_1[0:4]:   4,
}

// InitialReplicationAutoResolutionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func InitialReplicationAutoResolutionString(s string) (InitialReplicationAutoResolution, error) {
	if val, ok := _InitialReplicationAutoResolutionNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to InitialReplicationAutoResolution values", s)
}

// InitialReplicationAutoResolutionValues returns all values of the enum
func InitialReplicationAutoResolutionValues() []InitialReplicationAutoResolution {
	return _InitialReplicationAutoResolutionValues
}

// IsAInitialReplicationAutoResolution returns "true" if the value is listed in the enum definition. "false" otherwise
func (i InitialReplicationAutoResolution) IsAInitialReplicationAutoResolution() bool {
	for _, v := range _InitialReplicationAutoResolutionValues {
		if i == v {
			return true
		}
	}
	return false
}
