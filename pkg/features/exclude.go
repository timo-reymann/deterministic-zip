package features

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features/conditions"
	"github.com/timo-reymann/deterministic-zip/pkg/file"
)

type Exclude struct {
}

func (e Exclude) DebugName() string {
	return "Exclude"
}

func (e Exclude) IsEnabled(c *cli.Configuration) bool {
	return conditions.HasElements(&c.Exclude)
}

func (e Exclude) Execute(c *cli.Configuration) error {
	var fileExcluded bool
	files := make([]string, 0)
	excludes := file.Transform(&c.Exclude)

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
