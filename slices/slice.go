package slices

// SplitN splits the native slice into different pieces with length.
func SplitN[T any](native []T, length int) [][]T {
	if length <= 0 {
		return [][]T{native}
	}
	if len(native) <= length {
		return [][]T{native}
	}

	total := len(native) / length
	if len(native)%length != 0 {
		total++
	}
	splits := make([][]T, total)

	start, end := 0, 0
	for i := 0; i < total-1; i++ {
		start = i * length
		end = start + length
		splits[i] = native[start:end]
	}
	splits[total-1] = native[end:]
	return splits
}
