package features

import (
	"testing"
)

func TestFeatures(t *testing.T) {
	if Features() == nil {
		t.Fatal("Features should NEVER return nil")
	}
}
