package output

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features/conditions"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
)

type Quiet struct {
}

// IsEnabled checks if quiet mode is enabled
func (q Quiet) IsEnabled(c *cli.Configuration) bool {
	return conditions.OnFlag(c.Quiet)
}

// Execute and set the log level to silence
func (q Quiet) Execute(c *cli.Configuration) error {
	output.SetLevel(output.LevelSilence)
	return nil
}

// DebugName prints the debuggable name
func (q Quiet) DebugName() string {
	return "Quiet"
}
