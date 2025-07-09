// Написал по видео с ютуба для практики работы с каналами, по идее реализовывает пайплайн
package reader_doubler

import (
	"fmt"
	"sync"
	"time"
)

func Reader(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for val := range ch {
			fmt.Println("val = ", val)
		}
		defer wg.Done()
	}()
}

func Doubler(ch chan int) chan int {
	resultCh := make(chan int)
	go func() {
		for val := range ch {
			resultCh <- val * 2
			time.Sleep(500 * time.Millisecond)
		}
		close(resultCh)
	}()

	return resultCh
}

func Writer() chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i + 1
		}
		close(ch)
	}()

	return ch
}
