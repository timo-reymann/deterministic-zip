package features

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features/conditions"
	"github.com/timo-reymann/deterministic-zip/pkg/file"
	"sort"
)

type Recursive struct{}

func (r Recursive) DebugName() string {
	return "Recursive"
}

func (r Recursive) IsEnabled(c *cli.Configuration) bool {
	return conditions.OnFlag(c.Recursive)
}

func (r Recursive) Execute(c *cli.Configuration) error {
	files := make([]string, 0)

	for _, f := range c.SourceFiles {
		isDir, err := file.IsDir(f)
		if err != nil {
			return err
		}

		if !isDir {
			files = append(files, f)
			continue
		}

		paths, err := file.ReadDirRecursive(f)
		if err != nil {
			return err
		}

		for _, p := range paths {
			files = append(files, p)
		}
	}

	sort.Strings(files)

	c.SourceFiles = files

	return nil
}
