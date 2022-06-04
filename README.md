## flatqueue-go [![Tests](https://github.com/invisiblefunnel/flatqueue-go/actions/workflows/go.yml/badge.svg)](https://github.com/invisiblefunnel/flatqueue-go/actions/workflows/go.yml)

A Go 1.18+ port of the [flatqueue](https://github.com/mourner/flatqueue) priority queue library using generics.

`Peek`, `PeekValue`, and `Pop` will panic if called on an empty queue. You must check `Len` accordingly.

```go
import (
    "github.com/invisiblefunnel/flatqueue-go/v2"
)

type Item struct {
    Name  string
    Value int
}

func main() {
    items := []Item{
        {"X", 5},
        {"Y", 2},
        {"Z", 3},
    }

    var q flatqueue.FlatQueue[Item, int]

    for _, item := range items {
        q.Push(item, item.Value)
    }

    var (
        item  Item
        value int
    )

    item = q.Peek()       // top item
    value = q.PeekValue() // top item value
    item = q.Pop()        // remove and return the top item

    // loop while the queue is not empty
    for q.Len() > 0 {
        q.Pop()
    }
}
```
