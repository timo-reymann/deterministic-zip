package features

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/features/conditions"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
)

type Verbose struct {
}

func (v Verbose) DebugName() string {
	return "Verbose"
}

func (v Verbose) IsEnabled(c *cli.Configuration) bool {
	return conditions.OnFlag(c.Verbose)
}

func (v Verbose) Execute(c *cli.Configuration) error {
	output.SetLevel(output.LevelDebug)
	output.Debug("Enable verbose mode")
	return nil
}
