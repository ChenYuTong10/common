package strs

import (
	"strings"
)

// HasEmpty returns true if any string in ss is null or nothing in ss.
func HasEmpty(ss ...string) bool {
	if len(ss) == 0 {
		return true
	}
	i := 0
	for i = 0; i < len(ss) && len(ss[i]) > 0; i++ {
	}
	return i < len(ss)
}

// Cut slices and returns the text before the prefix,
// between the prefix and suffix and behind the suffix.
//
// Of course, you can also match three sections using regular expression.
//
// Notice here are two conditions that Cut returns three null strs:
//
// 1. Prefix or suffix does not appear in s.
//
// 2. Prefix or suffix is null string.
func Cut(s, prefix, suffix string) (string, string, string) {
	if len(prefix) == 0 || len(suffix) == 0 {
		return "", "", ""
	}
	before, rest, found := strings.Cut(s, prefix)
	if !found {
		return "", "", ""
	}
	middle, behind, found := strings.Cut(rest, suffix)
	if !found {
		return before, "", ""
	}
	return before, middle, behind
}
