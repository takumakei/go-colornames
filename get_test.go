package colornames_test

import (
	"fmt"

	"github.com/takumakei/go-colornames"
)

func ExampleGet() {
	i := 42
	n := colornames.Get(i)
	fmt.Println(n)
	// Output:
	// Abra Goldenrod
}

func ExampleRGB() {
	i := 42
	c := colornames.RGB(i)
	fmt.Printf("%q is the color value of %q", c.Hex(), colornames.Get(i))
	// Output:
	// "eec400" is the color value of "Abra Goldenrod"
}

func Example_n() {
	fmt.Println(colornames.N)
	// Output:
	// 27927
}

func Example() {
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
	// Output:
	// ffaaa5 Alright Then I Became a Princess
}
