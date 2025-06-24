package mirror_map

import (
	"reflect"
	"testing"
)

func TestMirrorMap(t *testing.T) {
	testData := []struct {
		data     map[string]int
		expected map[int][]string
	}{
		{
			data: map[string]int{
				"first":  1,
				"second": 2,
				"third":  3,
				"fourth": 4,
				"fifth":  5,
				"sixth":  6,
			},
			expected: map[int][]string{
				1: {"first"},
				2: {"second"},
				3: {"third"},
				4: {"fourth"},
				5: {"fifth"},
				6: {"sixth"},
			},
		},
		{
			data: map[string]int{
				"first":  1,
				"second": 1,
				"third":  3,
				"fourth": 3,
				"fifth":  3,
				"sixth":  6,
			},
			expected: map[int][]string{
				1: {"first", "second"},
				3: {"third", "fourth", "fifth"},
				6: {"sixth"},
			},
		},
	}

	for _, testCase := range testData {
		testResult := MirrorMap(testCase.data)

		if !reflect.DeepEqual(testResult, testCase.expected) {
			t.Errorf("Неверный результат. Ожидалось: %v, получено: %v", testCase.expected, testResult)
		}
	}
}
