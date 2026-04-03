//go:build darwin
// +build darwin

package changetime

import (
	"time"

	"golang.org/x/sys/unix"
)

func setTime(timestamp time.Time) error {
	utc := timestamp.UTC()
	return unix.Settimeofday(&unix.Timeval{
		Sec:  utc.Unix(),
		Usec: int32(utc.Nanosecond() / int(time.Microsecond)),
	})
}
