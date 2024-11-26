package sirkeji

import (
	"errors"
	"log"
)

// Subscriber defines the interface for receiving and processing events.
// Components that need to subscribe to a Streamer must implement this interface.
type Subscriber interface {
	// Uid returns the unique identifier of the subscriber.
	//
	// This identifier is used by the Streamer to manage subscriber connections.
	//
	// Returns:
	//   - A string representing the unique identifier of the subscriber.
	Uid() string

	// Process handles the received event.
	//
	// This method is invoked whenever an event is published to the Streamer
	// and is routed to this Subscriber.
	//
	// Parameters:
	//   - event: The Event instance to be processed by the subscriber.
	Process(event Event)

	// Subscribed is called when the subscriber is successfully connected to the Streamer.
	//
	// Use this method to perform any initialization or logging when the subscription is established.
	Subscribed()

	// Unsubscribed is called when the subscriber is disconnected from the Streamer.
	//
	// Use this method to perform any cleanup or logging when the subscription is terminated.
	Unsubscribed()
}

// SubscriptionManager manages the lifecycle of a Subscriber with a Streamer.
//
// It provides methods to connect and disconnect a Subscriber, handling the
// subscription process and ensuring events are delivered to the subscriber's
// Process method.
type SubscriptionManager struct {
	// streamer is the Streamer instance managing event subscriptions and publications.
	streamer Streamer

	// subscriber is the Subscriber instance receiving events from the Streamer.
	subscriber Subscriber
}

var (
	// ErrStreamerShouldNotBeNil is returned when a nil Streamer is provided to NewSubscriptionManager.
	ErrStreamerShouldNotBeNil = errors.New("streamer shouldn't be nil")

	// ErrSubscriberShouldNotBeNil is returned when a nil Subscriber is provided to NewSubscriptionManager.
	ErrSubscriberShouldNotBeNil = errors.New("subscriber shouldn't be nil")
)

// NewSubscriptionManager creates a new SubscriptionManager for a given Streamer and Subscriber.
//
// Parameters:
//   - streamer: The Streamer instance managing event delivery. Must not be nil.
//   - subscriber: The Subscriber instance to manage. Must not be nil.
//
// Returns:
//   - A pointer to a new SubscriptionManager instance.
//   - An error if either the streamer or subscriber is nil (ErrStreamerShouldNotBeNil or ErrSubscriberShouldNotBeNil).
//
// Errors:
//   - ErrStreamerShouldNotBeNil: Returned when the provided Streamer is nil.
//   - ErrSubscriberShouldNotBeNil: Returned when the provided Subscriber is nil.
//
// Example:
//
//	streamer := NewStreamer()
//	subscriber := &MySubscriber{}
//
//	manager, err := NewSubscriptionManager(streamer, subscriber)
//	if err != nil {
//	    log.Fatalf("Failed to create SubscriptionManager: %v", err)
//	}
func NewSubscriptionManager(streamer Streamer, subscriber Subscriber) (*SubscriptionManager, error) {
	if streamer == nil {
		return nil, ErrStreamerShouldNotBeNil
	}
	if subscriber == nil {
		return nil, ErrSubscriberShouldNotBeNil
	}

	return &SubscriptionManager{
		streamer:   streamer,
		subscriber: subscriber,
	}, nil
}

// Subscribe connects the subscriber to the Streamer and starts processing events.
//
// This method subscribes the Subscriber to the Streamer and spawns a goroutine
// to process incoming events by invoking the Subscriber's Process method.
//
// Parameters:
//   - None.
//
// Returns:
//   - An error if the subscription fails (e.g., duplicate subscriber UID).
//
// Behavior:
//   - Starts a goroutine to listen for events and route them to the Subscriber's Process method.
//   - Calls the Subscriber's Subscribed method upon successful subscription.
//
// Example:
//
//	manager := NewSubscriptionManager(streamer, subscriber)
//	err := manager.Subscribe()
//	if err != nil {
//	    log.Fatalf("failed to subscribe: %v", err)
//	}
func (sm *SubscriptionManager) Subscribe() error {
	ch, err := sm.streamer.Subscribe(sm.subscriber.Uid())
	if err != nil {
		return err
	}

	go func(ch chan Event) {
		for event := range ch {
			go sm.subscriber.Process(event)
		}
	}(ch)

	sm.subscriber.Subscribed()

	log.Printf("[%s] subscribed to the streamer\n", sm.subscriber.Uid())
	return nil
}

// Unsubscribe disconnects the subscriber from the Streamer.
//
// This method removes the Subscriber from the Streamer, ensuring it no longer
// receives events. It also invokes the Subscriber's Unsubscribed method.
//
// Parameters:
//   - None.
//
// Returns:
//   - None.
//
// Behavior:
//   - Calls the Subscriber's Unsubscribed method after successfully unsubscribing.
//
// Example:
//
//	manager := NewSubscriptionManager(streamer, subscriber)
//	manager.Unsubscribe()
func (sm *SubscriptionManager) Unsubscribe() {
	sm.streamer.Unsubscribe(sm.subscriber.Uid())
	sm.subscriber.Unsubscribed()

	log.Printf("[%s] unsubscribed from the streamer\n", sm.subscriber.Uid())
}
