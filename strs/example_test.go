package strs_test

import (
	"fmt"
	"github.com/ChenYuTong10/common/strs"
)

func ExampleHasEmpty() {
	fmt.Println(strs.HasEmpty("foo", "bar"))
	fmt.Println(strs.HasEmpty("foo", ""))
	fmt.Println(strs.HasEmpty(""))
	fmt.Println(strs.HasEmpty())
	// Output:
	// false
	// true
	// true
	// true
}

func ExampleCut() {
	fmt.Println(strs.Cut("2022 BBB Golang TTT It is good", "BBB", "TTT"))
	fmt.Println(strs.Cut("2022 BBB Golang TTT It is good", "BBB", ""))
	fmt.Println(strs.Cut("2022 Golang It is good", "BBB", "TTT"))
	// Output:
	// "2022", "Golang", "It is good"
	// "", "", ""
	// "", "", ""
}
