package unicodes

import "strings"

// Index returns index of the first unicode substr in s, or -1 if substr is not presented in s.
func Index(s, substr string) int {
	index := strings.Index(s, substr)
	if index != -1 {
		before := s[0:index]
		return len([]rune(before))
	}
	return -1
}

// Len returns the length of unicode s.
func Len(s string) int {
	return len([]rune(s))
}

// Cut returns the substr cut from unicode s between start and end.
// Substr includes the start but except end.
//
// 1. start >= end, an empty string will be returned.
//
// 2. end > Len(s), end will set with the Len(s).
func Cut(s string, start, end int) string {
	if start >= end {
		return ""
	}

	if start < 0 {
		start = 0
	}
	length := Len(s)
	if end > length {
		end = length
	}

	b := strings.Builder{}
	unicode := []rune(s)
	for i := start; i < end; i++ {
		b.WriteString(string(unicode[i]))
	}
	return b.String()
}
