package colornames

import (
	"image/color"
	"math"

	"github.com/takumakei/go-colornames/rgb"
)

// Nearest returns the name and RGB of the nearest color to c
func Nearest(c color.Color) (name string, colorOfName rgb.RGB) {
	R, G, B, _ := c.RGBA()
	ni := 0
	nd := uint64(math.MaxUint32)
	for i := 0; i < N; i++ {
		c := RGB(i)
		r, g, b, _ := c.RGBA()
		d := sqDiff(R, r) + sqDiff(G, g) + sqDiff(B, b)
		if d < nd {
			ni = i
			nd = d
			colorOfName = c
		}
	}
	name = Get(ni)
	return
}

// https://github.com/golang/go/blob/01af46f7cc419da19f8a6a444da8f6022c016803/src/image/color/color.go#L314
func sqDiff(x, y uint32) uint64 {
	d := x - y
	return uint64(d*d) >> 2
}
