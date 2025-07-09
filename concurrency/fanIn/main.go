package main

import (
	fanIn "FanIn/fanIn"
	"fmt"
)

func main() {
	ch1 := fanIn.Generator(1)
	ch2 := fanIn.Generator(2)
	ch3 := fanIn.Generator(3)
	ch4 := fanIn.Generator(4)

	result := fanIn.Merging(ch1, ch2, ch3, ch4)

	for val := range result {
		fmt.Println("Считано из результирующего канала: ", val)
	}
}
