package file

import (
	"testing"
)

func TestFindInString(t *testing.T) {
	testCases := []struct {
		path    string
		pattern string
		result  bool
	}{
		{
			path:    "testdata/glob/a.json",
			pattern: "testdata/glob/*.json",
			result:  true,
		},
		{
			path:    "testdata/glob/a.json",
			pattern: "testdata/glob/**",
			result:  true,
		},
		{
			path:    "testdata/glob",
			pattern: "testdata/glob*",
			result:  true,
		},
		{
			path:    "node_modules",
			pattern: "*node_modules*",
			result:  true,
		},
		{
			path:    "node_modules/some_package/index.js",
			pattern: "*node_modules*",
			result:  true,
		},
		{
			path:    "node_modules/some_package/index.js",
			pattern: "*node_modules/*",
			result:  true,
		},
		{
			path:    "node_modules/some_package/index.js",
			pattern: "node_modules/*",
			result:  true,
		},
		{
			path:    ".git",
			pattern: ".git*",
			result:  true,
		},
		{
			path:    ".git/HEAD",
			pattern: ".git/*",
			result:  true,
		},
		{
			path:    ".git",
			pattern: ".git*",
			result:  true,
		},
		{
			path:    ".git/hooks/applypatch-msg.sample",
			pattern: ".git/*",
			result:  true,
		},
		{
			path:    ".git/hooks/",
			pattern: ".git/*",
			result:  true,
		},
		{
			path:    ".git/asdf",
			pattern: ".git/*",
			result:  true,
		},
		{
			path:    "test.c",
			pattern: "*.[!o]",
			result:  true,
		},
		{
			path:    "test.c",
			pattern: "*.[hc]",
			result:  true,
		},
		{
			path:    "test.h",
			pattern: "*.[hc]",
			result:  true,
		},
		{
			path:    "path/to/test.h",
			pattern: "*.[hc]",
			result:  true,
		},
		{
			path:    "path/to/test.c",
			pattern: "*.[hc]",
			result:  true,
		},
		{
			path:    "z.file",
			pattern: "[a-f]*",
			result:  false,
		},
		{
			path:    "a.file",
			pattern: "[a-f]*",
			result:  true,
		},
		{
			path:    ".idea",
			pattern: ".idea",
			result:  true,
		},
	}

	for _, tc := range testCases {
		result := NewGlob(tc.pattern).Match(tc.path)

		if tc.result != result {
			t.Fatalf("Expected %v, but got %v for pattern %s and input %s", tc.result, result, tc.pattern, tc.path)
		}
	}
}
