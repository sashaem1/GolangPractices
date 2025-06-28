package stack_and_queue

type Stack struct {
	elements []int
}

func NewStack() Stack {
	return Stack{
		elements: []int{},
	}
}

func NewStackFromValues(values ...int) Stack {
	return Stack{
		elements: values,
	}
}

func (s *Stack) Push(val int) {
	s.elements = append(s.elements, val)
}

func (s *Stack) Pop() {
	s.elements = s.elements[:len(s.elements)-1]
}

func (s *Stack) Get() int {
	result := s.elements[len(s.elements)-1]
	return result
}
