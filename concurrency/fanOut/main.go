package main

import (
	"FanOut/fanOut"
	"sync"
)

func main() {
	testCh := fanOut.Generator(4)

	channels := fanOut.Distributor(testCh, 4)

	wg := &sync.WaitGroup{}

	fanOut.Reader(channels, wg)
	wg.Wait()
}
