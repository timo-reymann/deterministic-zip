package file

import (
	"github.com/gobwas/glob"
)

func NewGlob(pattern string) glob.Glob {
	return glob.MustCompile(pattern)
}
