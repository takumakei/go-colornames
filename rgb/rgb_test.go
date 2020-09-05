package rgb

import (
	"fmt"
	"testing"
)

func ExampleRGB() {
	c := RGB([]byte{0x12, 0x34, 0x56})
	fmt.Println(c.RGBA())
	fmt.Println(c.Hex())
	// Output:
	// 4626 13364 22102 65535
	// 123456
}

func ExampleRGB_zero() {
	var zero RGB
	fmt.Println(zero.RGBA())
	fmt.Println(zero.Hex())
	// Output:
	// 0 0 0 0
	// 000000
}

func BenchmarkRGB_Hex(b *testing.B) {
	b.Run("Hex", func(b *testing.B) {
		c := RGB([]byte{0x12, 0x34, 0x56})
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = c.Hex()
		}
	})

	b.Run("SlowHex", func(b *testing.B) {
		c := RGB([]byte{0x12, 0x34, 0x56})
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = c.SlowHex()
		}
	})
}

func (rgb RGB) SlowHex() string {
	v := []uint8(rgb)
	if len(v) < 3 {
		return "000000"
	}
	return fmt.Sprintf("%02x%02x%02x", v[0], v[1], v[2])
}
