package namedcolors

import (
	"github.com/sahilm/fuzzy"
	"github.com/takumakei/go-colornames"
	"github.com/takumakei/go-colornames/rgb"
)

// Find looks up the color name flexibly matches to s, returns the name and its color value.
// Find returns an empty string in case of no color matched.
func Find(s string) (name string, color rgb.RGB) {
	for _, e := range FindNames(s) {
		name = e.Name
		color = e.RGB()
		return
	}
	return
}

// FindNames looks up s in color names and returns matches in descending order of match score
func FindNames(s string) Matches {
	f := fuzzy.FindFrom(s, Source)
	m := make(Matches, len(f))
	for i, e := range f {
		m[i] = Match{
			Name:    e.Str,
			Index:   e.Index,
			Matches: e.MatchedIndexes,
			Score:   e.Score,
		}
	}
	return m
}

// Match represents a matched color name
type Match struct {
	// The matched color name
	Name string
	// The index of the named color
	Index int
	// The indexes of matched characters. Useful for highlighting matches.
	Matches []int
	// Score used to rank matches
	Score int
}

// RGB returns the RGB of the color.
func (m *Match) RGB() rgb.RGB {
	return colornames.RGB(m.Index)
}

// Matches is a slice of Match structs
type Matches []Match
