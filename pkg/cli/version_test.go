package cli

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// TestPrintVersionInfo_SmokeTest is a basic smoke test for PrintVersionInfo
func TestPrintVersionInfo_SmokeTest(t *testing.T) {
	// Create a pipe to capture stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}

	// Redirect stdout to our pipe
	os.Stdout = w
	origStdout := os.Stdout
	defer func() { os.Stdout = origStdout }()

	// Channel to receive the captured output
	outputChan := make(chan string, 1)

	// Start a goroutine to read from the pipe
	go func() {
		defer r.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outputChan <- buf.String()
	}()

	// Call the function we're testing
	PrintVersionInfo()

	// Close the writer to signal EOF
	w.Close()

	// Get the captured output
	output := <-outputChan

	// Basic smoke test assertions
	if output == "" {
		t.Error("PrintVersionInfo() produced no output")
	}

}

// TestPrintVersionInfo_DoesNotPanic tests that the function doesn't panic
func TestPrintVersionInfo_DoesNotPanic(t *testing.T) {
	// Test that the function doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PrintVersionInfo() panicked: %v", r)
		}
	}()

	PrintVersionInfo()
}

// TestPrintVersionInfo_OutputFormat tests basic output formatting
func TestPrintVersionInfo_OutputFormat(t *testing.T) {
	// Create a pipe to capture stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}

	// Redirect stdout to our pipe
	origStdout := os.Stdout
	defer func() { os.Stdout = origStdout }()
	os.Stdout = w

	// Channel to receive the captured output
	outputChan := make(chan string, 1)

	// Start a goroutine to read from the pipe
	go func() {
		defer r.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outputChan <- buf.String()
	}()

	// Call the function
	PrintVersionInfo()

	// Close the writer
	w.Close()

	// Get the output
	output := <-outputChan

	// Test that output is not just whitespace
	if strings.TrimSpace(output) == "" {
		t.Error("PrintVersionInfo() produced only whitespace")
	}

	// Test that output doesn't contain obvious error patterns
	errorPatterns := []string{"error", "panic", "fatal", "nil pointer"}
	for _, pattern := range errorPatterns {
		if strings.Contains(strings.ToLower(output), pattern) {
			t.Errorf("PrintVersionInfo() output contains error pattern %q: %s", pattern, output)
		}
	}

	// Test basic formatting - should be printable text
	for _, char := range output {
		if char < 32 && char != '\n' && char != '\t' && char != '\r' {
			t.Errorf("PrintVersionInfo() output contains non-printable character: %d", char)
			break
		}
	}
}
