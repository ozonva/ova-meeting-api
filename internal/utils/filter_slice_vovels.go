package utils

func FilterSlice(original []string) []string {
	result := make([]string, 0)
	filterMap := getFilterAsMapKeys()
	for _, v := range original {
		if _, ok := filterMap[v]; ok {
			continue
		}
		result = append(result, v)
	}
	return result
}

func getFilter() []string {
	return []string{"a", "e", "i", "o", "u", "y"}
}

func getFilterAsMapKeys() map[string]struct{} {
	result := make(map[string]struct{}, 0)
	for _, v := range getFilter() {
		if _, ok := result[v]; ok {
			continue
		}
		result[v] = struct{}{}
	}
	return result
}
