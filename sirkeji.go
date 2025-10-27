package sirkeji

import (
	"context"
	"os/signal"
	"syscall"
	"time"
)

// Subscribe is a helper function that simplifies the process of subscribing a Subscriber to a Streamer.
//
// This function creates a SubscriptionManager for the given Streamer and Subscriber,
// and immediately subscribes the Subscriber. If any errors occur during this process,
// the function panics, making it suitable for scenarios where failures during subscription
// are considered critical and should terminate the application.
//
// Parameters:
//   - streamer: The Streamer instance to which the Subscriber will be connected.
//   - subscriber: The Subscriber instance that will receive events from the Streamer.
//
// Panics:
//   - If the SubscriptionManager cannot be created (e.g., due to invalid arguments).
//   - If the subscription fails (e.g., due to duplicate Subscriber UIDs).
//
// Example Usage:
//
//	subscriber := &MySubscriber{}
//	streamer := sirkeji.NewStreamer()
//	sirkeji.Subscribe(streamer, subscriber)
func Subscribe(streamer Streamer, subscriber Subscriber) {
	manager, err := NewSubscriptionManager(streamer, subscriber)
	if err != nil {
		panic(err)
	}
	if err := manager.Subscribe(); err != nil {
		panic(err)
	}
}

// Unsubscribe is a helper function that simplifies the process of unsubscribe a Subscriber to a Streamer.
//
// This function creates a SubscriptionManager for the given Streamer and Subscriber,
// and immediately unsubscribes the Subscriber. If any errors occur during this process,
// the function panics, making it suitable for scenarios where failures during subscription
// are considered critical and should terminate the application.
//
// Parameters:
//   - streamer: The Streamer instance to which the Subscriber will be connected.
//   - subscriber: The Subscriber instance that will receive events from the Streamer.
//
// Panics:
//   - If the SubscriptionManager cannot be created (e.g., due to invalid arguments).
//
// Example Usage:
//
//	subscriber := &MySubscriber{}
//	streamer := sirkeji.NewStreamer()
//	sirkeji.UnSubscribe(streamer, subscriber)
func Unsubscribe(streamer Streamer, subscriber Subscriber) {
	manager, err := NewSubscriptionManager(streamer, subscriber)
	if err != nil {
		panic(err)
	}

	manager.Unsubscribe()
}

// WaitForTermination waits for OS termination signals and publishes a Shutdown event.
//
// Parameters:
//   - ctx: Global context.Context
//   - streamer: The Streamer instance to publish the Shutdown event.
//   - delay: Custom termination delay for application to close
//
// Behavior:
//   - Waits for SIGINT or SIGTERM signals.
//   - Publishes a Shutdown event with the publisher set to "main".
//   - Blocks execution until the termination signal is received and the event is published.
//
// Example:
//
//	sirkeji.WaitForTermination(streamer)
func WaitForTermination(ctx context.Context, streamer Streamer, delay time.Duration) {
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Wait for termination signal
	<-ctx.Done()

	// Publish a Shutdown event
	streamer.Publish(Event{
		Publisher: "main",
		Type:      Shutdown,
		Meta:      "Application is shutting down",
		Payload:   nil,
	})

	// Allow subscribers time to process the Shutdown event
	time.Sleep(delay)
}
