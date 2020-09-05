package colornames_test

import (
	"fmt"

	"github.com/takumakei/go-colornames"
	"github.com/takumakei/go-colornames/color"
)

func ExampleNearest() {
	c := color.RGB([]byte{0x33, 0x44, 0x55})
	name, rgb := colornames.Nearest(c)
	fmt.Println(rgb.Hex(), name)
	// Output:
	// 314459 Pickled Bluewood
}
