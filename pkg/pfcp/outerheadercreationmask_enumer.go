// Code generated by "enumer -type=OuterHeaderCreationMask -yaml"; DO NOT EDIT.

//
package pfcp

import (
	"fmt"
)

const (
	_OuterHeaderCreationMaskName_0 = "OUTER_HEADER_CREATION_GTPU_UDP_IPV4OUTER_HEADER_CREATION_GTPU_UDP_IPV6"
	_OuterHeaderCreationMaskName_1 = "OUTER_HEADER_CREATION_UDP_IPV4"
	_OuterHeaderCreationMaskName_2 = "OUTER_HEADER_CREATION_UDP_IPV6"
	_OuterHeaderCreationMaskName_3 = "OUTER_HEADER_CREATION_IPV4"
	_OuterHeaderCreationMaskName_4 = "OUTER_HEADER_CREATION_IPV6"
)

var (
	_OuterHeaderCreationMaskIndex_0 = [...]uint8{0, 35, 70}
	_OuterHeaderCreationMaskIndex_1 = [...]uint8{0, 30}
	_OuterHeaderCreationMaskIndex_2 = [...]uint8{0, 30}
	_OuterHeaderCreationMaskIndex_3 = [...]uint8{0, 26}
	_OuterHeaderCreationMaskIndex_4 = [...]uint8{0, 26}
)

func (i OuterHeaderCreationMask) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _OuterHeaderCreationMaskName_0[_OuterHeaderCreationMaskIndex_0[i]:_OuterHeaderCreationMaskIndex_0[i+1]]
	case i == 4:
		return _OuterHeaderCreationMaskName_1
	case i == 8:
		return _OuterHeaderCreationMaskName_2
	case i == 16:
		return _OuterHeaderCreationMaskName_3
	case i == 32:
		return _OuterHeaderCreationMaskName_4
	default:
		return fmt.Sprintf("OuterHeaderCreationMask(%d)", i)
	}
}

var _OuterHeaderCreationMaskValues = []OuterHeaderCreationMask{1, 2, 4, 8, 16, 32}

var _OuterHeaderCreationMaskNameToValueMap = map[string]OuterHeaderCreationMask{
	_OuterHeaderCreationMaskName_0[0:35]:  1,
	_OuterHeaderCreationMaskName_0[35:70]: 2,
	_OuterHeaderCreationMaskName_1[0:30]:  4,
	_OuterHeaderCreationMaskName_2[0:30]:  8,
	_OuterHeaderCreationMaskName_3[0:26]:  16,
	_OuterHeaderCreationMaskName_4[0:26]:  32,
}

// OuterHeaderCreationMaskString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func OuterHeaderCreationMaskString(s string) (OuterHeaderCreationMask, error) {
	if val, ok := _OuterHeaderCreationMaskNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to OuterHeaderCreationMask values", s)
}

// OuterHeaderCreationMaskValues returns all values of the enum
func OuterHeaderCreationMaskValues() []OuterHeaderCreationMask {
	return _OuterHeaderCreationMaskValues
}

// IsAOuterHeaderCreationMask returns "true" if the value is listed in the enum definition. "false" otherwise
func (i OuterHeaderCreationMask) IsAOuterHeaderCreationMask() bool {
	for _, v := range _OuterHeaderCreationMaskValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalYAML implements a YAML Marshaler for OuterHeaderCreationMask
func (i OuterHeaderCreationMask) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for OuterHeaderCreationMask
func (i *OuterHeaderCreationMask) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = OuterHeaderCreationMaskString(s)
	return err
}
