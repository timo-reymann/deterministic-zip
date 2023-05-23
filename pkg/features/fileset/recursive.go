package fileset

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features/conditions"
	"github.com/timo-reymann/deterministic-zip/pkg/file"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
)

// Recursive adds all children folders recursively
type Recursive struct{}

// DebugName prints the debuggable name
func (r Recursive) DebugName() string {
	return "Recursive"
}

// IsEnabled checks if recursive adding was requested
func (r Recursive) IsEnabled(c *cli.Configuration) bool {
	return conditions.OnFlag(c.Recursive)
}

// Execute and read all directories recursively and add them back to source files
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

		output.Debugf("read directory %s", f)
		paths, err := file.ReadDirRecursive(f)
		if err != nil {
			return err
		}

		for _, p := range paths {
			files = append(files, p)
		}
	}

	c.SourceFiles = files

	return nil
}
