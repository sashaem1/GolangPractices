package duplicate_removal

func RemovalSliceDuplicate(slice []int) []int {
	result := make([]int, 0)
	valMap := ValuesMap(slice)

	for key := range valMap {
		result = append(result, key)
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
