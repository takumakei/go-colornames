package colornames_test

import (
	"fmt"

	"github.com/takumakei/go-colornames"
	"github.com/takumakei/go-colornames/rgb"
)

func ExampleNearest() {
	c := rgb.RGB([]byte{0x33, 0x44, 0x55})
	name, color := colornames.Nearest(c)
	fmt.Println(color.Hex(), name)
	// Output:
	// 314459 Pickled Bluewood
}
