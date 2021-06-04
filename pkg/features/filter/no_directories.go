package filter

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features/conditions"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
	"os"
)

type NoDirectories struct {
}

// DebugName prints the debuggable name
func (n NoDirectories) DebugName() string {
	return "NoDirectories"
}

// IsEnabled checks if no directories should be created
func (n NoDirectories) IsEnabled(c *cli.Configuration) bool {
	return conditions.OnFlag(c.NoDirEntries)
}

// Execute filters out all directories using stat
func (n NoDirectories) Execute(c *cli.Configuration) error {
	files := make([]string, 0)

	for _, f := range c.SourceFiles {
		stat, err := os.Stat(f)
		if err != nil {
			return err
		}

		if !stat.IsDir() {
			files = append(files, f)
		} else {
			output.Debugf("%s is a directory, skipping", f)
		}
	}

	c.SourceFiles = files
	return nil
}
