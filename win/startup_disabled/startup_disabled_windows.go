//go:build windows
// +build windows

package startupdisabled

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

// IsStartupDisabled checks the StartupApproved Run registry keys for the
// given startup entry name. Returns true if the entry exists and is disabled.
func IsStartupDisabled(name string) (bool, error) {
	// Check current user first, then local machine
	paths := []struct {
		root registry.Key
		path string
	}{
		{registry.CURRENT_USER, `Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\StartupApproved\\Run`},
		{registry.LOCAL_MACHINE, `Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\StartupApproved\\Run`},
	}

	for _, p := range paths {
		k, err := registry.OpenKey(p.root, p.path, registry.READ)
		if err != nil {
			continue
		}
		defer k.Close()

		val, _, err := k.GetBinaryValue(name)
		if err != nil {
			continue // entry not found in this hive
		}

		if len(val) < 1 {
			return false, fmt.Errorf("unexpected empty value for %q", name)
		}

		// First byte: 02/06 = enabled, 03/07 = disabled
		disabled := val[0] == 0x03 || val[0] == 0x07
		return disabled, nil
	}

	return false, fmt.Errorf("NOT_FOUND")
}
