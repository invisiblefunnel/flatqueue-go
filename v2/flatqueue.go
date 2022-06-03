package flatqueue

import (
	"golang.org/x/exp/constraints"
)

type FlatQueue[T any, V constraints.Ordered] struct {
	items  []T
	values []V
	length int
}

func New[T any, V constraints.Ordered]() *FlatQueue[T, V] {
	return &FlatQueue[T, V]{}
}

func (q *FlatQueue[_, _]) Len() int {
	return q.length
}

func (q *FlatQueue[T, V]) Push(item T, value V) {
	pos := q.length
	q.length++

	q.items = append(q.items, item)
	q.values = append(q.values, value)

	for pos > 0 {
		parent := (pos - 1) >> 1
		parentValue := q.values[parent]

		if value > parentValue {
			break
		}

		q.items[pos] = q.items[parent]
		q.values[pos] = parentValue

		pos = parent
	}

	q.items[pos] = item
	q.values[pos] = value
}

func (q *FlatQueue[T, _]) Pop() T {
	top := q.items[0]
	q.length--

	if q.length > 0 {
		id := q.items[q.length]
		value := q.values[q.length]

		q.items[0] = id
		q.values[0] = value

		halfLength := q.length >> 1
		pos := 0

		for pos < halfLength {
			left := (pos << 1) + 1
			right := left + 1

			bestIndex := q.items[left]
			bestValue := q.values[left]
			rightValue := q.values[right]

			if right < q.length && rightValue < bestValue {
				left = right
				bestIndex = q.items[right]
				bestValue = rightValue
			}

			if bestValue > value {
				break
			}

			q.items[pos] = bestIndex
			q.values[pos] = bestValue
			pos = left
		}

		q.items[pos] = id
		q.values[pos] = value
	}

	return top
}

func (q *FlatQueue[T, _]) Peek() T {
	return q.items[0]
}

func (q *FlatQueue[_, V]) PeekValue() V {
	return q.values[0]
}
