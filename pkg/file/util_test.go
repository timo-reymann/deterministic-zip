package file

import (
	"reflect"
	"strconv"
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

	for idx, tc := range testCases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			if isDir, _ := IsDir(tc.path); isDir != tc.result {
				t.Fatalf("Expected result to be %v but got %v", tc.result, isDir)
			}
		})
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
				"testdata/util",
				"testdata/util/one",
				"testdata/util/one/two",
				"testdata/util/one/two/three",
				"testdata/util/one/two/three/file",
			},
		},
	}

	for idx, tc := range testCases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			results, _ := ReadDirRecursive(tc.path)
			if !reflect.DeepEqual(results, tc.results) {
				t.Fatalf("Expected %v, but got %v", tc.results, results)
			}
		})
	}
}
