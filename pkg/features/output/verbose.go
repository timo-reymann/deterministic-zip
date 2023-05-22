package output

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features/conditions"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
)

// Verbose enable verbose outputs
type Verbose struct {
}

// DebugName prints the debuggable name
func (v Verbose) DebugName() string {
	return "Verbose"
}

// IsEnabled checks if vebose mode is active
func (v Verbose) IsEnabled(c *cli.Configuration) bool {
	return conditions.OnFlag(c.Verbose)
}

// Execute and set level to Debug
func (v Verbose) Execute(c *cli.Configuration) error {
	output.SetLevel(output.LevelDebug)
	return nil
}
