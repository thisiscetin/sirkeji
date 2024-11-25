package sirkeji

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
