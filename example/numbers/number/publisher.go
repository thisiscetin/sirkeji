package number

import (
	"github.com/thisiscetin/sirkeji"
	"github.com/thisiscetin/sirkeji/example/numbers/events"
	"math/rand/v2"
	"strconv"
	"time"
)

type Publisher struct {
	uid     string
	publish func(e sirkeji.Event)
}

func (p *Publisher) Uid() string {
	return p.uid
}

func (p *Publisher) Process(event sirkeji.Event) {}

func (p *Publisher) Subscribed() {
	go func() {
		for range time.Tick(time.Second * 2) {
			n := rand.IntN(1_000)

			p.publish(sirkeji.Event{
				Publisher: p.uid,
				Type:      events.Number,
				Meta:      strconv.Itoa(n),
				Payload:   n,
			})
		}
	}()
}

func (p *Publisher) Unsubscribed() {}

func NewPublisher(uid string, publish func(e sirkeji.Event)) *Publisher {
	return &Publisher{
		uid,
		publish,
	}
}
