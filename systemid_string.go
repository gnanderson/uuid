// Code generated by "stringer -type=SystemId"; DO NOT EDIT.

package uuid

import "fmt"

const _SystemId_name = "SystemIdNoneSystemIdUserSystemIdEffectiveUserSystemIdGroupSystemIdEffectiveGroupSystemIdCallerProcessSystemIdCallerProcessParentSystemIdUnknown"

var _SystemId_index = [...]uint8{0, 12, 24, 45, 58, 80, 101, 128, 143}

func (i SystemId) String() string {
	if i >= SystemId(len(_SystemId_index)-1) {
		return fmt.Sprintf("SystemId(%d)", i)
	}
	return _SystemId_name[_SystemId_index[i]:_SystemId_index[i+1]]
}
