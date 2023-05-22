package fileset

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"reflect"
	"testing"
)

func TestDirectories_IsEnabled(t *testing.T) {
	c := cli.Configuration{Directories: true}
	directories := Directories{}
	if !directories.IsEnabled(&c) {
		t.Fatal("Execution for directories fallback not working")
	}
}

func TestDirectories_DebugName(t *testing.T) {
	testDebugName(t, (Directories{}).DebugName(), "Directories")
}

func TestDirectories_Execute(t *testing.T) {
	directories := Directories{}
	testCases := []struct {
		sources      []string
		err          *error
		sourcesAfter []string
	}{
		{
			sources: []string{
				"testdata/recursive",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata/",
				"testdata/recursive/",
			},
		},
		{
			sources: []string{
				"testdata/recursive/folder/file.txt",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata/",
				"testdata/recursive/",
				"testdata/recursive/folder/",
				"testdata/recursive/folder/file.txt",
			},
		},
		{
			sources: []string{
				"testdata/recursive/folder/subfolder/",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata/",
				"testdata/recursive/",
				"testdata/recursive/folder/",
				"testdata/recursive/folder/subfolder/",
			},
		},
		{
			sources: []string{
				"nonExistent",
			},
			err:          mockErr("stat nonExistent: no such file or directory"),
			sourcesAfter: []string{},
		},
	}

	for _, tc := range testCases {
		c := cli.Configuration{
			Directories: true,
			SourceFiles: tc.sources,
		}
		err := directories.Execute(&c)
		if tc.err == nil && err != nil {
			t.Fatalf("Expected no error but got %v", err)
		} else if err != nil {
			if (*tc.err).Error() != err.Error() {
				t.Fatalf("Expected error %v, but got %v", *tc.err, err)
			} else {
				// Skip checking -> error thrown
				continue
			}
		}

		if !reflect.DeepEqual(tc.sourcesAfter, c.SourceFiles) {
			t.Fatalf("Expected %v, but got %v", tc.sourcesAfter, c.SourceFiles)
		}
	}

}
