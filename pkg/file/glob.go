package file

import (
	"github.com/gobwas/glob"
)

// NewGlob compiles the given pattern
func NewGlob(pattern string) glob.Glob {
	return glob.MustCompile(pattern)
}

// Transform compiles all given patterns to globs and maps them
func Transform(patterns *[]string) []glob.Glob {
	globs := make([]glob.Glob, len(*patterns))

	for i, e := range *patterns {
		globs[i] = NewGlob(e)
	}

	return globs
}
