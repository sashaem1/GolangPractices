package list

import (
	"testing"
)

func TestListGet(t *testing.T) {
	testData := []struct {
		data     List
		expected int
	}{
		{
			*NewList(1, 2, 3),
			2,
		},
	}

	for _, testCase := range testData {

		callResult := testCase.data.Get(1)

		if !(callResult == testCase.expected) {
			t.Errorf("Неверный результат, получено: %v, ожидалось %v", callResult, testCase.expected)
		}
	}

}

func TestNewListFillFromSlice(t *testing.T) {
	testData := []struct {
		data     []int
		expected List
	}{
		{
			[]int{1, 2, 3, 4, 5},
			*NewList(1, 2, 3, 4, 5),
		},
	}

	for _, testCase := range testData {

		callResult := NewListFillFromSlice(testCase.data)

		if !IsListEqual(callResult, &testCase.expected) {
			t.Errorf("Неверный результат, получено: %v, ожидалось %v", callResult.ListPrintString(), testCase.expected.ListPrintString())
		}
	}

}
