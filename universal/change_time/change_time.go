package changetime

import (
	"fmt"
	"time"
)

// ZoneMode controls how timestamps without a timezone are interpreted.
type ZoneMode int

const (
	ZoneAuto ZoneMode = iota
	ZoneLocal
	ZoneUTC
)

// SetTime parses an ISO8601 timestamp and sets the system time.
func SetTime(timestamp string) error {
	return SetTimeWithZone(timestamp, ZoneAuto)
}

// SetTimeWithZone parses an ISO8601 timestamp and sets the system time.
// ZoneMode only applies to timestamps without an explicit offset.
func SetTimeWithZone(timestamp string, mode ZoneMode) error {
	parsed, err := parseTimestampWithZone(timestamp, mode)
	if err != nil {
		return err
	}
	return setTime(parsed)
}

func parseTimestampWithZone(timestamp string, mode ZoneMode) (time.Time, error) {
	if timestamp == "" {
		return time.Time{}, fmt.Errorf("timestamp is required")
	}

	if parsed, err := time.Parse(time.RFC3339Nano, timestamp); err == nil {
		return parsed, nil
	}
	if parsed, err := time.Parse(time.RFC3339, timestamp); err == nil {
		return parsed, nil
	}

	loc := time.Local
	if mode == ZoneUTC {
		loc = time.UTC
	}

	parsed, err := time.ParseInLocation("2006-01-02T15:04:05", timestamp, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timestamp: %s", timestamp)
	}
	return parsed, nil
}
