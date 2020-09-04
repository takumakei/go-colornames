package colornames

import "github.com/takumakei/go-colornames/color"

// N is the number of the color names in the list.
const N = count

const count = len(_Names_index) - 1

// Get returns the name at index of the list.
//
// The index must be equal or greater than 0 and smaller than colornames.N.
func Get(index int) string {
	return _Names_name[_Names_index[index]:_Names_index[index+1]]
}

// RGB returns the color at index of the list.
//
// The index must be equal or greater than 0 and smaller than colornames.N.
func RGB(index int) color.RGB {
	index *= 3
	return color.RGB(_RGBs[index : index+3])
}

//go:generate go run ./source
