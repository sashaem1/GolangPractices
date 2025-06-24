package mirror_map

func MirrorMap(baseMap map[string]int) map[int][]string {
	result := make(map[int][]string)

	for key, val := range baseMap {
		result[val] = append(result[val], key)
	}

	return result
}
