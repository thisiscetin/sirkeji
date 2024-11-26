package events

import "github.com/thisiscetin/sirkeji"

var (
	Number            sirkeji.EventType = "Number"
	SquaredNumber     sirkeji.EventType = "SquaredNumber"
	NumberCountUpdate sirkeji.EventType = "NumberCountUpdate"
)

func init() {
	sirkeji.RegisterEventTypes(Number, SquaredNumber, NumberCountUpdate)
}
