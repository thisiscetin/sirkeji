package sirkeji

import (
	"sync"
	"testing"
)

// TestSubscribe ensures that subscribers can connect and receive events.
func TestSubscribe(t *testing.T) {
	streamer := NewStreamer()

	t.Run("Subscribe new subscriber", func(t *testing.T) {
		ch, err := streamer.Subscribe("user123")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ch == nil {
			t.Fatal("expected non-nil channel")
		}
	})

	t.Run("Subscribe duplicate subscriber", func(t *testing.T) {
		_, _ = streamer.Subscribe("user123") // First subscription
		_, err := streamer.Subscribe("user123")
		if err == nil {
			t.Fatal("expected error for duplicate subscription, got nil")
		}
		if err.Error() != "subscriber user123 already subscribed" {
			t.Fatalf("unexpected error message: %v", err)
		}
	})
}

// TestUnsubscribe ensures that subscribers can disconnect and their channels are closed.
func TestUnsubscribe(t *testing.T) {
	streamer := NewStreamer()

	t.Run("Unsubscribe existing subscriber", func(t *testing.T) {
		ch, _ := streamer.Subscribe("user123")
		streamer.Unsubscribe("user123")

		// Test that the channel is closed
		_, ok := <-ch
		if ok {
			t.Fatal("expected channel to be closed, but it is still open")
		}
	})

	t.Run("Unsubscribe non-existent subscriber", func(t *testing.T) {
		// Should not panic or throw an error
		streamer.Unsubscribe("nonexistent_user")
	})
}

// TestPublish ensures that events are delivered to all active subscribers.
func TestPublish(t *testing.T) {
	t.Run("Publish to multiple subscribers", func(t *testing.T) {
		streamer := NewStreamer()

		ch1, _ := streamer.Subscribe("user1")
		ch2, _ := streamer.Subscribe("user2")

		event := Event{Publisher: "system", Type: Info, Meta: "Test Event"}

		wg := sync.WaitGroup{}
		wg.Add(2)

		go func() {
			// Verify the event is received by both subscribers
			received1 := <-ch1
			if received1 != event {
				t.Errorf("expected event %+v, got %+v", event, received1)
				return
			}
			wg.Done()
		}()

		go func() {
			received2 := <-ch2
			if received2 != event {
				t.Errorf("expected event %+v, got %+v", event, received2)
				return
			}
			wg.Done()
		}()

		go streamer.Publish(event)

		wg.Wait()
	})

	t.Run("Publish with no subscribers", func(t *testing.T) {
		streamer := NewStreamer()

		// No subscribers; publish should not block or panic
		event := Event{Publisher: "system", Type: Info, Meta: "No Subscribers"}
		streamer.Publish(event)
	})
}

// TestConcurrentAccess ensures thread safety for Subscribe, Unsubscribe, and Publish.
func TestConcurrentAccess(t *testing.T) {
	streamer := NewStreamer()

	// Subscribe multiple subscribers concurrently
	go func() {
		_, _ = streamer.Subscribe("user1")
	}()
	go func() {
		_, _ = streamer.Subscribe("user2")
	}()
	go func() {
		streamer.Unsubscribe("user1")
	}()

	// Publish events concurrently
	go func() {
		event := Event{Publisher: "system", Type: Info, Meta: "Concurrent Event"}
		streamer.Publish(event)
	}()
}
