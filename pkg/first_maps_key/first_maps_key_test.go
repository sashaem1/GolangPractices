package first_maps_key

import "testing"

func TestCombineMapsFirstKeys(t *testing.T) {
	testData := []struct {
		data     map[string]bool
		expected []string
	}{
		{
			data: map[string]bool{
				"first":  true,
				"second": true,
			},
			expected: []string{
				"firstsecond",
				"secondfirst",
			},
		},
		{
			data: map[string]bool{
				"first":  true,
				"second": true,
				"third":  true,
			},
			expected: []string{
				"firstsecond",
				"firstthird",
				"secondfirst",
				"secondthird",
				"thirdfirst",
				"thirdsecond",
			},
		},
	}

	for _, testCase := range testData {
		callResult := CombineMapsFirstKeys(testCase.data)
		testResult := false
		for _, val := range testCase.expected {
			if callResult[val] {
				testResult = true
			}
		}
		if !testResult {
			t.Errorf("Неверный результат. Отсутствует комбинированный ключ, получено: %v", callResult)
		}
	}
}
