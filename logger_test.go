package sirkeji

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// TestProcess verifies the behavior of the Logger's Process method.
//
// It checks that the Logger correctly logs the details of an event, including
// the Publisher, Type, Meta, and whether the Payload is "empty" or "full".
//
// Scenarios Tested:
//   - Event with an empty Payload.
//   - Event with a non-empty Payload.
func TestProcess(t *testing.T) {
	// Create a buffer to capture log output.
	var buf bytes.Buffer
	logger := NewLogger(&buf)

	// Subtest: Verify logging for an event with an empty Payload.
	t.Run("Process Event with Empty Payload", func(t *testing.T) {
		// Create an event with no Payload.
		event := Event{
			Publisher: "system",
			Type:      "Info",
			Meta:      "Application started",
			Payload:   nil,
		}

		// Process the event.
		logger.Process(event)

		// Verify the log output.
		logOutput := buf.String()
		expected := fmt.Sprintf("[%s] *%s*, m: %s | pl: %s", event.Publisher, event.Type, event.Meta, "empty")
		if !strings.Contains(logOutput, expected) {
			t.Errorf("expected log to contain '%s', got '%s'", expected, logOutput)
		}

		// Clear the buffer for the next subtest.
		buf.Reset()
	})

	// Subtest: Verify logging for an event with a non-empty Payload.
	t.Run("Process Event with Non-Empty Payload", func(t *testing.T) {
		// Create an event with a Payload.
		event := Event{
			Publisher: "user-service",
			Type:      "Error",
			Meta:      "User creation failed",
			Payload:   map[string]string{"username": "johndoe"},
		}

		// Process the event.
		logger.Process(event)

		// Verify the log output.
		logOutput := buf.String()
		expected := fmt.Sprintf("[%s] *%s*, m: %s | pl: %s", event.Publisher, event.Type, event.Meta, "full")
		if !strings.Contains(logOutput, expected) {
			t.Errorf("expected log to contain '%s', got '%s'", expected, logOutput)
		}

		// Clear the buffer for safety.
		buf.Reset()
	})
}
