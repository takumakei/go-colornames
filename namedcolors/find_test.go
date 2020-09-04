package namedcolors_test

import (
	"fmt"

	"github.com/takumakei/go-colornames/namedcolors"
)

func Example() {
	if c, ok := namedcolors.Find("oceanblue"); ok {
		fmt.Println(c.Hex(), c.Name)
	}
	// Output:
	// 009dc4 Ocean Blue
}
