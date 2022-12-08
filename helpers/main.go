package helpers

type Stack[T any] []T

func (m *Stack[T]) Push(i T) {
	(*m) = append(*m, i)
}

func (m *Stack[T]) Pop() T {
	return (*m)[len(*m)-1]
}

func (m *Stack[T]) Shift() T {
	return (*m)[0]
}

func (m *Stack[T]) Present() bool {
	return len(*m) > 0
}
