// infinite slices (task 1)
package inf_slices

import "sort"

func CombineSlices(slices ...[]int) []int {
	result := make([]int, 0)
	for _, temp_slice := range slices {
		result = append(result, temp_slice...)
	}
	sort.Ints(result)
	return result
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
