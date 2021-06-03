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
			pattern: "*.json",
			result: []string{
				"testdata/glob/a.json",
				"testdata/glob/b.json",
				"testdata/glob/c.json",
			},
		},
		{
			pattern: "*.csv",
			result:  []string{},
		},
		{
			pattern: "a*",
			result: []string{
				"testdata/glob/a.json",
				"testdata/glob/a.txt",
			},
		},
	}

	for _, tc := range testCases {
		results := FindByGlob("testdata/glob/" + tc.pattern)

		// DeepEquals doesnt like empty arrays
		if len(tc.result) == 0 && len(results) == 0 {
			continue
		}

		if !reflect.DeepEqual(results, tc.result) {
			t.Fatalf("Expected %d results, but got %d", len(tc.result), len(results))
		}
	}
}
