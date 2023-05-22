package fileset

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"reflect"
	"testing"
)

func TestIgnoreTargetZip_IsEnabled(t *testing.T) {
	ignoreZip := IgnoreTargetZip{}
	if !ignoreZip.IsEnabled(&cli.Configuration{}) {
		t.Fatal("The feature should be enabled by default")
	}
}

func TestIgnoreTargetZip_DebugName(t *testing.T) {
	testDebugName(t, (IgnoreTargetZip{}).DebugName(), "IgnoreTargetZip (builtin)")
}

func TestIgnoreTargetZip_Execute(t *testing.T) {
	ignoreZip := IgnoreTargetZip{}

	testCases := []struct {
		sources      []string
		zipFile      string
		err          *error
		sourcesAfter []string
	}{
		{
			zipFile: "test.zip",
			sources: []string{
				"testdata",
				"test.zip",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata",
			},
		},
		{
			zipFile: "test.zip",
			sources: []string{
				"testdata",
				"test.zip",
				"archive.zip",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata",
				"archive.zip",
			},
		},
		{
			zipFile: "test.zip",
			sources: []string{
				"testdata",
				"test.zip",
				"folder/test.zip",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata",
				"folder/test.zip",
			},
		},
	}

	for _, tc := range testCases {
		c := cli.Configuration{
			ZipFile:     tc.zipFile,
			SourceFiles: tc.sources,
		}
		err := ignoreZip.Execute(&c)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(tc.sourcesAfter, c.SourceFiles) {
			t.Fatalf("Expected %v, but got %v", tc.sourcesAfter, c.SourceFiles)
		}
	}
}
