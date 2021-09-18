## flatqueue-go [![Tests](https://github.com/invisiblefunnel/flatqueue-go/actions/workflows/go.yml/badge.svg)](https://github.com/invisiblefunnel/flatqueue-go/actions/workflows/go.yml)

A Go port of the [flatqueue](https://github.com/mourner/flatqueue) priority queue library. Push items by identifier (`int`) and value (`float64`).

```go
q := flatqueue.New()

for i, item := range items {
    q.Push(i, item.Value)
}

var (
    id    int
    value float64
    ok    bool
)

id, ok = q.Peek()         // top item index, ok if not empty
value, ok = q.PeekValue() // top item value, ok if not empty
id, ok = q.Pop()          // remove and return the top item index, ok if not empty

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
