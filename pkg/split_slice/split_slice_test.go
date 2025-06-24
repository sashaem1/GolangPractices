package split_slice

import (
	"reflect"
	"testing"
)

func TestSplitSlice(t *testing.T) {
	testData := []struct {
		dataSlice []int
		dataN     int
		expected  [][]int
	}{
		{
			dataSlice: []int{2, 3, 1, 5, 6, 8, 10, 7},
			dataN:     2,
			expected: [][]int{
				{2, 3},
				{1, 5},
				{6, 8},
				{10, 7},
			},
		},
		{
			dataSlice: []int{2, 3, 1, 5, 6, 8, 10, 7},
			dataN:     3,
			expected: [][]int{
				{2, 3, 1},
				{5, 6, 8},
				{10, 7},
			},
		},
	}

	for _, testCase := range testData {
		testResult := SplitSlice(testCase.dataSlice, testCase.dataN)

		if !reflect.DeepEqual(testResult, testCase.expected) {
			t.Errorf("Неверный результат. Ожидалось: %v, получено: %v", testCase.expected, testResult)
		}
	}
}
