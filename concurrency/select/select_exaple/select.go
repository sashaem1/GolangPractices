package select_exaple

import (
	"fmt"
	"time"
)

func SelectForMultiplexing() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Massage from gorutine 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Massage from gorutine 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}
