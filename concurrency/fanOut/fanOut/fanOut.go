package fanOut

import (
	"fmt"
	"sync"
)

func Generator(count int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			out <- i + 1
		}
		close(out)
	}()

	return out
}

func Distributor(inCh <-chan int, distrCount int) []<-chan int {
	tempSlice := make([]chan int, distrCount)

	for i := 0; i < distrCount; i++ {
		tempSlice[i] = make(chan int)
	}

	// wg := &sync.WaitGroup{}
	// wg.Add(1)

	go func() {
		// defer wg.Done()
		innerWg := &sync.WaitGroup{}

		for chanVal := range inCh {
			for _, ch := range tempSlice {
				innerWg.Add(1)
				go func(outCh chan int, val int) {
					defer innerWg.Done()
					outCh <- val
				}(ch, chanVal)
			}
		}

		innerWg.Wait()

		for _, ch := range tempSlice {
			close(ch)
		}
	}()

	// wg.Wait()

	// Приводим тип каналов слайса к "только для чтения"
	result := ConvertSliceToReadOnly(tempSlice)

	return result
}

func ConvertSliceToReadOnly(input []chan int) []<-chan int {
	out := make([]<-chan int, len(input))

	for i, ch := range input {
		out[i] = ch // безопасное приведение
	}

	return out
}

func Reader(channels []<-chan int, wg *sync.WaitGroup) {
	for key, ch := range channels {
		wg.Add(1)
		go func(k int, c <-chan int) {
			for val := range c {
				fmt.Printf("Channel №%d, value = %d\n", k, val)
			}
			wg.Done()
		}(key, ch)
	}
}
