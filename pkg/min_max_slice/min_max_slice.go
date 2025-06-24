package min_max_slice

import "slices"

func MinMaxSlice(slice []int) (min, max int) {
	min = slices.Min(slice)
	max = slices.Max(slice)

	return min, max
}
