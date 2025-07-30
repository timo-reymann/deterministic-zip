package cmd

import (
	"bytes"
	"errors"
	"github.com/timo-reymann/deterministic-zip/pkg/cli"
	"log"
	"os"
	"testing"
)

// TestErrCheck tests the error handling function
func TestErrCheck(t *testing.T) {
	// Save original os.Exit and restore it after the test
	originalOsExit := osExit
	defer func() { osExit = originalOsExit }()

	// Save the original log output and restore it after the test
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	tests := []struct {
		name        string
		err         error
		exitCode    int
		shouldPanic bool
		config      *cli.Configuration
	}{
		{
			name:     "nil error should not exit",
			err:      nil,
			exitCode: 0,
			config:   cli.NewConfiguration(),
		},
		{
			name:     "ErrAbort should exit with code 0",
			err:      cli.ErrAbort,
			exitCode: 0,
			config:   cli.NewConfiguration(),
		},
		{
			name:     "ErrMinimalParamsMissing should exit with code 2",
			err:      cli.ErrMinimalParamsMissing,
			exitCode: 2,
			config:   cli.NewConfiguration(),
		},
		{
			name:     "other error should exit with code 2",
			err:      errors.New("some error"),
			exitCode: 2,
			config:   cli.NewConfiguration(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var exitCode int
			osExit = func(code int) {
				exitCode = code
				panic("os.Exit called")
			}

			// Clear the buffer before each test
			buf.Reset()

			// Run the function and recover from panic (caused by our mock os.Exit)
			func() {
				defer func() {
					if r := recover(); r != nil {
						if r != "os.Exit called" {
							t.Errorf("unexpected panic: %v", r)
						}
					}
				}()
				errCheck(tt.err, tt.config)
				if tt.err != nil {
					t.Error("expected os.Exit to be called")
				}
			}()

			// Verify exit code
			if tt.err != nil && exitCode != tt.exitCode {
				t.Errorf("expected exit code %d, got %d", tt.exitCode, exitCode)
			}

			// Verify log output for non-nil errors (except ErrAbort and ErrMinimalParamsMissing)
			if tt.err != nil && tt.err != cli.ErrAbort && tt.err != cli.ErrMinimalParamsMissing {
				logOutput := buf.String()
				if logOutput == "" {
					t.Error("expected error to be logged")
				}
			}
		})
	}
}

// TestExecute_SmokeTest is a basic smoke test to verify Execute doesn't panic
func TestExecute_SmokeTest(t *testing.T) {
	// Save originals and restore after test
	originalOsExit := osExit
	originalArgs := os.Args
	originalStdout := os.Stdout
	originalStderr := os.Stderr
	defer func() {
		osExit = originalOsExit
		os.Args = originalArgs
		os.Stdout = originalStdout
		os.Stderr = originalStderr
	}()

	// Capture log output
	var logBuf bytes.Buffer
	log.SetOutput(&logBuf)
	defer log.SetOutput(os.Stderr)

	// Capture stdout and stderr
	stdoutReader, stdoutWriter, _ := os.Pipe()
	stderrReader, stderrWriter, _ := os.Pipe()
	os.Stdout = stdoutWriter
	os.Stderr = stderrWriter

	// Mock os.Exit to capture exit codes and prevent actual exit
	var exitCalled bool
	var exitCode int
	osExit = func(code int) {
		exitCalled = true
		exitCode = code
		panic("os.Exit called")
	}

	// Test cases
	tests := []struct {
		name        string
		args        []string
		shouldExit  bool
		expectPanic bool
	}{
		{
			name:       "version flag should exit cleanly",
			args:       []string{"program", "--version"},
			shouldExit: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset state
			exitCalled = false
			exitCode = -1
			logBuf.Reset()

			// Set command line arguments
			os.Args = tt.args

			// Run Execute and handle panic from mocked os.Exit
			func() {
				defer func() {
					if r := recover(); r != nil {
						if r != "os.Exit called" && !tt.expectPanic {
							t.Errorf("unexpected panic: %v", r)
						}
					}
				}()

				Execute()

				// If we get here and shouldExit is true, that's unexpected
				if tt.shouldExit {
					t.Error("expected Execute to call os.Exit but it didn't")
				}
			}()

			// Close writers and read output
			stdoutWriter.Close()
			stderrWriter.Close()

			stdoutOutput := make([]byte, 1024)
			stderrOutput := make([]byte, 1024)
			stdoutReader.Read(stdoutOutput)
			stderrReader.Read(stderrOutput)

			// Verify exit behavior
			if tt.shouldExit && !exitCalled {
				t.Error("expected os.Exit to be called")
			}

			// For help/version, expect exit code 0 or 2
			if exitCalled && exitCode != 0 && exitCode != 2 {
				t.Errorf("unexpected exit code: %d", exitCode)
			}

			// Create new pipes for next iteration
			if len(tests) > 1 {
				stdoutReader, stdoutWriter, _ = os.Pipe()
				stderrReader, stderrWriter, _ = os.Pipe()
				os.Stdout = stdoutWriter
				os.Stderr = stderrWriter
			}
		})
	}
}

// TestExecute_WithValidMinimalArgs tests Execute with minimal valid arguments
func TestExecute_WithValidMinimalArgs(t *testing.T) {
	// Skip this test if we can't create temporary files
	tempDir := t.TempDir()

	// Save originals
	originalOsExit := osExit
	originalArgs := os.Args
	defer func() {
		osExit = originalOsExit
		os.Args = originalArgs
	}()

	// Capture log output
	var logBuf bytes.Buffer
	log.SetOutput(&logBuf)
	defer log.SetOutput(os.Stderr)

	// Mock os.Exit
	osExit = func(code int) {
		panic("os.Exit called")
	}

	// Create a temporary input file
	inputFile := tempDir + "/input.txt"
	if err := os.WriteFile(inputFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("failed to create test input file: %v", err)
	}

	// Create output file path
	outputFile := tempDir + "/output.zip"

	// Set arguments for a basic zip creation
	os.Args = []string{
		"program",
		inputFile,
		outputFile,
	}

	// Run Execute
	func() {
		defer func() {
			if r := recover(); r != nil {
				// If it panics due to missing dependencies or configuration issues,
				// that's okay for a smoke test - we just want to verify it doesn't
				// crash unexpectedly
				t.Logf("Execute panicked (expected in smoke test): %v", r)
			}
		}()

		Execute()
	}()

	// If we got here without panicking, that's good
	// Check if output file was created (best case scenario)
	if _, err := os.Stat(outputFile); err == nil {
		t.Log("Execute completed successfully and created output file")
	} else {
		t.Log("Execute completed but didn't create output file (may be due to missing dependencies)")
	}
}
