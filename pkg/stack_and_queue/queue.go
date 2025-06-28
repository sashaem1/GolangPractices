package stack_and_queue

type Queue struct {
	elements []int
}

func NewQueue() Queue {
	return Queue{
		elements: []int{},
	}
}

func NewQueueFromValues(values ...int) Queue {
	return Queue{
		elements: values,
	}
}

func (q *Queue) Push(val int) {
	q.elements = append(q.elements, val)
}

func (q *Queue) Pop() {
	q.elements = q.elements[1:]
}

func (q *Queue) Get() int {
	result := q.elements[0]
	return result
}
