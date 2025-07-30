package fileset

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"reflect"
	"strconv"
	"testing"
)

func TestIgnoreNonExistent_IsEnabled(t *testing.T) {
	ignoreNonExistent := IgnoreNonExistent{}
	if !ignoreNonExistent.IsEnabled(&cli.Configuration{}) {
		t.Fatal("The feature should be enabled by default")
	}
}

func TestIgnoreNonExistent_DebugName(t *testing.T) {
	testDebugName(t, (IgnoreNonExistent{}).DebugName(), "IgnoreNonExistent (builtin)")
}

func TestIgnoreNonExistent_Execute(t *testing.T) {
	ignoreNonExistent := IgnoreNonExistent{}

	testCases := []struct {
		sources      []string
		err          *error
		sourcesAfter []string
	}{
		{
			sources: []string{
				"testdata",
				"nonExistent",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata",
			},
		},
		{
			sources: []string{
				"testdata",
				"testdata/file.txt",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata",
				"testdata/file.txt",
			},
		},
		{
			sources: []string{
				"testdata",
				"testdata/foo/file.txt",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata",
			},
		},
	}

	for idx, tc := range testCases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			c := cli.Configuration{
				SourceFiles: tc.sources,
			}
			err := ignoreNonExistent.Execute(&c)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.sourcesAfter, c.SourceFiles) {
				t.Fatalf("Expected %v, but got %v", tc.sourcesAfter, c.SourceFiles)
			}
		})

	}
}
