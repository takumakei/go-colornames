Color Names
======================================================================

The Go package colornames contains the list of color names.

The color names have taken from https://github.com/meodai/color-names.

### example

```go
// Find the longest color name and its color value.
var name string
var index int
for i := 0; i < colornames.N; i++ {
    v := colornames.Get(i)
    if len(v) > len(name) {
        name = v
        index = i
    }
}
c := colornames.RGB(index)
fmt.Println(c.Hex(), name)
```
