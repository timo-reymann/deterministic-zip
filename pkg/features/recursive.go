package features

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/file"
	"github.com/timo-reymann/deterministic-zip/pkg/log"
)

type Recursive struct{}

func (r Recursive) IsEnabled(c *cli.Configuration) bool {
	return OnFlag(c.Recursive)
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
			log.Debugf("Add path %s", p)
			files = append(files, p)
		}
	}

	c.SourceFiles = files

	return nil
}
