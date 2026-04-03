//go:build !windows && !linux && !darwin
// +build !windows,!linux,!darwin

package changetime

import (
	"fmt"
	"time"
)

func setTime(timestamp time.Time) error {
	return fmt.Errorf("change-time not supported on this platform")
}
