package cli

import (
	"testing"
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
