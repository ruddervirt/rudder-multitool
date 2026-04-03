package startupdisabled

import "testing"

func TestIsStartupDisabled_Mock(t *testing.T) {
	name := "rudder_test_startup_entry"
	defer ClearMockStartupValue(name)

	// Write a disabled marker (0x03)
	SetMockStartupValue(name, []byte{0x03})

	disabled, err := IsStartupDisabled(name)
	if err != nil {
		t.Fatalf("IsStartupDisabled returned error: %v", err)
	}
	if !disabled {
		t.Fatalf("expected startup entry to be detected as disabled")
	}

	// Now mark as enabled (0x02)
	SetMockStartupValue(name, []byte{0x02})

	disabled, err = IsStartupDisabled(name)
	if err != nil {
		t.Fatalf("IsStartupDisabled returned error: %v", err)
	}
	if disabled {
		t.Fatalf("expected startup entry to be detected as enabled")
	}
}
