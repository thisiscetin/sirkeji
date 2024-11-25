package sirkeji

import (
	"fmt"
	"sync"
)

// Streamer defines the interface for an event stream.
// It allows subscribers to connect and disconnect, and it broadcasts events to all active subscribers.
type Streamer interface {
	// Subscribe connects a subscriber to the Streamer and returns a channel for receiving events.
	// Parameters:
	//   - subscriberUid: A unique identifier for the subscriber.
	//
	// Returns:
	//   - A channel for receiving events.
	//   - An error if the subscriberUid is already subscribed.
	Subscribe(subscriberUid string) (chan Event, error)

	// Unsubscribe removes a subscriber from the Streamer and closes its event channel.
	// Parameters:
	//   - subscriberUid: The unique identifier of the subscriber to remove.
	Unsubscribe(subscriberUid string)

	// Publish broadcasts an event to all connected subscribers.
	// Parameters:
	//   - event: The Event to be published.
	Publish(event Event)
}

// DefaultStreamer is the default implementation of the Streamer interface.
// It manages subscribers and broadcasts events to all active channels.
type DefaultStreamer struct {
	// subscribers holds a map of subscriber IDs to their event channels.
	subscribers map[string]chan Event
	// RWMutex ensures thread-safe access to the subscribers map.
	sync.RWMutex
}

// NewStreamer creates and returns a new instance of DefaultStreamer.
//
// Returns:
//   - A pointer to a new DefaultStreamer.
func NewStreamer() *DefaultStreamer {
	return &DefaultStreamer{
		subscribers: make(map[string]chan Event),
	}
}

// Subscribe connects a subscriber to the DefaultStreamer and returns its event channel.
//
// Parameters:
//   - subscriberUid: A unique identifier for the subscriber.
//
// Returns:
//   - A channel for receiving events.
//   - An error if the subscriberUid is already subscribed.
//
// Behavior:
//   - If the subscriberUid is already in use, an error is returned.
//   - A new channel is created for the subscriber and added to the subscribers map.
//
// Example:
//
//	streamer := NewStreamer()
//	ch, err := streamer.Subscribe("user123")
//	if err != nil {
//	    log.Fatalf("failed to subscribe: %v", err)
//	}
func (s *DefaultStreamer) Subscribe(subscriberUid string) (chan Event, error) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.subscribers[subscriberUid]; ok {
		return nil, fmt.Errorf("subscriber %s already subscribed", subscriberUid)
	}

	ch := make(chan Event)
	s.subscribers[subscriberUid] = ch
	return ch, nil
}

// Unsubscribe removes a subscriber from the DefaultStreamer and closes its event channel.
//
// Parameters:
//   - subscriberUid: The unique identifier of the subscriber to remove.
//
// Behavior:
//   - Closes the subscriber's event channel.
//   - Removes the subscriberUid from the subscribers map.
//   - If the subscriberUid is not found, no action is taken.
//
// Example:
//
//	streamer := NewStreamer()
//	streamer.Unsubscribe("user123")
func (s *DefaultStreamer) Unsubscribe(subscriberUid string) {
	s.Lock()
	defer s.Unlock()

	if ch, ok := s.subscribers[subscriberUid]; ok {
		close(ch)
	}
	delete(s.subscribers, subscriberUid)
}

// Publish broadcasts an event to all connected subscribers.
//
// Parameters:
//   - event: The Event to be published.
//
// Behavior:
//   - Sends the event to all active subscriber channels.
//   - If a channel is blocked or slow, the operation may pause.
//
// Example:
//
//	streamer := NewStreamer()
//	event := Event{Publisher: "system", Type: Info, Meta: "App started"}
//	streamer.Publish(event)
func (s *DefaultStreamer) Publish(event Event) {
	s.RLock()
	defer s.RUnlock()

	for _, subscriber := range s.subscribers {
		subscriber <- event
	}
}
