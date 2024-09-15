# GoToolshed 🛠️

**GoToolshed** is a collection of re-usable Go packages that I've built over time. There's no overarching theme to the collection, but these tools are designed to be modular and useful across different projects.

## Available Packages

| Package                | Description                                                                                                   |
| ---------------------- | ------------------------------------------------------------------------------------------------------------- |
| [Stack](./stack/README.md)  🥞 | A thread-safe generic stack implementation in Go. Allows any type, with a `Contains` method that only works for comparable types. |

More packages will be added as they are developed!

## Package Details

### 🥞 Stack Package
The stack package is a thread-safe implementation of a generic stack. You can use it to push, pop, and peek at elements, in a thread safe way.

### Example

Here's a quick example of how to use the `Stack` package:

```go
package main

import (
    "fmt"
    "github.com/yourusername/gotoolshed/stack"
)
func printSquare(x int) {
  fmt.Println(x * x)
}

func main() {
    maxSize := 16
    s := stack.New[int](maxSize)  // Create a new thread-safe stack of ints
    s.Push(10)
    s.Push(20)
    s.Push(30)
    fmt.Println(s.Pop())       // 30
    fmt.Println(s.Peek())      // 20
    fmt.Println(s.Size())      // 2
    fmt.Println(s.IsFull())    // false
    fmt.Println(s.Elements())  // [2, 20]
    s.Traverse(printSquare)    // 4
                               // 40
    s.Clear()
    fmt.Println(s.IsEmpty())   // true
}
```
