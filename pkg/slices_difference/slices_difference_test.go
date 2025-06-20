package slices_difference_intersection

import (
	"slices"
	"testing"
)

func TestSlicesDifference(t *testing.T) {
	testData := []struct {
		data     [2][]int
		expected []int
	}{
		{
			data: [2][]int{
				{1, 2, 3, 4, 5},
				{3, 5, 6, 7, 8},
			},
			expected: []int{1, 2, 4},
		},
		{
			data: [2][]int{
				{1, 2, 3},
				{3, 4, 5},
			},
			expected: []int{1, 2},
		},
		{
			data: [2][]int{
				{},
				{},
			},
			expected: []int{},
		},
	}

	for _, testCase := range testData {
		testresult := SlicesDifference(testCase.data[0], testCase.data[1])

		if !slices.Equal(testresult, testCase.expected) {
			t.Errorf("Неверный результат. Ожидалось: %d, получено: %d", testCase.expected, testresult)
		}
	}
}

func TestSlicesIntersection(t *testing.T) {
	testData := []struct {
		data     [2][]int
		expected []int
	}{
		{
			data: [2][]int{
				{1, 2, 3, 4, 5},
				{3, 5, 6, 7, 8},
			},
			expected: []int{3, 5},
		},
		{
			data: [2][]int{
				{1, 2, 3},
				{3, 4, 5},
			},
			expected: []int{3},
		},
		{
			data: [2][]int{
				{},
				{},
			},
			expected: []int{},
		},
	}

	for _, testCase := range testData {
		testresult := SlicesIntersection(testCase.data[0], testCase.data[1])

		if !slices.Equal(testresult, testCase.expected) {
			t.Errorf("Неверный результат. Ожидалось: %d, получено: %d", testCase.expected, testresult)
		}
	}
}
