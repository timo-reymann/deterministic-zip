package file

import (
	"reflect"
	"testing"
)

func TestFindByGlob(t *testing.T) {
	testCases := []struct {
		pattern string
		result  []string
	}{
		{
			pattern: "testdata/glob/*.json",
			result: []string{
				"testdata/glob/a.json",
				"testdata/glob/b.json",
				"testdata/glob/c.json",
			},
		},
		{
			pattern: "testdata/glob/*.csv",
			result:  []string{},
		},
		{
			pattern: "testdata/glob/a*",
			result: []string{
				"testdata/glob/a.json",
				"testdata/glob/a.txt",
			},
		},
		{
			pattern: "testdata/glob/**/a*",
			result: []string{
				"testdata/glob/a.json",
				"testdata/glob/a.txt",
			},
		},
		{
			pattern: "testdata/glob/**",
			result: []string{
				"testdata/glob/",
				"testdata/glob/a.json",
				"testdata/glob/a.txt",
				"testdata/glob/b.json",
				"testdata/glob/b.txt",
				"testdata/glob/c.json",
				"testdata/glob/c.txt",
			},
		},
	}

	for _, tc := range testCases {
		results := FindByGlob(tc.pattern)

		// DeepEquals doesnt like empty arrays
		if len(tc.result) == 0 && len(results) == 0 {
			continue
		}

		if !reflect.DeepEqual(results, tc.result) {
			t.Fatalf("Expected %v, but got %v for pattern %s", tc.result, results, tc.pattern)
		}
	}
}
