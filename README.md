# ☱ Sirkeji

![Tests](https://github.com/thisiscetin/sirkeji/actions/workflows/tests.yml/badge.svg)
[![Go Report](https://goreportcard.com/badge/github.com/thisiscetin/sirkeji)](https://goreportcard.com/report/github.com/thisiscetin/sirkeji)
[![Go Reference](https://pkg.go.dev/badge/github.com/thisiscetin/sirkeji.svg)](https://pkg.go.dev/github.com/thisiscetin/sirkeji)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Sirkeji is a lightweight, in-memory event streaming library for Go designed to enable modular, event-centric architectures by allowing components to produce and consume events seamlessly.

Named after the historic Sirkeci Train Station, Sirkeji promotes a decoupled and predictable flow of interactions. It eliminates tightly coupled dependencies and simplifies extensibility without the performance overhead of external message brokers.

## Features

- **Event-Centric Architecture**: Focus on events (messages) as the primary means of communication to simplify your application design.
- **In-Memory Streaming**: Ultra-fast event processing without the complexity of external message brokers.
- **Decoupled Components**: Promote modularity by eliminating tightly coupled dependencies.
- **Extensible Subscribers**: Add, modify, or replace subscribers easily without breaking the system.

## Installation

Install Sirkeji using `go get`

```bash
go get github.com/thisiscetin/sirkeji
```

## Getting Started

To demonstrate Sirkeji, let's build a system where one component randomly publishes numbers, and others react to that event.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes. Make sure your code is well-tested and aligns with the project's goals.
- **In-Memory Streaming**: Ultra-fast event processing without the complexity of external message brokers.
- **Decoupled Components**: Promote modularity by eliminating tightly coupled dependencies.
- **Extensible Subscribers**: Add, modify, or replace subscribers easily without breaking the system.

## Installation

Install Sirkeji using `go get`

```bash
go get github.com/thisiscetin/sirkeji
```

## Getting Started

To demonstrate Sirkeji, let's build a system where one component randomly publishes numbers, and others react to following events.

Start by subscribing a built-in logger to visualize events, and a blocking function waiting for SIGTERM.
```go
package main

import "github.com/thisiscetin/sirkeji"

var (
	gStreamer = sirkeji.NewStreamer()
)

func main() {
	sirkeji.Subscribe(gStreamer, sirkeji.NewLogger())
	
	sirkeji.WaitForTermination(gStreamer)
}
```

Register an event and define your first component.

```go
func main() {
    // ... add to main
    sirkeji.Subscribe(gStreamer, NewNumberPublisher(gStreamer.Publish))
    // ...
}

var Number sirkeji.EventType = "Number"

func init() {
	sirkeji.RegisterEventType(Number)
}

type NumberPublisher struct {
	uid     string
	publish func(e sirkeji.Event)
}

func NewNumberPublisher(publish func(e sirkeji.Event)) NumberPublisher {
	return NumberPublisher{
		uid:     fmt.Sprintf("number-publisher-%d", time.Now().UnixMilli()),
		publish: publish,
	}
}

func (n NumberPublisher) Uid() string {
	return n.uid
}

func (n NumberPublisher) Process(event sirkeji.Event) {}

func (n NumberPublisher) OnSubscribed() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			randomNumber := rand.Intn(100)
			
			n.publish(sirkeji.Event{
				Publisher: n.uid,
				Type:      Number,
				Meta:      strconv.Itoa(randomNumber),
				Payload:   randomNumber,
			})
		}
	}()
}

func (n NumberPublisher) OnUnsubscribed() {}
```

When you run the application, you will immediately see a partially working software, and an output like below.

```bash
2024/11/26 00:12:38 [logger-1732569158496] subscribed to the streamer
2024/11/26 00:12:38 [number-publisher-1732569158496] subscribed to the streamer
2024/11/26 00:12:39.497028 [number-publisher-1732569158496] *Number*, m: 59 | pl: full
2024/11/26 00:12:40.497705 [number-publisher-1732569158496] *Number*, m: 85 | pl: full
2024/11/26 00:12:41.497917 [number-publisher-1732569158496] *Number*, m: 95 | pl: full
2024/11/26 00:12:42.498675 [number-publisher-1732569158496] *Number*, m: 3 | pl: full
```

Let's create another component which listens for a `Number` event, squares it and emits a `SquaredNumber` event.

```go
var SquaredNumber sirkeji.EventType = "SquaredNumber"

func init() {
	sirkeji.RegisterEventType(SquaredNumber)
}

type SquaredNumberPublisher struct {
	uid     string
	publish func(e sirkeji.Event)
}

func NewSquaredNumberPublisher(publish func(e sirkeji.Event)) SquaredNumberPublisher {
	return SquaredNumberPublisher{
		uid:     fmt.Sprintf("squared-number-publisher-%d", time.Now().UnixMilli()),
		publish: publish,
	}
}

func (s SquaredNumberPublisher) Uid() string {
	return s.uid
}

func (s SquaredNumberPublisher) Process(event sirkeji.Event) {
	if event.Type == Number {
		number := event.Payload.(int)
		numberSq := number * number

		s.publish(sirkeji.Event{
			Publisher: s.uid,
			Type:      SquaredNumber,
			Meta:      strconv.Itoa(numberSq),
			Payload:   numberSq,
		})
	}
}

func (s SquaredNumberPublisher) OnSubscribed() {}

func (s SquaredNumberPublisher) OnUnsubscribed() {}

```

Don't forget to attach components to the streamer.

```go
func main() {
	sirkeji.Subscribe(gStreamer, sirkeji.NewLogger())
	sirkeji.Subscribe(gStreamer, NewNumberPublisher(gStreamer.Publish))
	sirkeji.Subscribe(gStreamer, NewSquaredNumberPublisher(gStreamer.Publish))

	sirkeji.WaitForTermination(gStreamer)
}
```

When you run this code you will see an output like below.

```bash
2024/11/26 00:22:51 [logger-1732569771508] subscribed to the streamer
2024/11/26 00:22:51 [number-publisher-1732569771508] subscribed to the streamer
2024/11/26 00:22:51 [squared-number-publisher-1732569771508] subscribed to the streamer
2024/11/26 00:22:52.509421 [squared-number-publisher-1732569771508] *SquaredNumber*, m: 8464 | pl: full
2024/11/26 00:22:52.509539 [number-publisher-1732569771508] *Number*, m: 92 | pl: full
2024/11/26 00:22:53.510132 [squared-number-publisher-1732569771508] *SquaredNumber*, m: 2704 | pl: full
2024/11/26 00:22:53.510227 [number-publisher-1732569771508] *Number*, m: 52 | pl: full
2024/11/26 00:22:54.510833 [squared-number-publisher-1732569771508] *SquaredNumber*, m: 6889 | pl: full
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes. Make sure your code is well-tested and aligns with the project's goals.