package main

import (
	"github.com/thisiscetin/sirkeji"
	"github.com/thisiscetin/sirkeji/example/numbers/number"
	"github.com/thisiscetin/sirkeji/example/numbers/number_count"
	"github.com/thisiscetin/sirkeji/example/numbers/squared_number"
)

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
