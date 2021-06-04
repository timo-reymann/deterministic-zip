package file

import (
	"github.com/gobwas/glob"
)

func NewGlob(pattern string) glob.Glob {
	return glob.MustCompile(pattern)
}

func Transform(patterns *[]string) []glob.Glob {
	globs := make([]glob.Glob, len(*patterns))

	for i, e := range *patterns {
		globs[i] = NewGlob(e)
	}

	return globs
}
