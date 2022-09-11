package strs

import "testing"

type EmptyTest struct {
	ss  []string
	out bool
}

var emptyTests = []EmptyTest{
	{[]string{}, true},
	{[]string{""}, true},
	{[]string{"", "", ""}, true},
	{[]string{"", "", ""}, true},
	{[]string{"1"}, false},
	{[]string{"a"}, false},
	{[]string{"中"}, false},
	{[]string{"123", "", ""}, true},
	{[]string{"123", "abc", ""}, true},
	{[]string{"123", "", "中文"}, true},
	{[]string{"", "abc", "中文"}, true},
	{[]string{"123", "abc", "中文"}, false},
}

func TestHasEmpty(t *testing.T) {
	for _, test := range emptyTests {
		actual := HasEmpty(test.ss...)
		if actual != test.out {
			t.Errorf("HasEmpty %v Error: actual is %v, want %v", test.ss, actual, test.out)
		}
	}
}

type CutTest struct {
	s, prefix, suffix      string
	before, middle, behind string
}

var cutTests = []CutTest{
	{s: "", prefix: "", suffix: "", before: "", middle: "", behind: ""},
	{s: "", prefix: "xxx", suffix: "yyy", before: "", middle: "", behind: ""},
	{s: "BeforeaaaMiddlebbbBehind", prefix: "", suffix: "yyy", before: "", middle: "", behind: ""},
	{s: "BeforeaaaMiddlebbbBehind", prefix: "xxx", suffix: "", before: "", middle: "", behind: ""},
	{s: "BeforeaaaMiddlebbbBehind", prefix: "", suffix: "", before: "", middle: "", behind: ""},
	{s: "BeforexxxMiddleyyyBehind", prefix: "xxx", suffix: "yyy", before: "Before", middle: "Middle", behind: "Behind"},
	{s: "BeforeaaaMiddlebbbBehind", prefix: "xxx", suffix: "yyy", before: "", middle: "", behind: ""},
}

func TestCut(t *testing.T) {
	for _, test := range cutTests {
		before, middle, behind := Cut(test.s, test.prefix, test.suffix)
		if before != test.before ||
			middle != test.middle ||
			behind != test.behind {
			t.Errorf("Cut(%v, %v, %v) = %v, %v ,%v, want %v, %v, %v",
				test.s, test.prefix, test.suffix,
				before, middle, behind,
				test.before, test.middle, test.behind,
			)
		}
	}
}
