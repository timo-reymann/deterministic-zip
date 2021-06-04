package file

import (
	"github.com/bmatcuk/doublestar/v4"
	"path/filepath"
	"strings"
)

func MachStringByGlob(pattern string, path string) (bool, error) {
	val, err := doublestar.PathMatch(pattern, path)
	if err != nil {
		return false, err
	}

	if val {
		return val, nil
	}

	val, err = doublestar.Match(pattern, filepath.Base(path))
	if err != nil {
		return false, err
	}

	// Try with directory match
	if strings.HasSuffix(pattern, "*") {
		return doublestar.PathMatch(pattern+"*", path)
	}

	return val, nil
}
