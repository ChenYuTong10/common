package archives_test

import (
	"fmt"
	"github.com/ChenYuTong10/common/archives"
)

func ExampleUnzip() {
	err := archives.Unzip("testdata/foo.zip")
	if err != nil {
		fmt.Printf("unzip error: %s", err)
		return
	}
	// do anything you want
}
