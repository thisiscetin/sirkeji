package sirkeji

import (
	"bytes"
	"fmt"
	"log"
	"time"
)

// Logger is a Subscriber implementation that logs all received events.
//
// It provides unique identification via `Uid()` and processes events by logging their details.
// A custom output can be provided during instantiation to direct log messages for testing or debugging purposes.
type Logger struct {
	uid    string
	logger *log.Logger
}

// NewLogger creates a new Logger instance with the specified output.
//
// Parameters:
//   - output: A *bytes.Buffer for capturing log output (useful for testing).
//
// Returns:
//   - A pointer to a new Logger instance.
//
// Example:
//
//	var buf bytes.Buffer
//	logger := NewLogger(&buf)
func NewLogger(output *bytes.Buffer) *Logger {
	return &Logger{
		uid:    fmt.Sprintf("logger-%d", time.Now().UnixNano()),
		logger: log.New(output, "", log.LstdFlags|log.Lmicroseconds),
	}
}

// Uid returns the unique identifier of the Logger.
//
// This identifier is used to manage Logger connections in the Streamer.
//
// Returns:
//   - A string representing the unique identifier of the Logger.
//
// Example:
//
//	fmt.Println(logger.Uid())
func (l *Logger) Uid() string {
	return l.uid
}

// Process logs the details of the received event.
//
// Parameters:
//   - event: The Event to be logged.
//
// Behavior:
//   - Logs the Publisher, Type, Meta, and Payload status of the event.
//   - Indicates whether the Payload is "empty" or "full".
//
// Example:
//
//	event := Event{Publisher: "system", Type: "Info", Meta: "App started", Payload: nil}
//	logger.Process(event)
func (l *Logger) Process(event Event) {
	payloadStatus := "empty"
	if event.Payload != nil {
		payloadStatus = "full"
	}
	l.logger.Printf("[%s] *%s*, m: %s | pl: %s", event.Publisher, event.Type, event.Meta, payloadStatus)
}

// OnSubscribed is called when the Logger is successfully subscribed to a Streamer.
//
// Behavior:
//   - Currently a no-op but can be extended to perform initialization or logging.
//
// Example:
//
//	logger.OnSubscribed()
func (l *Logger) OnSubscribed() {}

// OnUnsubscribed is called when the Logger is unsubscribed from a Streamer.
//
// Behavior:
//   - Currently a no-op but can be extended to perform cleanup or logging.
//
// Example:
//
//	logger.OnUnsubscribed()
func (l *Logger) OnUnsubscribed() {}
