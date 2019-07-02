package libenv

// CopyStringMap copies a map and returns the copy
func CopyStringMap(original map[string]string) (copy map[string]string) {
	copy = make(map[string]string)

	for k, v := range original {
		copy[k] = v
	}

	return copy
}