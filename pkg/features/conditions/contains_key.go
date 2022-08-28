package conditions

// ContiainsKey checks if the provided string (needle) is in the map (haystack)
func ContainsKey(haystack *map[string]string, needle string) bool {
	if _, ok := (*haystack)[needle]; ok {
		return ok
	}
	return false
}
