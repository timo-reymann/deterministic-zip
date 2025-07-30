package fileset

import (
	"errors"
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"reflect"
	"strconv"
	"testing"
)

func TestRecursive_IsEnabled(t *testing.T) {
	c := cli.Configuration{Recursive: true}
	recursive := Recursive{}
	if !recursive.IsEnabled(&c) {
		t.Fatal("Execution for recursive flag not working")
	}
}

func TestRecursive_DebugName(t *testing.T) {
	testDebugName(t, (Recursive{}).DebugName(), "Recursive")
}

func mockErr(msg string) *error {
	err := errors.New(msg)
	return &err
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
				"testdata/recursive",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata/recursive",
				"testdata/recursive/folder",
				"testdata/recursive/folder/file.txt",
				"testdata/recursive/folder/subfolder",
				"testdata/recursive/folder/subfolder/test",
			},
		},
		{
			sources: []string{
				"testdata/recursive/folder/file.txt",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata/recursive/folder/file.txt",
			},
		},
		{
			sources: []string{
				"testdata/recursive/folder/subfolder/",
			},
			err: nil,
			sourcesAfter: []string{
				"testdata/recursive/folder/subfolder/",
				"testdata/recursive/folder/subfolder/test",
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

	for idx, tc := range testCases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			c := cli.Configuration{
				SourceFiles: tc.sources,
			}
			err := recursive.Execute(&c)
			if tc.err == nil && err != nil {
				t.Fatalf("Expected no error but got %v", err)
			} else if err != nil {
				if (*tc.err).Error() != err.Error() {
					t.Fatalf("Expected error %v, but got %v", *tc.err, err)
				} else {
					// Skip checking -> error thrown
					return
				}
			}

			if !reflect.DeepEqual(tc.sourcesAfter, c.SourceFiles) {
				t.Fatalf("Expected %v, but got %v", tc.sourcesAfter, c.SourceFiles)
			}
		})

	}

}
