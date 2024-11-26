# sirkeji

## numbers example

This example will use sirkeji to demonstrate,
- creating components with their lifecycle methods
- creating a component with an internal state
- subscribing components to the streamer
- subscribing to the streamer with same component multiple times
- defining events in a central sub-package

```go
func main() {
	sirkeji.Subscribe(gStreamer, sirkeji.NewLogger())
	sirkeji.Subscribe(gStreamer, number.NewPublisher("number-publisher-1", gStreamer.Publish))
	sirkeji.Subscribe(gStreamer, number.NewPublisher("number-publisher-2", gStreamer.Publish))
	sirkeji.Subscribe(gStreamer, squared_number.NewPublisher("squared-number-publisher-1", gStreamer.Publish))
	sirkeji.Subscribe(gStreamer, number_count.NewPublisher("number-count-publisher-1", gStreamer.Publish))

	sirkeji.WaitForTermination(gStreamer)
}
```

Running this package will produce an output like below:
```shell
2024/11/26 11:08:41 [logger-1732608521473] subscribed to the streamer
2024/11/26 11:08:41 [number-publisher-1] subscribed to the streamer
2024/11/26 11:08:41 [number-publisher-2] subscribed to the streamer
2024/11/26 11:08:41 [squared-number-publisher-1] subscribed to the streamer
2024/11/26 11:08:41 [number-count-publisher-1] subscribed to the streamer
2024/11/26 11:08:43.474799 [number-publisher-2] *Number*, m: 646 | pl: full
2024/11/26 11:08:43.475090 [squared-number-publisher-1] *SquaredNumber*, m: 289 | pl: full
2024/11/26 11:08:43.475067 [number-publisher-1] *Number*, m: 17 | pl: full
2024/11/26 11:08:43.475115 [squared-number-publisher-1] *SquaredNumber*, m: 417316 | pl: full
2024/11/26 11:08:45.474784 [number-publisher-2] *Number*, m: 491 | pl: full
2024/11/26 11:08:45.474819 [squared-number-publisher-1] *SquaredNumber*, m: 241081 | pl: full
2024/11/26 11:08:45.474846 [number-publisher-1] *Number*, m: 25 | pl: full
2024/11/26 11:08:45.474886 [squared-number-publisher-1] *SquaredNumber*, m: 625 | pl: full
2024/11/26 11:08:46.474780 [number-count-publisher-1] *NumberCountUpdate*, m: 4 | pl: full
```