package output

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"github.com/timo-reymann/deterministic-zip/pkg/output"
	"os"
	"testing"
)

func TestLogFileIsEnabled(t *testing.T) {
	lf := LogFile{}
	lf.IsEnabled(&cli.Configuration{LogFilePath: "/tmp/log"})
}

func TestLogFileExecute(t *testing.T) {
	lf := LogFile{}
	err := lf.Execute(&cli.Configuration{LogFilePath: "/tmp/log"})
	if err != nil {
		t.Fatal(err)
	}

	output.Info("test")

	_, err = os.Stat("/tmp/log")
	if err != nil {
		t.Fatal(err)
	}
}
