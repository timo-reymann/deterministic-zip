package features

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"reflect"
	"testing"
)

func TestRecursive_IsEnabled(t *testing.T) {
	c := cli.Configuration{Recursive: true}
	recursive := Recursive{}
	if !recursive.IsEnabled(&c) {
		t.Fatal("Execution for recursive flag not working")
	}
}

func TestRecursive_Execute(t *testing.T) {
	recursive := Recursive{}
	testCases := []struct {
		sources      []string
		err          *error
		sourcesAfter []string
	}{
		{
			sources: []string{
				"testdata",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata/",
				"testdata/folder/",
				"testdata/folder/file.txt",
				"testdata/folder/subfolder/",
				"testdata/folder/subfolder/test",
			},
		},
		{
			sources: []string{
				"testdata/folder/file.txt",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata/folder/file.txt",
			},
		},
		{
			sources: []string{
				"testdata/folder/subfolder",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata/folder/subfolder/",
				"testdata/folder/subfolder/test",
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
			SourceFiles: tc.sources,
		}
		err := recursive.Execute(&c)
		if tc.err == nil && err != nil {
			t.Fatalf("Expected no error but got %v", err)
			t.FailNow()
		} else if err != nil {
			if (*tc.err).Error() != err.Error() {
				t.Fatalf("Expected error %v, but got %v", *tc.err, err)
				t.FailNow()
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
