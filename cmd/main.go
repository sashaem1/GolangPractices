package main

import (
	"fmt"

	"github.com/sashaem1/GolangPractices/pkg"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6, 7}
	test1 := pkg.CombineSlices(slice1, slice2)
	fmt.Println(test1)
}
