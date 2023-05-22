package fileset

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"path/filepath"
)

// IgnoreTargetZip removes the zip file from the file set if found
type IgnoreTargetZip struct{}

func (i IgnoreTargetZip) DebugName() string {
	return "IgnoreTargetZip (builtin)"
}

// IsEnabled always returns true as this feature is a safety mechanism
func (i IgnoreTargetZip) IsEnabled(c *cli.Configuration) bool {
	return true
}

// Execute removes the target zip file if it finds it in the fileset list and exits
func (i IgnoreTargetZip) Execute(c *cli.Configuration) error {
	target := filepath.Clean(c.ZipFile)

	for idx, f := range c.SourceFiles {
		if target == f {
			c.SourceFiles = append(c.SourceFiles[:idx], c.SourceFiles[idx+1:]...)
			break
		}
	}
	return nil
}
