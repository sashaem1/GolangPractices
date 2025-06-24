package maps_union

func UnionMaps(firstMap, secondMap map[string]int) map[string]int {
	result := make(map[string]int)

	for key, val := range firstMap {
		result[key] = val
	}

	for key, val := range secondMap {
		if _, ok := result[key]; !ok || result[key] < val {
			result[key] = val
		}
	}
	return result
}
