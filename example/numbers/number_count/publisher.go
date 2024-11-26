package number_count

import (
	"github.com/thisiscetin/sirkeji"
	"github.com/thisiscetin/sirkeji/example/numbers/events"
	"strconv"
	"sync"
	"time"
)

type Publisher struct {
	uid     string
	publish func(e sirkeji.Event)

	count int
	sync.RWMutex
}

func (p *Publisher) Uid() string {
	return p.uid
}

func (p *Publisher) Process(event sirkeji.Event) {
	if event.Type == events.Number {
		p.Lock()
		defer p.Unlock()

		p.count++
	}
}

func (p *Publisher) Subscribed() {
	go func() {
		for range time.Tick(5 * time.Second) {
			p.RLock()
			p.publish(sirkeji.Event{
				Publisher: p.uid,
				Type:      events.NumberCountUpdate,
				Meta:      strconv.Itoa(p.count),
				Payload:   p.count,
			})
			p.RUnlock()
		}
	}()
}

func (p *Publisher) Unsubscribed() {}

func NewPublisher(uid string, publish func(e sirkeji.Event)) *Publisher {
	return &Publisher{
		uid:     uid,
		publish: publish,
	}
}
