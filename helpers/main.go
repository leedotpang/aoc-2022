package helpers

type Stack[T any] []T

func (m *Stack[T]) Push(i T) {
	(*m) = append(*m, i)
}

func (m *Stack[T]) Pop() T {
	last := (*m)[len(*m)-1]
	*m = (*m)[:len((*m))-1]
	return last
}

func (m *Stack[T]) Shift() T {
	first := (*m)[0]
	*m = (*m)[1:]
	return first
}

func (m *Stack[T]) Present() bool {
	return len(*m) > 0
}
