```go
package main

import (

    "fmt"
)

func main() {
    var m = make(map[string]int)
    m["a"] = 1
    m["b"] = 2

    value, ok := m["c"]
    if ok {
        fmt.Println("Value found:", value)
    } else {
        fmt.Println("Key not found")
    }
}
```