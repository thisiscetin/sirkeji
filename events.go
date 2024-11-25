package sirkeji

import "sync"

// EventType represents the type of event.
// Used to categorize and handle different kinds of events within the system.
type EventType string

// Event represents an event in the system.
//
// Fields:
//   - Publisher: The originator of the event (e.g., system or component name).
//   - Type: The type of the event, defined by EventType.
//   - Meta: Optional metadata describing the event.
//   - Payload: Optional additional data associated with the event.
type Event struct {
	Publisher string
	Type      EventType
	Meta      string
	Payload   interface{}
}

// NewEvent creates a new Event with the required fields.
func NewEvent(publisher string, eventType EventType, meta string, payload interface{}) Event {
	if publisher == "" {
		panic("event must have a non-empty publisher")
	}
	if eventType == "" {
		panic("event must have a non-empty type")
	}
	return Event{
		Publisher: publisher,
		Type:      eventType,
		Meta:      meta,
		Payload:   payload,
	}
}

// Predefined EventTypes represent commonly used event categories.
const (
	// Error represents an error event, typically used to signal an issue.
	Error EventType = "Error"

	// Info represents an informational event, used for logging or status updates.
	Info EventType = "Info"

	// Shutdown represents a shutdown event, used during application termination.
	Shutdown EventType = "Shutdown"
)

// InfoEvent creates an informational event.
//
// Simplifies the creation of an Event with the predefined `Info` EventType.
//
// Parameters:
//   - publisher: The origin of the event (e.g., system or component name).
//   - message: A descriptive message about the event.
//
// Returns:
//   - An Event instance with the `Info` EventType.
//
// Example:
//
//	event := InfoEvent("system", "Application started")
//	gStreamer.Publish(event)
func InfoEvent(publisher, message string) Event {
	return Event{
		Publisher: publisher,
		Type:      Info,
		Meta:      message,
	}
}

// ErrorEvent creates an error event.
//
// Simplifies the creation of an Event with the predefined `Error` EventType.
//
// Parameters:
//   - publisher: The origin of the event (e.g., system or component name).
//   - message: A descriptive message about the error.
//
// Returns:
//   - An Event instance with the `Error` EventType.
//
// Example:
//
//	event := ErrorEvent("database", "Connection failed")
//	gStreamer.Publish(event)
func ErrorEvent(publisher, message string) Event {
	return Event{
		Publisher: publisher,
		Type:      Error,
		Meta:      message,
	}
}

// eventTypeRegistry is a thread-safe registry for EventTypes.
// Ensures that each EventType is unique within the system.
var (
	eventTypeRegistry = struct {
		sync.RWMutex
		types map[EventType]struct{}
	}{types: map[EventType]struct{}{
		Error:    struct{}{},
		Info:     struct{}{},
		Shutdown: struct{}{},
	}}
)

// RegisterEventType registers a new EventType to ensure uniqueness.
//
// Panics if the EventType is already registered, preventing duplication.
//
// Parameters:
//   - eventType: The EventType to register.
//
// Example:
//
//	RegisterEventType("CustomEvent")
func RegisterEventType(eventType EventType) {
	eventTypeRegistry.Lock()
	defer eventTypeRegistry.Unlock()

	if _, exists := eventTypeRegistry.types[eventType]; exists {
		panic("duplicate event type registration: " + string(eventType))
	}
	eventTypeRegistry.types[eventType] = struct{}{}
}

// IsEventTypeRegistered checks if an EventType is already registered.
//
// Parameters:
//   - eventType: The EventType to check.
//
// Returns:
//   - true if the EventType is found in the registry, false otherwise.
//
// Example:
//
//	if !IsEventTypeRegistered("CustomEvent") {
//	    RegisterEventType("CustomEvent")
//	}
func IsEventTypeRegistered(eventType EventType) bool {
	eventTypeRegistry.RLock()
	defer eventTypeRegistry.RUnlock()

	_, exists := eventTypeRegistry.types[eventType]
	return exists
}
