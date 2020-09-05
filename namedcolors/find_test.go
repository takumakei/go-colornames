package namedcolors_test

import (
	"encoding/json"
	"fmt"

	"github.com/takumakei/go-colornames/namedcolors"
)

func Example() {
	if name, rgb := namedcolors.Find("oceanblue"); name != "" {
		fmt.Println(rgb.Hex(), name)
	}
	// Output:
	// 009dc4 Ocean Blue
}

func ExampleFindNames() {
	m := namedcolors.FindNames("funpin")
	for _, e := range m {
		b, _ := json.Marshal(e)
		fmt.Println(string(b))
	}
	// Output:
	// {"Name":"Funki Porcini","Index":9007,"Matches":[0,1,2,6,10,11],"Score":88}
	// {"Name":"Fluorescent Pink","Index":8509,"Matches":[0,2,9,12,13,14],"Score":40}
}
