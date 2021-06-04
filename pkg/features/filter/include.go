package filter

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features/conditions"
	"github.com/timo-reymann/deterministic-zip/pkg/file"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
)

type Include struct {
}

// DebugName prints the debuggable name
func (i Include) DebugName() string {
	return "Include"
}

// IsEnabled checks if include patterns are defined
func (i Include) IsEnabled(c *cli.Configuration) bool {
	return conditions.HasElements(&c.Include)
}

// Execute checks all patterns for each file and excludes files that dont match all patterns
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
		} else {
			output.Debugf("%s doesnt match include patterns, skipping", f)
		}
	}

	c.SourceFiles = files
	return nil
}
