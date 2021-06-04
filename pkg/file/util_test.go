package file

import (
	"reflect"
	"testing"
)

func TestIsDir(t *testing.T) {
	testCases := []struct {
		path   string
		result bool
	}{
		{
			"nonexistent",
			false,
		},
		{
			"testdata/util/one",
			true,
		},
	}

	for _, tc := range testCases {
		if isDir, _ := IsDir(tc.path); isDir != tc.result {
			t.Fatalf("Expected result to be %v but got %v", tc.result, isDir)
		}
	}
}

func TestReadDirRecursive(t *testing.T) {
	testCases := []struct {
		path    string
		results []string
	}{
		{
			path: "testdata/util",
			results: []string{
				"testdata/util/",
				"testdata/util/one/",
				"testdata/util/one/two/",
				"testdata/util/one/two/three/",
				"testdata/util/one/two/three/file",
			},
		},
	}

	for _, tc := range testCases {
		results, _ := ReadDirRecursive(tc.path)
		if !reflect.DeepEqual(results, tc.results) {
			t.Fatalf("Expected %v, but got %v", tc.results, results)
		}
	}
}
