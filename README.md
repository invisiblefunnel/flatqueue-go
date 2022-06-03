## flatqueue-go [![Tests](https://github.com/invisiblefunnel/flatqueue-go/actions/workflows/go.yml/badge.svg)](https://github.com/invisiblefunnel/flatqueue-go/actions/workflows/go.yml)

A Go port of the [flatqueue](https://github.com/mourner/flatqueue) priority queue library.

`Peek`, `PeekValue`, and `Pop` will panic if called on an empty queue. You must check `Len` accordingly.

```go
items := []Item{/* ... */}

q := flatqueue.New[Item, float64]()

for _, item := range items {
    q.Push(item, item.Value)
}

var (
    item  Item
    value float64
)

item = q.Peek()       // top item
value = q.PeekValue() // top item value
item = q.Pop()        // remove and return the top item

// loop while queue is not empty
for q.Len() > 0 {
    q.Pop()
}
```

Specifying an initial capacity for the underlying slices may improve the performance of `Push` if you know, or can estimate, the maximum length of the queue. This does not limit the length of the queue.

```go
q := flatqueue.NewWithCapacity[Item, float64](len(items))

for _, item := range items {
    q.Push(item, item.Value)
}
```
