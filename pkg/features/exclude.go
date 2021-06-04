package features

import (
	"github.com/gobwas/glob"
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/file"
)

type Exclude struct {
}

func (e Exclude) IsEnabled(c *cli.Configuration) bool {
	return len(c.Exclude) > 0
}

func (e Exclude) Execute(c *cli.Configuration) error {
	var fileExcluded bool
	files := make([]string, 0)
	excludes := make([]glob.Glob, len(c.Exclude))

	for i, e := range c.Exclude {
		excludes[i] = file.NewGlob(e)
	}

	for _, f := range c.SourceFiles {
		fileExcluded = false
		for _, pattern := range excludes {
			isMatch := pattern.Match(f)

			if isMatch {
				fileExcluded = true
				break
			}
		}

		if !fileExcluded {
			files = append(files, f)
		}
	}

	c.SourceFiles = files
	return nil
}
