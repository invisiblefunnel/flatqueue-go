package flatqueue

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

const benchmarkN int = 100_000

func TestMaintainsPriorityQueue(t *testing.T) {
	n := 10000
	q := New[int, float64]()

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
	q := New[int, float64]()

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
	q := New[int, float64]()

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
	q := New[int, float64]()

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
	q := New[int, float64]()

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
	q := New[int, float64]()

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
	New[int, float64]().Peek()
}

func TestPeekValueEmpty(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal()
		}
	}()
	New[int, float64]().PeekValue()
}

func TestPopEmpty(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal()
		}
	}()
	New[int, float64]().Pop()
}

func BenchmarkPush(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	var (
		q      *FlatQueue[int, float64]
		values []float64 = make([]float64, benchmarkN)
	)

	for i := 0; i < b.N; i++ {
		b.StopTimer()

		q = New[int, float64]()

		for j := 0; j < benchmarkN; j++ {
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

	var q *FlatQueue[int, float64]

	for i := 0; i < b.N; i++ {
		b.StopTimer()

		q = New[int, float64]()

		for j := 0; j < benchmarkN; j++ {
			q.Push(j, rand.Float64())
		}

		b.StartTimer()

		// Empty the queue
		for q.Len() > 0 {
			q.Pop()
		}
	}
}
