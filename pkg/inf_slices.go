// infinite slices task 1
package inf_slices

func CombineSlices(slices ...[]int) []int {
	result := make([]int, 0)
	for _, temp_slice := range slices {
		result = append(result, temp_slice...)
	}
	return result
}
