package colorname_test

import (
	"encoding/json"
	"fmt"

	"github.com/takumakei/go-colornames/colorname"
)

func Example() {
	if name, rgb := colorname.Find("oceanblue"); name != "" {
		fmt.Println(rgb.Hex(), name)
	}
	// Output:
	// 009dc4 Ocean Blue
}

func ExampleFindNames() {
	m := colorname.FindNames("funpin")
	for _, e := range m {
		b, _ := json.Marshal(e)
		fmt.Println(string(b))
	}
	// Output:
	// {"Name":"Funki Porcini","Index":9745,"Matches":[0,1,2,6,10,11],"Score":88}
	// {"Name":"Fluorescent Pink","Index":9192,"Matches":[0,2,9,12,13,14],"Score":40}
}
