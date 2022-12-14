package helpers

type Stack[T any] []T
type Queue[T any] []T

func (s *Stack[T]) Push(i T) {
	(*s) = append(*s, i)
}

func (s *Stack[T]) PushMany(list []T) {
	*s = append(*s, list...)
}

func (s *Stack[T]) Pop() T {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len((*s))-1]
	return last
}

func (s *Stack[T]) PopMany(count int) []T {
	last_index := len(*s) - count
	last := (*s)[last_index:]
	*s = (*s)[:last_index]
	return last
}

func (s *Stack[T]) Present() bool {
	return len(*s) > 0
}

func (s *Stack[T]) Reverse() {
	for i, j := len(*s)-1, 0; i >= 0; i, j = i-1, j+1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func (q *Queue[T]) Shift() T {
	first := (*q)[0]
	*q = (*q)[1:]
	return first
}

func (q *Queue[T]) ShiftMany(count int) []T {
	first := (*q)[:count]
	*q = (*q)[count:]
	return first
}

func (q *Queue[T]) Push(i T) {
	(*q) = append(*q, i)
}

func (q *Queue[T]) Present() bool {
	return len(*q) > 0
}
