// uniform
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"

	"github.com/takumakei/go-colornames/namedcolors"
)

var (
	FlagWidth  = 320
	FlagHeight = 320
	FlagOutput = "uniform.png"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	flag.IntVar(&FlagWidth, "w", FlagWidth, "width")
	flag.IntVar(&FlagHeight, "h", FlagHeight, "height")
	flag.StringVar(&FlagOutput, "o", FlagOutput, "output PNG image file name")
	flag.Parse()

	q := strings.Join(flag.Args(), " ")
	c, ok := namedcolors.Find(q)
	if !ok {
		return fmt.Errorf("error: color not found for %q", q)
	}
	fmt.Printf("#%s %q\n", c.Hex(), c.Name)

	m := NewUniform(c, image.Rect(0, 0, FlagWidth, FlagHeight))

	f, err := os.Create(FlagOutput)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, m)
}

type Uniform struct {
	*image.Uniform
	bounds image.Rectangle
}

func NewUniform(c color.Color, bounds image.Rectangle) *Uniform {
	return &Uniform{
		Uniform: image.NewUniform(c),
		bounds:  bounds,
	}
}

func (u *Uniform) Bounds() image.Rectangle {
	return u.bounds
}
