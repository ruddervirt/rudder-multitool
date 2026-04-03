//go:build linux
// +build linux

package changetime

import (
	"time"

	"golang.org/x/sys/unix"
)

func setTime(timestamp time.Time) error {
	utc := timestamp.UTC()
	return unix.Settimeofday(&unix.Timeval{
		Sec:  utc.Unix(),
		Usec: int64(utc.Nanosecond() / int(time.Microsecond)),
	})
}
