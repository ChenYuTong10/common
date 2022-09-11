package slices

import (
	"strings"
	"testing"
)

var natives = [][]string{
	{},
	{"1", "2", "3"},
	{"1", "2", "3", "4"},
	{"1", " 2", "3", "4", "5"},
	{
		"1", " 2", "3", "4", "5",
		"1",
	},
	{
		"1", " 2", "3", "4", "5",
		"1", "2",
	},
	{
		"1", " 2", "3", "4", "5",
		"1", "2", "3",
	},
	{
		"1", " 2", "3", "4", "5",
		"1", "2", "3", "4",
	},
	{
		"1", " 2", "3", "4", "5",
		"1", "2", "3", "4", "5",
	},
}

func TestSplit(t *testing.T) {
	t.Logf("Splite Size 0:")
	for _, native := range natives {
		t.Logf("Split Result : %v", SplitN[string](native, 0))
	}
	t.Log(strings.Repeat("-", 50))
	t.Logf("Splite Size -1:")
	for _, native := range natives {
		t.Logf("Split Result : %v", SplitN[string](native, -1))
	}
	t.Log(strings.Repeat("-", 50))
	t.Logf("Split Size 5:")
	for _, native := range natives {
		t.Logf("Split Result : %v", SplitN[string](native, 5))
	}
	t.Log(strings.Repeat("-", 50))
	t.Logf("Split Size 4:")
	for _, native := range natives {
		t.Logf("Split Result : %v", SplitN[string](native, 4))
	}
}
