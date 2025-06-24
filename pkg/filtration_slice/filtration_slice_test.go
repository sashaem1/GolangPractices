package filtration_slice

import (
	"slices"
	"testing"
)

func TestFiltrationSlice(t *testing.T) {
	testData := []struct {
		data     []int
		filter   func(int) bool
		expected []int
	}{
		{
			data: []int{1, 2, 3, 4, 5, 6},
			filter: func(val int) bool {
				return val%2 == 0
			},
			expected: []int{2, 4, 6},
		},
		{
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			filter: func(val int) bool {
				return val < 7
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, testCase := range testData {
		testResult := FiltrationSlice(testCase.data, testCase.filter)

		slices.Sort(testResult)

		if !slices.Equal(testResult, testCase.expected) {
			t.Errorf("Неверный результат. Ожидалось: %d, получено: %d", testCase.expected, testResult)
		}
	}
}
