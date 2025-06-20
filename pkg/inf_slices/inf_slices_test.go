package inf_slices

import (
	"slices"
	"testing"
)

func TestCombineSlices(t *testing.T) {
	testData := []struct {
		data     [][]int
		expected []int
	}{
		{
			data: [][]int{
				{1},
				{3, 2},
				{6, 4, 5},
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			data: [][]int{
				{},
				{7, 1},
			},
			expected: []int{1, 7},
		},
		{
			data: [][]int{
				{},
				{},
			},
			expected: []int{},
		},
	}

	for _, testCase := range testData {
		result := CombineSlices(testCase.data...)

		if !slices.Equal(result, testCase.expected) {
			t.Errorf("Неверный результат. Ожидалось: %d, получено: %d", testCase.expected, result)
		}
	}
}
