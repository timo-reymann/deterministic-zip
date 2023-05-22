package output

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
	"testing"
)

func TestVerboseIsEnabled(t *testing.T) {
	verboseConfig := cli.Configuration{Verbose: true}
	verbose := Verbose{}
	if !verbose.IsEnabled(&verboseConfig) {
		t.Fatal("Execution for debug flag not working")
	}
}

func TestVerboseDebugName(t *testing.T) {
	testDebugName(t, (Verbose{}).DebugName(), "Verbose")
}

func TestVerboseExecute(t *testing.T) {
	verbose := Verbose{}
	if err := verbose.Execute(nil); err != nil {
		t.Fatal(err)
	}

	if output.Level() != output.LevelDebug {
		t.Fatal("Log level not set")
	}
}
