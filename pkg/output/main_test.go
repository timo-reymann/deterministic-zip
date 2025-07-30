package output

import (
	"path/filepath"
	"testing"
)

func TestInfo(t *testing.T) {
	Info("test")
}

func TestDebug(t *testing.T) {
	Debug("test")
}

func TestLevelAndSetOutput(t *testing.T) {
	// Test setting and getting log levels
	originalLevel := Level()

	// Test setting different levels
	SetLevel(LevelInfo) // assuming you have SetLevel and level constants
	if Level() != LevelInfo {
		t.Errorf("Expected level %v, got %v", LevelInfo, Level())
	}

	SetLevel(LevelDebug)
	if Level() != LevelDebug {
		t.Errorf("Expected level %v, got %v", LevelInfo, Level())
	}

	// Restore original level
	SetLevel(originalLevel)
}

func TestInfof(t *testing.T) {
	Infof("test %s", "foo")
}

func TestDebugf(t *testing.T) {
	Debugf("test %s", "foo")
}

func TestSetOutput(t *testing.T) {
	logPath := filepath.Join(t.TempDir(), "test.log")
	SetOutput(logPath)
	if logFile != logPath {
		t.Errorf("Expected log file %s, got %s", logPath, logFile)
	}
}
