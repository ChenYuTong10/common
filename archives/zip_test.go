package archives

import "testing"

type ZipTest struct {
	path    string
	pkgname string
}

var zipTests = []ZipTest{
	{"testdata\\test.zip", "test.zip"},
	{"testdata\\test-chinese.zip", "test-chinese.zip"},
}

func TestUnzip(t *testing.T) {
	for _, test := range zipTests {
		if err := Unzip(test.path); err != nil {
			t.Errorf("%v Unzip Error: %v", test.pkgname, err)
		} else {
			t.Logf("%v Unzip Result: succuss", test.pkgname)
		}
	}
}
