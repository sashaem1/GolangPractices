package maps_union

import (
	"reflect"
	"testing"
)

func TestUnionMaps(t *testing.T) {
	testData := []struct {
		firstDataMap  map[string]int
		secondDataMap map[string]int
		expected      map[string]int
	}{
		{
			firstDataMap: map[string]int{
				"first":  1,
				"second": 2,
				"third":  3,
			},
			secondDataMap: map[string]int{
				"second": 4,
				"third":  6,
				"fourth": 4,
				"fifth":  5,
				"sixth":  6,
			},
			expected: map[string]int{
				"first":  1,
				"second": 4,
				"third":  6,
				"fourth": 4,
				"fifth":  5,
				"sixth":  6,
			},
		},
		{
			firstDataMap: map[string]int{
				"first":  1,
				"second": 2,
				"third":  3,
			},
			secondDataMap: map[string]int{
				"fourth": 4,
				"fifth":  5,
				"sixth":  6,
			},
			expected: map[string]int{
				"first":  1,
				"second": 2,
				"third":  3,
				"fourth": 4,
				"fifth":  5,
				"sixth":  6,
			},
		},
		{
			firstDataMap: map[string]int{
				"first":  -1,
				"second": -2,
				"third":  -3,
			},
			secondDataMap: map[string]int{
				"second": 4,
				"third":  6,
				"fourth": 4,
				"fifth":  5,
				"sixth":  6,
			},
			expected: map[string]int{
				"first":  -1,
				"second": 4,
				"third":  6,
				"fourth": 4,
				"fifth":  5,
				"sixth":  6,
			},
		},
	}

	for _, testCase := range testData {
		testResult := UnionMaps(testCase.firstDataMap, testCase.secondDataMap)

		if !reflect.DeepEqual(testResult, testCase.expected) {
			t.Errorf("Неверный результат. Ожидалось: %v, получено: %v", testCase.expected, testResult)
		}
	}
}
