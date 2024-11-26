package sirkeji

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Logger is a Subscriber implementation that logs all received events.
//
// Logs are written to the screen by default (os.Stdout). Additional outputs
// such as a bytes.Buffer or a file can be added dynamically.
type Logger struct {
	uid              string
	baseOutput       io.Writer
	additionalOutput io.Writer
	logger           *log.Logger
}

// NewLogger creates a new Logger instance.
//
// Logs are written to the screen (os.Stdout) by default.
//
// Returns:
//   - A pointer to a new Logger instance.
//
// Example:
//
//	logger := NewLogger()
func NewLogger() *Logger {
	return &Logger{
		uid:        fmt.Sprintf("logger-%d", time.Now().UnixMilli()),
		baseOutput: os.Stdout,
		logger:     log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds),
	}
}

// SetAdditionalOutput sets an optional second output for the logger.
//
// Parameters:
//   - output: An io.Writer (e.g., bytes.Buffer, file) to receive log messages in addition to the screen.
//
// Behavior:
//   - Updates the logger to write to both os.Stdout and the provided output.
//   - If no additional output is set, the logger writes only to os.Stdout.
//
// Example:
//
//	var buf bytes.Buffer
//	logger.SetAdditionalOutput(&buf)
//
//	file, _ := os.Create("log.txt")
//	defer file.Close()
//	logger.SetAdditionalOutput(file)
func (l *Logger) SetAdditionalOutput(output io.Writer) {
	if output != nil {
		l.additionalOutput = output
		multiWriter := io.MultiWriter(l.baseOutput, l.additionalOutput)
		l.logger.SetOutput(multiWriter)
	} else {
		l.logger.SetOutput(l.baseOutput)
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
//	logger.Subscribed()
func (l *Logger) Subscribed() {}

// OnUnsubscribed is called when the Logger is unsubscribed from a Streamer.
//
// Behavior:
//   - Currently a no-op but can be extended to perform cleanup or logging.
//
// Example:
//
//	logger.Unsubscribed()
func (l *Logger) Unsubscribed() {}
