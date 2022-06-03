package flatqueue

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestMaintainsPriorityQueue(t *testing.T) {
	n := 10000
	testMaintainsPriorityQueue(t, n, New[int, float64]())
	testMaintainsPriorityQueue(t, n, NewWithCapacity[int, float64](n))
	testMaintainsPriorityQueue(t, n, NewWithCapacity[int, float64](n/4))
	testMaintainsPriorityQueue(t, n, NewWithCapacity[int, float64](n*2))
}

func testMaintainsPriorityQueue(t *testing.T, n int, q *FlatQueue[int, float64]) {
	data := make([]float64, n)
	sorted := make([]float64, n)
	for i := 0; i < n; i++ {
		data[i] = randFloat64(-100, 100)
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
	testLen(t, New[int, float64]())
	testLen(t, NewWithCapacity[int, float64](100))
}

func testLen(t *testing.T, q *FlatQueue[int, float64]) {
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
	testPop(t, New[int, float64]())
	testPop(t, NewWithCapacity[int, float64](100))
}

func testPop(t *testing.T, q *FlatQueue[int, float64]) {
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
	testPeek(t, New[int, float64]())
	testPeek(t, NewWithCapacity[int, float64](100))
}

func testPeek(t *testing.T, q *FlatQueue[int, float64]) {
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
	testPeekValue(t, New[int, float64]())
	testPeekValue(t, NewWithCapacity[int, float64](100))
}

func testPeekValue(t *testing.T, q *FlatQueue[int, float64]) {
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
	testEdgeCasesWithFewElements(t, New[int, float64]())
	testEdgeCasesWithFewElements(t, NewWithCapacity[int, float64](100))
}

func testEdgeCasesWithFewElements(t *testing.T, q *FlatQueue[int, float64]) {
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
	q := New[int, float64]()

	values := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		values[i] = randFloat64(-1000, 1000)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Push(i, values[i])
	}
}

func BenchmarkPushWithCapacity(b *testing.B) {
	q := NewWithCapacity[int, float64](b.N)

	values := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		values[i] = randFloat64(-1000, 1000)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Push(i, values[i])
	}
}

func BenchmarkPop(b *testing.B) {
	q := New[int, float64]()

	for i := 0; i < b.N; i++ {
		q.Push(i, randFloat64(-1000, 1000))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkPopWithCapacity(b *testing.B) {
	q := NewWithCapacity[int, float64](b.N)

	for i := 0; i < b.N; i++ {
		q.Push(i, randFloat64(-1000, 1000))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func randFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
