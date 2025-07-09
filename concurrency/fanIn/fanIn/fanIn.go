package fanIn

import "sync"

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

func Merging(channels ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}

		for _, ch := range channels {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for val := range ch {
					out <- val
				}
			}()
		}
		wg.Wait()
		close(out)
	}()

	return out
}
