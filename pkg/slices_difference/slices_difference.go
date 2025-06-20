package slices_difference_intersection

func SlicesDifference(firstSlice, secondSlice []int) []int {
	result := []int{}
	secondSliceValueMap := ValuesMap(secondSlice)

	for _, val := range firstSlice {
		if !secondSliceValueMap[val] {
			result = append(result, val)
		}
	}

	return result
}

func SlicesIntersection(firstSlice, secondSlice []int) []int {
	result := []int{}
	firstSliceValueMap := ValuesMap(firstSlice)

	for _, val := range secondSlice {
		if firstSliceValueMap[val] {
			result = append(result, val)
		}
	}

	return result
}

func ValuesMap(slice []int) map[int]bool {
	result := make(map[int]bool)
	for _, val := range slice {
		result[val] = true
	}

	return result
}
