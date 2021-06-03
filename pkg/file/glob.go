package file

import (
	"path/filepath"
)

// FindByGlob searches files by pattern
func FindByGlob(pattern string) []string {
	results, _ := filepath.Glob(pattern)
	return results
}
