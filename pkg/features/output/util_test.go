package output

import (
	"testing"
)

func testDebugName(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Fatalf("Expected debug name to be '%s' but got '%s' instead", expected, actual)
	}
}
