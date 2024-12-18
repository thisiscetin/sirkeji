package sirkeji

import (
	"os"
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

// WaitForTermination waits for OS termination signals and publishes a Shutdown event.
//
// Parameters:
//   - streamer: The Streamer instance to publish the Shutdown event.
//
// Behavior:
//   - Waits for SIGINT or SIGTERM signals.
//   - Publishes a Shutdown event with the publisher set to "main".
//   - Blocks execution until the termination signal is received and the event is published.
//
// Example:
//
//	sirkeji.WaitForTermination(streamer)
func WaitForTermination(streamer Streamer) {
	chExit := make(chan os.Signal, 1)
	signal.Notify(chExit, os.Interrupt, syscall.SIGTERM)

	// Wait for termination signal
	<-chExit

	// Publish a Shutdown event
	streamer.Publish(Event{
		Publisher: "main",
		Type:      "Shutdown",
		Meta:      "Application is shutting down",
		Payload:   nil,
	})

	// Allow subscribers time to process the Shutdown event
	time.Sleep(2 * time.Second)
}
