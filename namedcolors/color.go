package namedcolors

import "github.com/takumakei/go-colornames/rgb"

// Color represents the color.
type Color struct {
	Name string
	rgb.RGB
}
