package utils

// ReverseKeys Function swaps keys and values in an array of strings with string keys
// If original array have duplicated values, function will panic
func ReverseKeys(original map[string]string) map[string]string {
	resultMap := make(map[string]string, len(original))
	for k, v := range original {
		if _, ok := resultMap[v]; ok {
			panic("Non-unique values in original map")
		}
		resultMap[v] = k
	}
	return resultMap
}
