package duplicate_removal

import (
	"slices"
	"testing"
)

func TestRemovalSliceDuplicate(t *testing.T) {
	testData := []struct {
		data     []int
		expected []int
	}{
		{
			data:     []int{1, 2, 2, 3, 3, 3, 4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			data:     []int{1, 2, 3, 4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			data:     []int{},
			expected: []int{},
		},
	}

	for _, testCase := range testData {
		testResult := RemovalSliceDuplicate(testCase.data)

		slices.Sort(testResult)

		if !slices.Equal(testResult, testCase.expected) {
			t.Errorf("Неверный результат. Ожидалось: %d, получено: %d", testCase.expected, testResult)
		}
	}
}
