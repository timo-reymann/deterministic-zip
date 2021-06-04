package conditions

// HasElements verify that the parameter has elements
func HasElements(v *[]string) bool {
	return len(*v) > 0
}
