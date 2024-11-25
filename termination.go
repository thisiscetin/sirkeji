package sirkeji

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

// HandleTerminationBlocking waits for OS termination signals and publishes a Shutdown event.
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
//	sirkeji.HandleTerminationBlocking(streamer)
func HandleTerminationBlocking(streamer Streamer) {
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
