package features

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/file"
)

type Exclude struct {
}

func (e Exclude) IsEnabled(c *cli.Configuration) bool {
	return len(c.Exclude) > 0
}

func (e Exclude) Execute(c *cli.Configuration) error {
	files := make([]string, 0)
	var fileExcluded bool

	for _, f := range c.SourceFiles {
		fileExcluded = false
		for _, pattern := range c.Exclude {

			isMatch, err := file.MachStringByGlob(pattern, f)
			if err != nil {
				return err
			}

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
