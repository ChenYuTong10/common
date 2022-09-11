package slices_test

import (
	"fmt"
	"github.com/ChenYuTong10/common/slices"
)

func ExampleSplitN() {
	ss := []string{
		"foo", "bar",
		"foo1", "bar1",
		"foo2", "bar2",
	}
	fmt.Println(slices.SplitN(ss, 2))
	fmt.Println(slices.SplitN(ss, 8))
	fmt.Println(slices.SplitN(ss, 0))
	fmt.Println(slices.SplitN(ss, -1))
	// Output:
	// [["foo", "bar"] ["foo1", "bar1"] ["foo2", "bar2"]]
	// [["foo", "bar", "foo1", "bar1", "foo2", "bar2"]]
	// [["foo", "bar", "foo1", "bar1", "foo2", "bar2"]]
	// [["foo", "bar", "foo1", "bar1", "foo2", "bar2"]]
}
