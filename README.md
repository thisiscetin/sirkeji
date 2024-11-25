## Sirkeji


**Sirkeji** is a lightweight, in-memory event streamer designed for building **modular**, **maintainable**, and **event-centric applications** in Go. Inspired by the principles of event-driven architectures, Sirkeji focuses on simplicity, performance, and observability, making it ideal for applications where using an external event streamer like Kafka is not feasible.

## 🌟 Key Features

- **In-Memory Event Streaming**  
  Eliminate the overhead of external tools like Kafka or RabbitMQ for applications with low latency requirements.

- **Event-Centric Design**  
  Build applications as collections of small, single-purpose components that produce and consume events without tight coupling.

- **Modular and Maintainable**  
  Sirkeji encourages modularity, enabling you to expand your system organically as new requirements emerge.

---

## 🚀 Why Sirkeji?

An **event-centered application** is a living system, built on small, purpose-driven components—crafted, packaged, and set in motion within a message-driven lifecycle. At its core is a simple principle:

> Components announce what they produce and what they need, listening for answers without concern for where those answers will come from.

Think of it like a city:  
Just as great cities flourish around a central river, event-centric systems thrive around a **data stream**. New features emerge like neighborhoods and bridges, adapting to new demands while remaining cohesive.

Sirkeji embraces this philosophy to simplify the development of maintainable and observable monoliths without compromising on performance.

---

## 🛠️ Getting Started

### Installation

```bash
go get github.com/thisiscetin/sirkeji