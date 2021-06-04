package features

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features/conditions"
	"github.com/timo-reymann/deterministic-zip/pkg/file"
)

type Include struct {
}

func (i Include) DebugName() string {
	return "Include"
}

func (i Include) IsEnabled(c *cli.Configuration) bool {
	return conditions.HasElements(&c.Include)
}

func (i Include) Execute(c *cli.Configuration) error {
	files := make([]string, 0)
	patterns := file.Transform(&c.Include)

	var fileIncluded bool

	for _, f := range c.SourceFiles {
		fileIncluded = true
		for _, p := range patterns {
			if !p.Match(f) {
				fileIncluded = false
				break
			}
		}

		if fileIncluded {
			files = append(files, f)
		}
	}

	c.SourceFiles = files
	return nil
}
