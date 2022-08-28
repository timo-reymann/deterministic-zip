package fileset

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features/conditions"
	"github.com/timo-reymann/deterministic-zip/pkg/file"
	"os"
	"path/filepath"
	"sort"
)

// Directories adds all children folders recursively
type Directories struct{}

// DebugName prints the debuggable name
func (r Directories) DebugName() string {
	return "Directories"
}

// IsEnabled checks if recursive adding was requested
func (r Directories) IsEnabled(c *cli.Configuration) bool {
	return conditions.OnFlag(c.Directories)
}

// Execute and read all directories recursively and add them back to source files
func (r Directories) Execute(c *cli.Configuration) error {
	files := make(map[string]string, 0)
	sort.Strings(c.SourceFiles)

	for _, f := range c.SourceFiles {
		isDir, err := file.IsDir(f)
		if err != nil {
			return err
		}

		cf := filepath.Clean(f)
		if cf != "." {
			if isDir {
				cf += string(os.PathSeparator)
			}
			files[cf] = f
		}
		includeParentDirs(&files, f)
	}

	c.SourceFiles = sortedKeySlice(&files)

	return nil
}

func includeParentDirs(m *map[string]string, path string) {
	parent := filepath.Dir(path)

	for parent != "." {
		if !conditions.ContainsKey(m, parent) {
			suffixed := parent + string(os.PathSeparator)
			(*m)[suffixed] = parent
		}
		parent = filepath.Dir(parent)
	}
}

func sortedKeySlice(m *map[string]string) []string {
	keys := make([]string, 0, len(*m))
	for k := range *m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
