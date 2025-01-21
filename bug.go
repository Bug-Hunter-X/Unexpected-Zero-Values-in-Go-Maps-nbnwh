```go
func main() {
    var m = make(map[string]int)
    m["a"] = 1
    m["b"] = 2
    fmt.Println(m["c"]) // This will print 0, not an error
}
```