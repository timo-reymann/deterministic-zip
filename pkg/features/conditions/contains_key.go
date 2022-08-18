package conditions

func ContainsKey(haystack *map[string]string, needle string) bool {
	if _, ok := (*haystack)[needle]; ok {
		return ok
	}
	return false
}
