package stack_and_queue

import (
	"slices"
	"testing"
)

func TestQueue(t *testing.T) {
	testData := []struct {
		data     Queue
		expected Queue
	}{
		{
			NewQueueFromValues(1, 2, 3, 4, 5),
			NewQueueFromValues(3, 4, 5, 6, 7),
		},
	}

	for _, testCase := range testData {

		testCase.data.Pop()
		testCase.data.Pop()
		testCase.data.Push(6)
		testCase.data.Push(7)

		if !slices.Equal(testCase.data.elements, testCase.expected.elements) {
			t.Errorf("Неверный результат, получено: %v, ожидалось %v", testCase.data.elements, testCase.expected.elements)
		}
	}

}
