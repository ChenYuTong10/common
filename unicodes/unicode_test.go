package unicodes

import (
	"testing"
)

type IndexTest struct {
	s   string
	sep string
	out int
}

var indexTests = []IndexTest{
	{"", "", 0},
	{"", "a", -1},
	{"", "中", -1},
	{"", "1", -1},
	{"", "abc", -1},
	{"", "中文", -1},
	{"", "123", -1},
	{"abc中文123", "bc", 1},
	{"abc中文123", "中文", 3},
	{"abc中文123", "23", 6},
	// special character
	{"abc�文123", "文", 4},
	{"abc中☹123", "☹", 4},
}

func TestIndex(t *testing.T) {
	for _, test := range indexTests {
		actual := Index(test.s, test.sep)
		if actual != test.out {
			t.Errorf("Index(%v,%v) = %v; want %v", test.s, test.sep, actual, test.out)
		}
	}
}

type LenTest struct {
	s   string
	out int
}

var lenTests = []LenTest{
	{"", 0},
	{"abc", 3},
	{"123", 3},
	{"大家好", 3},
	{"abc123", 6},
	{"abc大家好", 6},
	{"123大家好", 6},
	// special character
	{"abc�", 4},
	{"abc☹", 4},
	{"大家��", 4},
}

func TestLen(t *testing.T) {
	for _, test := range lenTests {
		actual := Len(test.s)
		if actual != test.out {
			t.Errorf("Len(%v) = %v; want %v", test.s, actual, test.out)
		}
	}
}

type CutTest struct {
	s          string
	start, end int
	out        string
}

var cutTests = []CutTest{
	{"", 0, 0, ""},
	{"", 1, 4, ""},
	{"", 4, 1, ""},
	{"abcdefghijklmn", 0, 0, ""},
	{"abcdefghijklmn", 10, 0, ""},
	{"abcdefghijklmn", 0, 4, "abcd"},
	{"abcdefghijklmn", 4, 8, "efgh"},
	{"abcdefghijklmn", 0, 26, "abcdefghijklmn"},
	{"大家好，我是小明", 0, 0, ""},
	{"大家好，我是小明", 4, 0, ""},
	{"大家好，我是小明", 0, 3, "大家好"},
	{"大家好，我是小明", 4, 8, "我是小明"},
	{"大家好，我是小明", 0, 10, "大家好，我是小明"},
	// special characters
	{"大家�，我是小明", 0, 2, "大家"},
	{"大家�，我是小明", 4, 8, "我是小明"},
}

func TestCut(t *testing.T) {
	for _, test := range cutTests {
		actual := Cut(test.s, test.start, test.end)
		if actual != test.out {
			t.Errorf("Cut(%v, %v, %v) = %v; want %v", test.s, test.start, test.end, actual, test.out)
		}
	}
}
