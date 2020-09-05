// Package rgb implements the 24-bit color type RGB
package rgb

import (
	"image/color"
)

// RGB represents the color
type RGB []byte

// RGB must be of color.Color
var _ color.Color = RGB{}

// RGBA returns the red, green, blue and alpha values for the color.
// Each value ranges within [0, 0xffff].
func (rgb RGB) RGBA() (r, g, b, a uint32) {
	v := []uint8(rgb)
	if len(v) < 3 {
		return 0, 0, 0, 0
	}
	r = uint32(v[0])
	g = uint32(v[1])
	b = uint32(v[2])
	r |= r << 8
	g |= g << 8
	b |= b << 8
	a = 0xffff
	return
}

// Hex returns hex string.
func (rgb RGB) Hex() string {
	v := []uint8(rgb)
	if len(v) < 3 {
		return "000000"
	}
	var a [6]byte
	a[0] = hex[v[0]>>4]
	a[1] = hex[v[0]&0xf]
	a[2] = hex[v[1]>>4]
	a[3] = hex[v[1]&0xf]
	a[4] = hex[v[2]>>4]
	a[5] = hex[v[2]&0xf]
	return string(a[:])
}

const hex = "0123456789abcdef"
