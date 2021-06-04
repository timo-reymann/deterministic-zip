package features

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features/conditions"
	"os"
)

type NoDirectories struct {
}

func (n NoDirectories) DebugName() string {
	return "NoDirectories"
}

func (n NoDirectories) IsEnabled(c *cli.Configuration) bool {
	return conditions.OnFlag(c.NoDirEntries)
}

func (n NoDirectories) Execute(c *cli.Configuration) error {
	files := make([]string, 0)

	for _, f := range c.SourceFiles {
		stat, err := os.Stat(f)
		if err != nil {
			return err
		}

		if !stat.IsDir() {
			files = append(files, f)
		}
	}

	c.SourceFiles = files
	return nil
}
