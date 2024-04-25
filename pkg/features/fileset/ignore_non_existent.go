package fileset

import (
	"fmt"
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
	"os"
)

// IgnoreNonExistent filters out files that don't exist and prints a warning if not run in quiet mode
type IgnoreNonExistent struct{}

// DebugName prints the debuggable name
func (i IgnoreNonExistent) DebugName() string {
	return "IgnoreNonExistent (builtin)"
}

// IsEnabled always is the case as this is a builtin functionality to achieve parity with zip
func (i IgnoreNonExistent) IsEnabled(c *cli.Configuration) bool {
	return true
}

// Execute and filter out all files and folders that cannot be accessed
func (i IgnoreNonExistent) Execute(c *cli.Configuration) error {
	files := make([]string, 0)

	for _, entry := range c.SourceFiles {
		if _, err := os.Stat(entry); err != nil {
			if !c.Quiet {
				output.Info(fmt.Sprintf("Skipping %s from adding to archive, no such entry or directory.", entry))
			}
			continue
		}

		files = append(files, entry)
	}

	c.SourceFiles = files

	return nil
}
