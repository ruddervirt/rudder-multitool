package changetime

import (
	"testing"
	"time"
)

func TestParseTimestamp_RFC3339(t *testing.T) {
	parsed, err := parseTimestampWithZone("2026-02-20T12:00:00Z", ZoneAuto)
	if err != nil {
		t.Fatalf("expected parse to succeed: %v", err)
	}
	if parsed.UTC().Format(time.RFC3339) != "2026-02-20T12:00:00Z" {
		t.Fatalf("unexpected parsed time: %s", parsed.UTC().Format(time.RFC3339))
	}
}

func TestParseTimestamp_LocalFallback(t *testing.T) {
	parsed, err := parseTimestampWithZone("2026-02-20T12:00:00", ZoneLocal)
	if err != nil {
		t.Fatalf("expected parse to succeed: %v", err)
	}
	if parsed.Location() != time.Local {
		t.Fatalf("expected local location for fallback format")
	}
}

func TestParseTimestamp_AutoUsesLocal(t *testing.T) {
	parsed, err := parseTimestampWithZone("2026-02-20T12:00:00", ZoneAuto)
	if err != nil {
		t.Fatalf("expected parse to succeed: %v", err)
	}
	if parsed.Location() != time.Local {
		t.Fatalf("expected auto mode to use local location")
	}
}

func TestParseTimestamp_UTCFlag(t *testing.T) {
	parsed, err := parseTimestampWithZone("2026-02-20T12:00:00", ZoneUTC)
	if err != nil {
		t.Fatalf("expected parse to succeed: %v", err)
	}
	if parsed.Location() != time.UTC {
		t.Fatalf("expected utc location for utc mode")
	}
	if parsed.Format(time.RFC3339) != "2026-02-20T12:00:00Z" {
		t.Fatalf("unexpected parsed time: %s", parsed.Format(time.RFC3339))
	}
}

func TestParseTimestamp_Invalid(t *testing.T) {
	if _, err := parseTimestampWithZone("not-a-time", ZoneAuto); err == nil {
		t.Fatalf("expected parse to fail for invalid input")
	}
}
