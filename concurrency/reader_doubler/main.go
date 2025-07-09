package main

import (
	rd "reader/reader_doubler"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	rd.Reader(rd.Doubler(rd.Writer()), wg)
	wg.Wait()
}
