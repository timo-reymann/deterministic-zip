package filter

import (
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"reflect"
	"testing"
)

func TestInclude_IsEnabled(t *testing.T) {
	config := cli.Configuration{Include: []string{
		"foo.*",
	}}
	include := Include{}
	if !include.IsEnabled(&config) {
		t.Fatalf("Should include for non empty include")
	}
}

func TestInclude_DebugName(t *testing.T) {
	testDebugName(t, (Include{}).DebugName(), "Include")
}

func TestInclude_Execute(t *testing.T) {
	testCases := []struct {
		sourceFiles []string
		targetFiles []string
		patterns    []string
	}{
		{
			sourceFiles: []string{
				"export.json",
				"export.csv",
			},
			targetFiles: []string{
				"export.json",
			},
			patterns: []string{
				"*.json",
			},
		},
		{
			sourceFiles: []string{
				"export.json",
				"export.csv",
				"private/export.json",
			},
			targetFiles: []string{
				"private/export.json",
			},
			patterns: []string{
				"*.json",
				"private*",
			},
		},
		{
			sourceFiles: []string{
				"export.json",
				"export.csv",
				"private/export.json",
			},
			targetFiles: []string{},
			patterns: []string{
				"*.prop",
				"private*",
			},
		},
	}

	for _, tc := range testCases {
		config := cli.Configuration{
			SourceFiles: tc.sourceFiles,
			Include:     tc.patterns,
		}

		include := Include{}
		if err := include.Execute(&config); err != nil {
			t.Fatal(err)
		}

		// DeepEquals doesnt like empty arrays
		if len(tc.targetFiles) == 0 && len(config.SourceFiles) == 0 {
			continue
		}

		if !reflect.DeepEqual(tc.targetFiles, config.SourceFiles) {
			t.Fatalf("Expected %v, but got %v for patterns %v", tc.targetFiles, config.SourceFiles, tc.patterns)
		}
	}
}
