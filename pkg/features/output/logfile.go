package output

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
	"os"
)

// LogFile enables writing output to a log file
type LogFile struct{}

// DebugName returns the debuggable name
func (l LogFile) DebugName() string {
	return "LogFile"
}

// IsEnabled check if the logfile flag is enabled
func (l LogFile) IsEnabled(c *cli.Configuration) bool {
	return c.LogFilePath != ""
}

// Execute sets the output log file
func (l LogFile) Execute(c *cli.Configuration) error {
	if !c.LogFileAppend {
		_ = os.Remove(c.LogFilePath)
	}

	output.SetOutput(c.LogFilePath)
	return nil
}
