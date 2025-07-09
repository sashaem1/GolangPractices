package cancellation

import (
	"fmt"
	"time"
)

func CancellationExample() {
	stopCh := make(chan bool)

	go WorkImitation(stopCh)

	time.Sleep(3 * time.Second)

	stopCh <- true
}

func WorkImitation(stopCh chan bool) {
	for i := range 20 {
		select {
		case <-stopCh:
			return
		default:
			fmt.Printf("work №%d is completed\n", i)
			time.Sleep(500 * time.Millisecond)
		}
	}

}
