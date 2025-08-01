package cli

import (
	"os"
	"testing"
	"time"
)

func TestNewConfiguration(t *testing.T) {
	c := NewConfiguration()
	if c == nil {
		t.Errorf("Expected a new configuration, got nil")
	}
}

func TestDefineFlags(t *testing.T) {
	c := NewConfiguration()
	c.defineFlags()
}

func TestCleanPath(t *testing.T) {
	tests := []struct {
		name               string
		input              string
		isRetardedPlatform bool
		expected           string
	}{
		{
			name:               "Unix path",
			input:              "/home/user/../user/docs",
			isRetardedPlatform: false,
			expected:           "/home/user/docs",
		},
		{
			name:               "Windows path",
			input:              "C:\\Users\\user\\docs",
			isRetardedPlatform: true,
			expected:           "C:/Users/user/docs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cleanPath(tt.input, tt.isRetardedPlatform)
			if result != tt.expected {
				t.Errorf("cleanPath(%q, %v) = %q; want %q", tt.input, tt.isRetardedPlatform, result, tt.expected)
			}
		})
	}
}

func TestCleanPaths(t *testing.T) {
	tests := []struct {
		name        string
		sourceFiles []string
		expected    []string
	}{
		{
			name:        "Unix paths",
			sourceFiles: []string{"/home/user/../user/docs", "/var/log/../log/syslog"},
			expected:    []string{"/home/user/docs", "/var/log/syslog"},
		},
		{
			name:        "Windows paths",
			sourceFiles: []string{"C:\\Users\\user\\docs", "D:\\data\\files"},
			expected:    []string{"C:\\Users\\user\\docs", "D:\\data\\files"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Configuration{SourceFiles: tt.sourceFiles}
			conf.CleanPaths()
			for i, got := range conf.SourceFiles {
				if got != tt.expected[i] {
					t.Errorf("CleanPaths() = %q; want %q", got, tt.expected[i])
				}
			}
		})
	}
}

func TestConfiguration_Help_SmokeTest(t *testing.T) {
	// Create a Configuration instance
	config := NewConfiguration()

	// Call Help method - should not fail
	config.Help()

	// If we reach here, the method completed without panicking
	t.Log("Configuration.Help() completed successfully")
}

// TestConfiguration_Parse_HelpFlag tests Parse with help flag
func TestConfiguration_Parse_HelpFlag(t *testing.T) {
	config := NewConfiguration()

	testCases := []struct {
		name string
		args []string
	}{
		{
			name: "short help flag",
			args: []string{"program", "-h"},
		},
		{
			name: "long help flag",
			args: []string{"program", "--help"},
		},
		{
			name: "help flag with other args",
			args: []string{"program", "-h", "extra", "args"},
		},
		{
			name: "help flag at end",
			args: []string{"program", "command", "--help"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Save original os.Args and restore after test
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			// Set test arguments
			os.Args = tt.args

			// Call Parse - should handle help flag gracefully
			config.Parse()

			t.Logf("Parse() with args %v completed successfully", tt.args)
		})
	}
}

// TestConfiguration_Parse_VersionFlag tests Parse with version flag
func TestConfiguration_Parse_VersionFlag(t *testing.T) {
	config := NewConfiguration()

	testCases := []struct {
		name string
		args []string
	}{
		{
			name: "long version flag",
			args: []string{"program", "--version"},
		},
		{
			name: "version flag with other args",
			args: []string{"program", "--version", "ignored"},
		},
		{
			name: "version flag at different position",
			args: []string{"program", "command", "-v"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			// Save original os.Args and restore after test
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			// Set test arguments
			os.Args = tt.args

			// Call Parse - should handle version flag gracefully
			config.Parse()

			t.Logf("Parse() with args %v completed successfully", tt.args)
		})
	}
}

// TestConfiguration_Parse_HelpAndVersionFlags tests Parse with both help and version flags
func TestConfiguration_Parse_HelpAndVersionFlags(t *testing.T) {
	config := NewConfiguration()

	testCases := []struct {
		name string
		args []string
	}{
		{
			name: "help and version together",
			args: []string{"program", "-h", "-v"},
		},
		{
			name: "version and help together",
			args: []string{"program", "--version", "--help"},
		},
		{
			name: "mixed short and long flags",
			args: []string{"program", "-h", "--version"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Test that Parse doesn't panic with multiple flags
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Configuration.Parse() with help and version flags panicked: %v", r)
				}
			}()

			// Save original os.Args and restore after test
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			// Set test arguments
			os.Args = tt.args

			// Call Parse
			config.Parse()

			t.Logf("Parse() with args %v completed successfully", tt.args)
		})
	}
}

// TestConfiguration_Parse_EmptyArgs tests Parse with no arguments
func TestConfiguration_Parse_EmptyArgs(t *testing.T) {
	config := NewConfiguration()

	// Save original os.Args and restore after test
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Set minimal arguments (just program name)
	os.Args = []string{"program"}

	// Call Parse
	config.Parse()

	t.Log("Parse() with empty args completed successfully")
}

// TestConfiguration_Parse_InvalidHelpVariations tests Parse with invalid help flag variations
func TestConfiguration_Parse_InvalidHelpVariations(t *testing.T) {
	config := NewConfiguration()

	testCases := []struct {
		name string
		args []string
	}{
		{
			name: "malformed help flag",
			args: []string{"program", "-help"},
		},
		{
			name: "help as argument",
			args: []string{"program", "help"},
		},
		{
			name: "help with equals",
			args: []string{"program", "--help=true"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Test that Parse handles malformed help flags gracefully
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Configuration.Parse() with malformed help flag panicked: %v", r)
				}
			}()

			// Save original os.Args and restore after test
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			// Set test arguments
			os.Args = tt.args

			// Call Parse
			config.Parse()

			t.Logf("Parse() with args %v completed successfully", tt.args)
		})
	}
}

// TestConfiguration_Parse_MultipleInstances tests Parse on multiple Configuration instances
func TestConfiguration_Parse_MultipleInstances(t *testing.T) {
	testCases := []struct {
		name string
		args []string
	}{
		{
			name: "help flag",
			args: []string{"program", "--help"},
		},
		{
			name: "version flag",
			args: []string{"program", "--version"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Test multiple Configuration instances
			config1 := NewConfiguration()
			config2 := NewConfiguration()

			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Multiple Configuration.Parse() calls panicked: %v", r)
				}
			}()

			// Save original os.Args and restore after test
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			// Set test arguments
			os.Args = tt.args

			// Call Parse on both instances
			config1.Parse()
			config2.Parse()

			t.Logf("Multiple Parse() calls with args %v completed successfully", tt.args)
		})
	}
}
func TestParseModifiedDate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
		wantErr  bool
	}{
		{
			name:     "valid unix timestamp",
			input:    "1703520645",
			expected: time.Unix(1703520645, 0).UTC(),
			wantErr:  false,
		},
		{
			name:     "zero timestamp",
			input:    "0",
			expected: time.Unix(0, 0).UTC(),
			wantErr:  false,
		},
		{
			name:     "negative timestamp",
			input:    "-86400",
			expected: time.Unix(-86400, 0).UTC(),
			wantErr:  false,
		},
		{
			name:     "future timestamp",
			input:    "2147483647",
			expected: time.Unix(2147483647, 0).UTC(),
			wantErr:  false,
		},
		{
			name:    "invalid non-numeric input",
			input:   "not-a-number",
			wantErr: true,
		},
		{
			name:     "empty string",
			input:    "",
			wantErr:  false,
			expected: DefaultModifiedTimestamp,
		},
		{
			name:    "floating point number",
			input:   "1703520645.5",
			wantErr: true,
		},
		{
			name:    "timestamp with letters",
			input:   "1703520645abc",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Configuration{}
			err := os.Setenv("SOURCE_DATE_EPOCH", tt.input)
			if err != nil {
				t.Fatal(err)
			}
			result, err := conf.parseModifiedDate()

			if tt.wantErr {
				if err == nil {
					t.Errorf("parseModifiedDate() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("parseModifiedDate() unexpected error: %v", err)
				return
			}

			if !result.Equal(tt.expected) {
				t.Errorf("parseModifiedDate() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestConfiguration_ModifiedDate(t *testing.T) {
	tests := []struct {
		name          string
		configuration *Configuration
		expectedTime  time.Time
		expectedErr   bool
	}{
		{
			name:          "modified date not set",
			configuration: &Configuration{
				// Assuming the modified date field is empty/zero
			},
			expectedTime: DefaultModifiedTimestamp,
			expectedErr:  true,
		},
		{
			name: "modified date set",
			configuration: &Configuration{
				modifiedDate: time.Unix(1703520645, 0).UTC(),
			},
			expectedTime: time.Unix(1703520645, 0).UTC(),
			expectedErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.configuration.ModifiedDate()

			if !result.Equal(tt.expectedTime) {
				t.Errorf("Configuration.ModifiedDate() = %v, expected %v", result, tt.expectedTime)
			}
		})
	}
}
