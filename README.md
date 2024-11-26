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

*Please head to the examples folder for a working example.*

Sirkeji needs components to implement `sirkeji.Subscriber` interface to connect them to the streamer.

- a `Uid() string` function which returns a `string` unique id of the component
- a `Process(event sirkeji.Event)` function, which is called in a dedicated goroutine for processing events
- a `Subscribed()` function to perform boot-up operations like initializing a ticker in a separate goroutine
- a `Unsubscribed()` function to perform clean-up operations and handling graceful shutdowns

```go

type Publisher struct {}

func (p *Publisher) Uid() string {}

func (p *Publisher) Process(event sirkeji.Event) {}

func (p *Publisher) Subscribed() {}

func (p *Publisher) Unsubscribed() {}
```

One way to subscribe components to a stream is as follows:

```go
var (
	gStreamer = sirkeji.NewStreamer()
)

func main() {
	sirkeji.Subscribe(gStreamer, sirkeji.NewLogger())
	sirkeji.Subscribe(gStreamer, number.NewPublisher("number-publisher-1", gStreamer.Publish))
	sirkeji.Subscribe(gStreamer, number.NewPublisher("number-publisher-2", gStreamer.Publish))
	sirkeji.Subscribe(gStreamer, squared_number.NewPublisher("squared-number-publisher-1", gStreamer.Publish))
	sirkeji.Subscribe(gStreamer, number_count.NewPublisher("number-count-publisher-1", gStreamer.Publish))

	sirkeji.WaitForTermination(gStreamer)
}
```

*Note: With Sirkeji, you can also subscribe and unsubscribe components dynamically and perform much more complex operations. Please refer to the godoc for details.*

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes. Make sure your code is well-tested and aligns with the project's goals.