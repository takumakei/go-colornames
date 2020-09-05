package colorname

import "github.com/takumakei/go-colornames"

// Source is the source of the color names.
var Source ColorSource

// ColorSource represents the source of the color names.
type ColorSource struct{}

func (ColorSource) String(i int) string { return colornames.Get(i) }
func (ColorSource) Len() int            { return colornames.N }
