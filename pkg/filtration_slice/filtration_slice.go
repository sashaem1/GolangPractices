package filtration_slice

func FiltrationSlice(slice []int, filter func(int) bool) []int {
	result := []int{}

	for _, val := range slice {
		if filter(val) {
			result = append(result, val)
		}
	}
	return result
}
