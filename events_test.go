package sirkeji

import (
	"testing"
)

// TestNewEvent ensures NewEvent correctly constructs an Event and handles invalid inputs.
func TestNewEvent(t *testing.T) {
	t.Run("Valid Event", func(t *testing.T) {
		event := NewEvent("test-publisher", Info, "metadata", nil)

		if event.Publisher != "test-publisher" {
			t.Errorf("expected Publisher 'test-publisher', got '%s'", event.Publisher)
		}
		if event.Type != Info {
			t.Errorf("expected Type 'Info', got '%s'", event.Type)
		}
		if event.Meta != "metadata" {
			t.Errorf("expected Meta 'metadata', got '%s'", event.Meta)
		}
		if event.Payload != nil {
			t.Errorf("expected Payload 'nil', got '%v'", event.Payload)
		}
	})

	t.Run("Empty Publisher Panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic for empty Publisher")
			}
		}()
		NewEvent("", Info, "metadata", nil)
	})

	t.Run("Empty Type Panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic for empty Type")
			}
		}()
		NewEvent("test-publisher", "", "metadata", nil)
	})
}

// TestInfoEvent ensures InfoEvent helper constructs an Info event correctly.
func TestInfoEvent(t *testing.T) {
	event := InfoEvent("system", "Application started")

	if event.Publisher != "system" {
		t.Errorf("expected Publisher 'system', got '%s'", event.Publisher)
	}
	if event.Type != Info {
		t.Errorf("expected Type 'Info', got '%s'", event.Type)
	}
	if event.Meta != "Application started" {
		t.Errorf("expected Meta 'Application started', got '%s'", event.Meta)
	}
	if event.Payload != nil {
		t.Errorf("expected Payload 'nil', got '%v'", event.Payload)
	}
}

// TestErrorEvent ensures ErrorEvent helper constructs an Error event correctly.
func TestErrorEvent(t *testing.T) {
	event := ErrorEvent("database", "Connection failed")

	if event.Publisher != "database" {
		t.Errorf("expected Publisher 'database', got '%s'", event.Publisher)
	}
	if event.Type != Error {
		t.Errorf("expected Type 'Error', got '%s'", event.Type)
	}
	if event.Meta != "Connection failed" {
		t.Errorf("expected Meta 'Connection failed', got '%s'", event.Meta)
	}
	if event.Payload != nil {
		t.Errorf("expected Payload 'nil', got '%v'", event.Payload)
	}
}

// TestRegisterEventType ensures the event type registry behaves correctly.
func TestRegisterEventType(t *testing.T) {
	customEventType := EventType("CustomEvent")

	t.Run("Register New EventType", func(t *testing.T) {
		RegisterEventType(customEventType)

		if !IsEventTypeRegistered(customEventType) {
			t.Errorf("expected EventType 'CustomEvent' to be registered")
		}
	})

	t.Run("Duplicate Registration Panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic for duplicate EventType registration")
			}
		}()
		RegisterEventType(customEventType)
	})
}

// TestIsEventTypeRegistered ensures IsEventTypeRegistered behaves as expected.
func TestIsEventTypeRegistered(t *testing.T) {
	t.Run("Predefined EventTypes", func(t *testing.T) {
		if !IsEventTypeRegistered(Error) {
			t.Errorf("expected EventType 'Error' to be registered")
		}
		if !IsEventTypeRegistered(Info) {
			t.Errorf("expected EventType 'Info' to be registered")
		}
		if !IsEventTypeRegistered(Shutdown) {
			t.Errorf("expected EventType 'Shutdown' to be registered")
		}
	})

	t.Run("Unregistered EventType", func(t *testing.T) {
		if IsEventTypeRegistered("UnregisteredEvent") {
			t.Errorf("did not expect 'UnregisteredEvent' to be registered")
		}
	})
}
