package file

import (
	"path/filepath"
	"sort"
)

// FindByGlob searches files by pattern and returns a sorted list
func FindByGlob(pattern string) []string {
	results, _ := filepath.Glob(pattern)
	sort.Strings(results)
	return results
}
