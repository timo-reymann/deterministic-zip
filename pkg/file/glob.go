package file

import (
	"github.com/bmatcuk/doublestar/v4"
	"os"
	"path/filepath"
	"sort"
)

// FindByGlob searches files by pattern
func FindByGlob(pattern string) []string {
	path, pattern := doublestar.SplitPattern(pattern)
	fsys := os.DirFS(path)
	matches, _ := doublestar.Glob(fsys, pattern)

	results := make([]string, len(matches))
	for i, m := range matches {
		if m == "." {
			m = ""
		}
		results[i] = path + string(filepath.Separator) + m
	}

	sort.Strings(results)

	return results
}
