//go:build windows
// +build windows

package changetime

import (
	"fmt"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	kernel32          = windows.NewLazySystemDLL("kernel32.dll")
	procSetSystemTime = kernel32.NewProc("SetSystemTime")
)

func setTime(timestamp time.Time) error {
	utc := timestamp.UTC()
	systemTime := windows.Systemtime{
		Year:         uint16(utc.Year()),
		Month:        uint16(utc.Month()),
		Day:          uint16(utc.Day()),
		Hour:         uint16(utc.Hour()),
		Minute:       uint16(utc.Minute()),
		Second:       uint16(utc.Second()),
		Milliseconds: uint16(utc.Nanosecond() / int(time.Millisecond)),
	}
	r1, _, err := procSetSystemTime.Call(uintptr(unsafe.Pointer(&systemTime)))
	if r1 == 0 {
		return fmt.Errorf("SetSystemTime failed: %w", err)
	}
	return nil
}
