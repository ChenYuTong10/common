package unicodes_test

import (
	"fmt"
	"github.com/ChenYuTong10/common/unicodes"
)

func ExampleIndex() {
	fmt.Println(unicodes.Index("012中文abc", "中文"))
	fmt.Println(unicodes.Index("012中文abc", "文中"))
	// Output:
	// 3
	// -1
}

func ExampleLen() {
	fmt.Println(unicodes.Len("01234"))
	fmt.Println(unicodes.Len("abcd"))
	fmt.Println(unicodes.Len("大家好"))
	// Output:
	// 5
	// 4
	// 3
}

func ExampleCut() {
	fmt.Println(unicodes.Cut("012中文abc", 3, 5))
	fmt.Println(unicodes.Cut("012中文abc", 5, 1))
	fmt.Println(unicodes.Cut("012中文abc", 1, 10))
	// Output:
	// "中文"
	// ""
	// "12中文abc"
}
