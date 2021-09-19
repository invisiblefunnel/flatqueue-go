## flatqueue-go [![Tests](https://github.com/invisiblefunnel/flatqueue-go/actions/workflows/go.yml/badge.svg)](https://github.com/invisiblefunnel/flatqueue-go/actions/workflows/go.yml)

A Go port of the [flatqueue](https://github.com/mourner/flatqueue) priority queue library. Push items by identifier (`int`) and value (`float64`).

`Peek`, `PeekValue`, and `Pop` will panic if called on an empty queue. You must check `Len` accordingly.

```go
q := flatqueue.New()

for i, item := range items {
    q.Push(i, item.Value)
}

var (
    id    int
    value float64
)

id = q.Peek()         // top item index
value = q.PeekValue() // top item value
id = q.Pop()          // remove and return the top item index

// loop while queue is not empty
for q.Len() > 0 {
    q.Pop()
}
```

Specifying an initial capacity for the underlying slices may improve the performance of `Push` if you know, or can estimate, the maximum length of the queue. This does not limit the length of the queue.

```go
q := flatqueue.NewWithCapacity(len(items))

for i, item := range items {
    q.Push(i, item.Value)
}
```
