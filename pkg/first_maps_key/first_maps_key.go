// Take 2 first keys from map (task 4)
package first_maps_key

func CombineMapsFirstKeys(m map[string]bool) map[string]bool {
	i := 0
	prevKey := ""
	for key := range m {
		i++
		if i == 2 {
			if _, ok := m[prevKey+key]; !ok {
				m[prevKey+key] = true
			}
			return m
		}
		prevKey = key
	}
	return m
}
