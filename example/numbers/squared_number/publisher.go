package squared_number

import (
	"github.com/thisiscetin/sirkeji"
	"github.com/thisiscetin/sirkeji/example/numbers/events"
	"strconv"
)

type Publisher struct {
	uid     string
	publish func(e sirkeji.Event)
}

func (p *Publisher) Uid() string {
	return p.uid
}

func (p *Publisher) Process(event sirkeji.Event) {
	if event.Type == events.Number {
		n := event.Payload.(int)
		nSquare := n * n

		p.publish(sirkeji.Event{
			Publisher: p.uid,
			Type:      events.SquaredNumber,
			Meta:      strconv.Itoa(nSquare),
			Payload:   nSquare,
		})
	}
}

func (p *Publisher) Subscribed() {}

func (p *Publisher) Unsubscribed() {}

func NewPublisher(uid string, publish func(e sirkeji.Event)) *Publisher {
	return &Publisher{
		uid:     uid,
		publish: publish,
	}
}
