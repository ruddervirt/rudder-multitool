//go:build !windows
// +build !windows

package startupdisabled

import "fmt"

// Non-windows stub used for tests on Linux. Tests can set mock values via
// SetMockStartupValue and ClearMockStartupValue.
var mockStartup = map[string][]byte{}

func IsStartupDisabled(name string) (bool, error) {
	val, ok := mockStartup[name]
	if !ok {
		return false, fmt.Errorf("startup entry %q not found", name)
	}
	if len(val) < 1 {
		return false, fmt.Errorf("unexpected empty value for %q", name)
	}
	disabled := val[0] == 0x03 || val[0] == 0x07
	return disabled, nil
}

// SetMockStartupValue sets a mock binary value for a startup entry (tests only).
func SetMockStartupValue(name string, val []byte) {
	mockStartup[name] = val
}

// ClearMockStartupValue removes a mock startup entry.
func ClearMockStartupValue(name string) {
	delete(mockStartup, name)
}
