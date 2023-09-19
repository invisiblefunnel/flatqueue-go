package flatqueue

type FlatQueue struct {
	ids    []int
	values []float64
	length int
}

func New() *FlatQueue {
	return &FlatQueue{}
}

func NewWithCapacity(cap int) *FlatQueue {
	return &FlatQueue{
		ids:    make([]int, 0, cap),
		values: make([]float64, 0, cap),
	}
}

func (q *FlatQueue) Len() int {
	return q.length
}

func (q *FlatQueue) Push(id int, value float64) {
	pos := q.length
	q.length++

	if q.length > len(q.ids) {
		q.ids = append(q.ids, id)
		q.values = append(q.values, value)
	}

	for pos > 0 {
		parent := (pos - 1) >> 1
		parentValue := q.values[parent]

		if value > parentValue {
			break
		}

		q.ids[pos] = q.ids[parent]
		q.values[pos] = parentValue

		pos = parent
	}

	q.ids[pos] = id
	q.values[pos] = value
}

func (q *FlatQueue) Pop() int {
	top := q.ids[0]
	q.length--

	if q.length > 0 {
		id := q.ids[q.length]
		value := q.values[q.length]

		q.ids[0] = id
		q.values[0] = value

		halfLength := q.length >> 1
		pos := 0

		for pos < halfLength {
			left := (pos << 1) + 1
			right := left + 1

			bestIndex := q.ids[left]
			bestValue := q.values[left]
			rightValue := q.values[right]

			if right < q.length && rightValue < bestValue {
				left = right
				bestIndex = q.ids[right]
				bestValue = rightValue
			}

			if bestValue > value {
				break
			}

			q.ids[pos] = bestIndex
			q.values[pos] = bestValue
			pos = left
		}

		q.ids[pos] = id
		q.values[pos] = value
	}

	return top
}

func (q *FlatQueue) Peek() int {
	return q.ids[0]
}

func (q *FlatQueue) PeekValue() float64 {
	return q.values[0]
}

func (q *FlatQueue) Clear() {
	q.length = 0
	q.ids = q.ids[:0]
	q.values = q.values[:0]
}
