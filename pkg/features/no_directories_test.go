package features

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"testing"
)

func TestNoDirectories_IsEnabled(t *testing.T) {
	config := cli.Configuration{NoDirEntries: true}
	noDir := NoDirectories{}
	if !noDir.IsEnabled(&config) {
		t.Fatalf("Should execute for noDirDirectories")
	}
}

func TestNoDirectories_Execute(t *testing.T) {
	noDir := NoDirectories{}
	config := cli.Configuration{
		SourceFiles: []string{
			"testdata/no-dir/file.txt",
			"testdata/no-dir/subfolder/",
		},
		NoDirEntries: true,
	}

	err := noDir.Execute(&config)
	if err != nil {
		t.Fatal(err)
	}

	if len(config.SourceFiles) != 1 {
		t.Fatalf("Expected one file, got %v", config.SourceFiles)
	}
}
