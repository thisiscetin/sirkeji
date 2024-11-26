package sirkeji

import (
	"errors"
	"sync"
	"testing"
	"time"
)

// MockSubscriber is a mock implementation of the Subscriber interface for testing.
type MockSubscriber struct {
	uid          string
	processed    []Event
	subscribed   bool
	unsubscribed bool
	sync.Mutex
}

func NewMockSubscriber(uid string) *MockSubscriber {
	return &MockSubscriber{
		uid:       uid,
		processed: []Event{},
	}
}

func (ms *MockSubscriber) Uid() string {
	return ms.uid
}

func (ms *MockSubscriber) Process(event Event) {
	ms.Lock()
	defer ms.Unlock()

	ms.processed = append(ms.processed, event)
}

func (ms *MockSubscriber) Subscribed() {
	ms.Lock()
	defer ms.Unlock()

	ms.subscribed = true
}

func (ms *MockSubscriber) Unsubscribed() {
	ms.Lock()
	defer ms.Unlock()

	ms.unsubscribed = true
}

func (ms *MockSubscriber) GetProcessedEvents() []Event {
	ms.Lock()
	defer ms.Unlock()

	return ms.processed
}

// TestNewSubscriptionManager verifies the behavior of NewSubscriptionManager.
func TestNewSubscriptionManager(t *testing.T) {
	streamer := NewStreamer()
	subscriber := NewMockSubscriber("test-subscriber")

	t.Run("Valid Initialization", func(t *testing.T) {
		manager, err := NewSubscriptionManager(streamer, subscriber)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if manager == nil {
			t.Fatal("expected non-nil SubscriptionManager")
		}
	})

	t.Run("Nil Streamer", func(t *testing.T) {
		manager, err := NewSubscriptionManager(nil, subscriber)
		if !errors.Is(err, ErrStreamerShouldNotBeNil) {
			t.Fatalf("expected error: %v, got: %v", ErrStreamerShouldNotBeNil, err)
		}
		if manager != nil {
			t.Fatal("expected nil SubscriptionManager")
		}
	})

	t.Run("Nil Subscriber", func(t *testing.T) {
		manager, err := NewSubscriptionManager(streamer, nil)
		if !errors.Is(err, ErrSubscriberShouldNotBeNil) {
			t.Fatalf("expected error: %v, got: %v", ErrSubscriberShouldNotBeNil, err)
		}
		if manager != nil {
			t.Fatal("expected nil SubscriptionManager")
		}
	})
}

// TestSubscriptionManagerSubscribe verifies the behavior of the Subscribe method.
func TestSubscriptionManagerSubscribe(t *testing.T) {
	streamer := NewStreamer()
	subscriber := NewMockSubscriber("test-subscriber")
	manager, err := NewSubscriptionManager(streamer, subscriber)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = manager.Subscribe()
	if err != nil {
		t.Fatalf("unexpected error during subscription: %v", err)
	}

	// Verify that the subscriber's Subscribed method was called
	if !subscriber.subscribed {
		t.Errorf("expected Subscribed to be called, but it wasn't")
	}

	// Verify that the subscriber receives published events
	event := Event{Publisher: "system", Type: Info, Meta: "Test Event"}
	streamer.Publish(event)

	time.Sleep(100 * time.Millisecond) // Allow time for the event to propagate

	processedEvents := subscriber.GetProcessedEvents()
	if len(processedEvents) != 1 {
		t.Fatalf("expected 1 processed event, got %d", len(processedEvents))
	}
	if processedEvents[0] != event {
		t.Errorf("expected event %+v, got %+v", event, processedEvents[0])
	}
}

// TestSubscriptionManagerUnsubscribe verifies the behavior of the Unsubscribe method.
func TestSubscriptionManagerUnsubscribe(t *testing.T) {
	streamer := NewStreamer()
	subscriber := NewMockSubscriber("test-subscriber")
	manager, err := NewSubscriptionManager(streamer, subscriber)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_ = manager.Subscribe()
	manager.Unsubscribe()

	// Verify that the subscriber's Unsubscribed method was called
	if !subscriber.unsubscribed {
		t.Errorf("expected Unsubscribed to be called, but it wasn't")
	}

	// Verify that the subscriber no longer receives events
	event := Event{Publisher: "system", Type: Info, Meta: "Test Event"}
	streamer.Publish(event)

	time.Sleep(100 * time.Millisecond) // Allow time for events to propagate

	processedEvents := subscriber.GetProcessedEvents()
	if len(processedEvents) != 0 {
		t.Fatalf("expected 0 processed events, got %d", len(processedEvents))
	}
}
