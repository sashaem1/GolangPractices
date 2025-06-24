package min_max_slice

import (
	"testing"
)

func TestMinMaxSlice(t *testing.T) {
	testData := []struct {
		data     []int
		expected [2]int
	}{
		{
			data:     []int{2, 3, 1, 5, 6, 8, 10, 7},
			expected: [2]int{1, 10},
		},
		{
			data:     []int{-5, 89, 7, 34, 20},
			expected: [2]int{-5, 89},
		},
		{
			data:     []int{-123, -7, -1, -37},
			expected: [2]int{-123, -1},
		},
	}

	for _, testCase := range testData {
		testResultMin, testResultMax := MinMaxSlice(testCase.data)

		if !(testResultMin == testCase.expected[0] && testResultMax == testCase.expected[1]) {
			t.Errorf("Неверный результат. Ожидалось: %v, получено min:%v max:%v", testCase.expected, testResultMin, testResultMax)
		}
	}
}
