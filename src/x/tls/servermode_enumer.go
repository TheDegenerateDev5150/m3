// Code generated by "enumer -type ServerMode -text"; DO NOT EDIT.

package tls

import (
	"fmt"
	"strings"
)

const _ServerModeName = "DisabledPermissiveEnforced"

var _ServerModeIndex = [...]uint8{0, 8, 18, 26}

const _ServerModeLowerName = "disabledpermissiveenforced"

func (i ServerMode) String() string {
	if i >= ServerMode(len(_ServerModeIndex)-1) {
		return fmt.Sprintf("ServerMode(%d)", i)
	}
	return _ServerModeName[_ServerModeIndex[i]:_ServerModeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ServerModeNoOp() {
	var x [1]struct{}
	_ = x[Disabled-(0)]
	_ = x[Permissive-(1)]
	_ = x[Enforced-(2)]
}

var _ServerModeValues = []ServerMode{Disabled, Permissive, Enforced}

var _ServerModeNameToValueMap = map[string]ServerMode{
	_ServerModeName[0:8]:        Disabled,
	_ServerModeLowerName[0:8]:   Disabled,
	_ServerModeName[8:18]:       Permissive,
	_ServerModeLowerName[8:18]:  Permissive,
	_ServerModeName[18:26]:      Enforced,
	_ServerModeLowerName[18:26]: Enforced,
}

var _ServerModeNames = []string{
	_ServerModeName[0:8],
	_ServerModeName[8:18],
	_ServerModeName[18:26],
}

// ServerModeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ServerModeString(s string) (ServerMode, error) {
	if val, ok := _ServerModeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ServerModeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ServerMode values", s)
}

// ServerModeValues returns all values of the enum
func ServerModeValues() []ServerMode {
	return _ServerModeValues
}

// ServerModeStrings returns a slice of all String values of the enum
func ServerModeStrings() []string {
	strs := make([]string, len(_ServerModeNames))
	copy(strs, _ServerModeNames)
	return strs
}

// IsAServerMode returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ServerMode) IsAServerMode() bool {
	for _, v := range _ServerModeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for ServerMode
func (i ServerMode) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ServerMode
func (i *ServerMode) UnmarshalText(text []byte) error {
	var err error
	*i, err = ServerModeString(string(text))
	return err
}
