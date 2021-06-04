package features

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/log"
	"testing"
)

func TestVerbose_IsEnabled(t *testing.T) {
	verboseConfig := cli.Configuration{Verbose: true}
	verbose := Verbose{}
	if !verbose.IsEnabled(&verboseConfig) {
		t.Fatal("Execution for debug flag not working")
	}
}

func TestVerbose_Execute(t *testing.T) {
	verbose := Verbose{}
	if err := verbose.Execute(nil); err != nil {
		t.Fatal(err)
	}

	if log.Level() != log.LevelDebug {
		t.Fatal("Log level not set")
	}
}
