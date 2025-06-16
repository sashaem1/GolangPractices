package main

import (
	"fmt"

	"github.com/sashaem1/GolangPractices/pkg/inf_slices"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6, 7}
	test1 := inf_slices.CombineSlices(slice1, slice2)
	fmt.Println(test1)
}
