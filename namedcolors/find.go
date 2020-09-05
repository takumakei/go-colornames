package namedcolors

import (
	"github.com/sahilm/fuzzy"
	"github.com/takumakei/go-colornames"
)

// Find returns the color.
func Find(s string) (c Color, ok bool) {
	m := FindNames(s)
	if len(m) == 0 {
		return
	}
	c.Name = m[0].Str
	c.RGB = colornames.RGB(m[0].Index)
	ok = true
	return
}

// FindNames returns the maches from Source.
func FindNames(s string) fuzzy.Matches {
	return fuzzy.FindFrom(s, Source)
}
