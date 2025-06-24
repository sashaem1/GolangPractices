package split_slice

func SplitSlice(slice []int, n int) [][]int {
	result := make([][]int, 0)

	for i := 0; len(slice) > i; i += n {
		if len(slice)-i >= n {
			result = append(result, slice[i:i+n])
		} else {
			result = append(result, slice[i:])
		}
	}
	return result
}
