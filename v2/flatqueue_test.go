package flatqueue

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

const (
	benchN int = 100_000
	benchK int = benchN / 1_000
)

func TestMaintainsPriorityQueue(t *testing.T) {
	n := 10000
	var q FlatQueue[int, float64]

	data := make([]float64, n)
	sorted := make([]float64, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Float64()
		sorted[i] = data[i]
	}

	sort.Float64s(sorted)

	for i, v := range data {
		q.Push(i, v)
	}

	if q.PeekValue() != sorted[0] {
		t.Fatal()
	}

	if data[q.Peek()] != sorted[0] {
		t.Fatal()
	}

	result := make([]float64, n)
	i := 0
	for q.Len() > 0 {
		result[i] = data[q.Pop()]
		i++
	}

	if !reflect.DeepEqual(sorted, result) {
		t.Fatal()
	}
}

func TestLen(t *testing.T) {
	var q FlatQueue[int, float64]

	if q.Len() != 0 {
		t.Fatal()
	}

	if q.Push(0, 0); q.Len() != 1 {
		t.Fatal()
	}

	if q.Push(1, 1); q.Len() != 2 {
		t.Fatal()
	}

	if q.Pop(); q.Len() != 1 {
		t.Fatal()
	}

	if q.Pop(); q.Len() != 0 {
		t.Fatal()
	}
}

func TestPop(t *testing.T) {
	var q FlatQueue[int, float64]

	q.Push(1, 10)
	q.Push(2, 11)

	if q.Pop() != 1 {
		t.Fatal()
	}

	if q.Pop() != 2 {
		t.Fatal()
	}
}

func TestPeek(t *testing.T) {
	var q FlatQueue[int, float64]

	q.Push(1, 10)

	if q.Peek() != 1 {
		t.Fatal()
	}

	q.Push(2, 11)

	if q.Peek() != 1 {
		t.Fatal()
	}

	q.Push(3, 9)

	if q.Peek() != 3 {
		t.Fatal()
	}

	q.Pop()

	if q.Peek() != 1 {
		t.Fatal()
	}

	q.Pop()
	q.Pop()
}

func TestPeekValue(t *testing.T) {
	var q FlatQueue[int, float64]

	q.Push(1, 10)

	if q.PeekValue() != float64(10) {
		t.Fatal()
	}

	q.Push(2, 11)

	if q.PeekValue() != float64(10) {
		t.Fatal()
	}

	q.Push(3, 9)

	if q.PeekValue() != float64(9) {
		t.Fatal()
	}

	q.Pop()

	if q.PeekValue() != float64(10) {
		t.Fatal()
	}

	q.Pop()
	q.Pop()
}

func TestEdgeCasesWithFewElements(t *testing.T) {
	var q FlatQueue[int, float64]

	q.Push(0, 2)
	q.Push(1, 1)
	q.Pop()
	q.Pop()
	q.Push(2, 2)
	q.Push(3, 1)

	if q.Pop() != 3 {
		t.Fatal()
	}

	if q.Pop() != 2 {
		t.Fatal()
	}
}

func TestPeekEmpty(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal()
		}
	}()
	var q FlatQueue[int, float64]
	q.Peek()
}

func TestPeekValueEmpty(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal()
		}
	}()
	var q FlatQueue[int, float64]
	q.PeekValue()
}

func TestPopEmpty(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal()
		}
	}()
	var q FlatQueue[int, float64]
	q.Pop()
}

func TestClear(t *testing.T) {
	var q FlatQueue[int, float64]

	q.Clear() // ok to clear empty queue

	q.Push(1, 1)
	q.Push(2, 2)

	if q.Len() != 2 {
		t.Fatal()
	}

	q.Clear()

	if q.Len() != 0 {
		t.Fatal()
	}

	q.Push(3, 3)
	q.Push(4, 4)

	if q.Pop() != 3 {
		t.Fatal()
	}

	if q.Pop() != 4 {
		t.Fatal()
	}

	if q.Len() != 0 {
		t.Fatal()
	}
}

func BenchmarkPush(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()

		var q FlatQueue[int, float64]

		values := make([]float64, benchN)
		for j := 0; j < benchN; j++ {
			values[j] = rand.Float64()
		}

		b.StartTimer()

		// Fill the queue
		for j, value := range values {
			q.Push(j, value)
		}
	}
}

func BenchmarkPop(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()

		var q FlatQueue[int, float64]
		for j := 0; j < benchN; j++ {
			q.Push(j, rand.Float64())
		}

		b.StartTimer()

		// Empty the queue
		for q.Len() > 0 {
			q.Pop()
		}
	}
}

func BenchmarkPushPop(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()

		var q FlatQueue[int, float64]

		values := make([]float64, benchN)
		for i := 0; i < benchN; i++ {
			values[i] = rand.Float64()
		}

		b.StartTimer()

		for j := 0; j < benchN; j += benchK {
			for k := 0; k < benchK; k++ {
				q.Push(j+k, values[j+k])
			}
			for k := 0; k < benchK; k++ {
				q.Pop()
			}
		}
	}
}
